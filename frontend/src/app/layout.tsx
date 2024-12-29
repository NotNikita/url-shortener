import type {Metadata} from 'next';
import {nunito} from '@/app/fonts';
import {Theme} from '@radix-ui/themes';
import {ToastContainer} from 'react-toastify';

import './globals.css';
import '@radix-ui/themes/styles.css';
import './theme-config.css';
import {Sidebar} from '@/ui/Sidebar';

export const metadata: Metadata = {
  title: {
    template: '%s | Short.it',
    default: 'ShortIt!',
  },
  description: 'Make your links shorter!',
  metadataBase: new URL('https://github.com/NotNikita/url-shortener'),
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang='en'>
      <body className={`${nunito.variable} antialiased overflow-x-hidden`}>
        <Theme accentColor='crimson' grayColor='gray' panelBackground='solid' scaling='105%'>
          <Sidebar />
          {children}

          <ToastContainer position='bottom-right' />
        </Theme>
      </body>
    </html>
  );
}
