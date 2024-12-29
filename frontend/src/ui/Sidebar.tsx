import Link from 'next/link';
import Image from 'next/image';
import {Box, Button, Flex, Text} from '@radix-ui/themes';
import {EyeOpenIcon} from '@radix-ui/react-icons';
import {SigninButton} from '@/components/SigninButton';

// #1E1E24
// #92140c
// #fff8f0
// #ffcf99

export const Sidebar = () => {
  return (
    <Flex
      id='sidebar'
      direction='column'
      justify='between'
      style={{
        position: 'absolute',
        backgroundColor: '#1E1E24',
        top: 0,
        left: 0,
        height: '100vh',
        width: '240px',
        padding: '0.4rem 0.8rem',
        transition: 'all 0.5s ease',
      }}
    >
      <Box>
        <Link
          href='/'
          className='text-xl transition-colors duration-150 relative text-inherit whitespace-nowrap no-underline'
        >
          <Image
            src='/horizontal-logo.svg'
            alt='TutMenu logo for sidebar'
            width={200}
            height={50}
            className='block cursor-pointer'
            priority
          />
        </Link>
        <ul
          style={{
            marginTop: '3rem',
            gap: '20px',
          }}
        >
          <MenuItems />
          <MenuItems />
        </ul>
      </Box>

      <Box
        style={{
          marginBottom: '1rem',
          width: '100%',
        }}
      >
        <SigninButton />
      </Box>
    </Flex>
  );
};

const MenuItems = () => {
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
      <Link href='#'>
        <Button
          radius='medium'
          variant='outline'
          size='4'
          style={{
            color: '#FFF',
            border: '1px #fff8f0 solid',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'space-around',
            textDecoration: 'none',
            borderRadius: '0.8rem',
            width: '100%',
          }}
        >
          <EyeOpenIcon
            style={{
              height: '1.5rem',
              width: '1.5rem',
            }}
          />
          <Text style={{}}>Menu item 1</Text>
        </Button>
      </Link>
    </li>
  );
};
