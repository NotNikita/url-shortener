import Link from 'next/link';
import Image from 'next/image';
import {Box, Button, Flex, Text} from '@radix-ui/themes';
import {LinkBreak2Icon} from '@radix-ui/react-icons';
import {SigninButton} from '@/components/SigninButton';
import {LanguageButton} from '@/components/LanguageButton';
import {SidebarItem} from '@/components/SidebarItem';
import {IconQR} from '@/components/Icons/IconQR';

const menuIconProps = {
  style: {
    height: '1.5rem',
    width: '1.5rem',
    fill: 'currentColor',
  },
};

export const Sidebar = () => {
  return (
    <Flex
      data-testid='sidebar'
      direction='column'
      justify='between'
      style={{
        position: 'absolute',
        backgroundColor: 'var(--gray-a2)',
        top: 0,
        left: 0,
        height: '100vh',
        width: '240px',
        padding: '0.4rem 0.8rem',
        transition: 'all 0.5s ease',
      }}
    >
      <Box>
        <Link href='/'>
          <Image src='/horizontal-logo.svg' alt='Url shorting logo for sidebar' width={200} height={50} priority />
        </Link>
        <ul
          style={{
            marginTop: '3rem',
            gap: '20px',
          }}
        >
          <SidebarItem text='Shorten Url' link='/short' icon={<LinkBreak2Icon {...menuIconProps} />} />
          <SidebarItem text='Create QR' link='/qr-code' icon={<IconQR props={menuIconProps} />} />
        </ul>
      </Box>

      <Flex
        data-testid='sidebar-bottom'
        style={{
          marginBottom: '1rem',
          width: '100%',
        }}
        gap='1'
        direction='column'
      >
        <SigninButton />
        <LanguageButton />

        <Flex direction='row' gap='3'>
          <Link href='/privacy-policy'>
            <Text size='1'>privacy policy</Text>
          </Link>
          <Link href='/terms-conditions'>
            <Text size='1'>terms & conditions</Text>
          </Link>
        </Flex>
      </Flex>
    </Flex>
  );
};
