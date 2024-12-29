'use client';

import {Button} from '@radix-ui/themes';
import {PersonIcon} from '@radix-ui/react-icons';

export const SigninButton = () => {
  const isSignedIn = false;

  return (
    <Button
      variant='soft'
      size='3'
      style={{
        width: '100%',
      }}
    >
      <PersonIcon />
      {isSignedIn ? 'Sign out' : 'Sign in'}
    </Button>
  );
};
