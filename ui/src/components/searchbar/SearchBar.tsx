import { TextInput } from '@patternfly/react-core';
import React, { useState } from 'react';

const SearchBar: React.FC = () => {
  const [value, setValue] = useState('');

  const setInput = (event: any) => {
    setValue(event);
  };

  return (
    <TextInput
      value={value}
      type="search"
      onChange={setInput}
      aria-label="text input example"
      placeholder="Search for resources..."
      spellCheck='false'
      style={{
        marginLeft: '-10em',
        marginRight: '20em',
        borderRadius: '4px',
        background: '#2e2e2e',
        border: '0px',
        fontStyle: 'italic',
        color: '#ffff',
      }}
    />
  );
};

export default SearchBar;
