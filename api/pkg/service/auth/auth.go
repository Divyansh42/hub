// Copyright © 2020 The Tekton Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/go-github/github"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/tektoncd/hub/api/gen/auth"
	"github.com/tektoncd/hub/api/pkg/app"
	"github.com/tektoncd/hub/api/pkg/db/model"
	"github.com/tektoncd/hub/api/pkg/token"
)

type service struct {
	api           app.Config
	logger        *zap.SugaredLogger
	db            *gorm.DB
	oauth         *oauth2.Config
	jwtSigningKey string
}

var (
	invalidCode   = auth.MakeInvalidCode(fmt.Errorf("invalid authorization code"))
	internalError = auth.MakeInternalError(fmt.Errorf("failed to authenticate"))
)

// New returns the auth service implementation.
func New(api app.Config) auth.Service {
	return &service{api, api.Logger(), api.DB(), api.OAuthConfig(), api.JWTSigningKey()}
}

// Authenticates users against GitHub OAuth
func (s *service) Authenticate(ctx context.Context, p *auth.AuthenticatePayload) (res *auth.AuthenticateResult, err error) {

	// gets access_token for user using authorization_code
	token, err := s.oauth.Exchange(oauth2.NoContext, p.Code)
	if err != nil {
		return nil, invalidCode
	}

	// gets user details from github using the access_token
	oauthClient := s.oauth.Client(oauth2.NoContext, token)
	ghClient := github.NewClient(oauthClient)
	ghUser, _, err := ghClient.Users.Get(oauth2.NoContext, "")
	if err != nil {
		s.logger.Error(err)
		return nil, internalError
	}

	// adds user in db if not exist
	user, err := s.addUser(ghUser)
	if err != nil {
		return nil, err
	}

	// creates jwt using user details
	jwt, err := s.createJWT(user)
	if err != nil {
		return nil, err
	}

	return &auth.AuthenticateResult{Token: jwt}, nil
}

func (s *service) addUser(user *github.User) (*model.User, error) {

	q := s.db.Model(&model.User{}).Where(&model.User{GithubLogin: user.GetLogin()})

	newUser := &model.User{
		GithubName:  user.GetName(),
		GithubLogin: user.GetLogin(),
	}
	if err := q.FirstOrCreate(newUser).Error; err != nil {
		s.logger.Error(err)
		return nil, internalError
	}

	return newUser, nil
}

func (s *service) createJWT(user *model.User) (string, error) {

	claim := jwt.MapClaims{
		"id":     user.ID,
		"login":  user.GithubLogin,
		"name":   user.GithubName,
		"scopes": s.api.Data().Default.Scopes,
	}

	token, err := token.Create(claim, s.jwtSigningKey)
	if err != nil {
		s.logger.Error(err)
		return "", internalError
	}

	return token, nil
}
