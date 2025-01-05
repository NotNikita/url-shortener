import {Button, Text} from '@radix-ui/themes';
import Link from 'next/link';
import {JSX} from 'react';

export const SidebarItem = ({text, link, icon}: {text: string; link: string; icon: JSX.Element}) => {
  return (
    <li
      style={{
        position: 'relative',
        listStyleType: 'none',
        height: '3rem',
        lineHeight: '3.25rem',
        width: '100%',
        margin: '0.8rem auto',
      }}
    >
      <Link href={link}>
        <Button
          radius='medium'
          variant='outline'
          size='4'
          style={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'space-around',
            textDecoration: 'none',
            borderRadius: '0.8rem',
            width: '100%',
          }}
        >
          {icon}
          <Text>{text}</Text>
        </Button>
      </Link>
    </li>
  );
};
