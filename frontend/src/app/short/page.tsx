'use client';

import styles from './page.module.css';
import {Box, Button, Card, Container, Flex, Heading, IconButton, Text, TextField, Tooltip} from '@radix-ui/themes';
import {GlobeIcon, Share1Icon, GearIcon, Link2Icon, Cross1Icon} from '@radix-ui/react-icons';
import {MainHeading} from '@/ui/MainHeading';
import {useEffect, useState} from 'react';
import {toast} from 'react-toastify';
import Image from 'next/image';
import {ShareButton} from '@/components/ShareButton';
import Link from 'next/link';
import useQRCode from '@/hooks/useQRCode';
import {QRWithControls} from '@/ui/QRWithControls';

const MAX_LONG_URL_LENGTH = 100;

export default function ShortPage() {
  const [originUrl, setOriginUrl] = useState('');
  const [lastProcessedUrl, setLastProcessedUrl] = useState('');
  const [shortUrl, setShortUrl] = useState('');
  const [isDisabled, setDisabled] = useState(false);

  const {qrCode, qrCodeSvg, qrCodePngJpeg, generateQRCode} = useQRCode();

  const trimmedShortUrl = shortUrl.replace(/^https?:\/\//, '');
  const symbolsLeft = MAX_LONG_URL_LENGTH - originUrl.length;

  const onInputChange = (e: any) => {
    setOriginUrl(e.target.value);
  };

  const onShortButtonClick = () => {
    setLastProcessedUrl(originUrl);
    // fake:
    setShortUrl('https://short.url/abc123');

    // TODO: fetch api
  };

  const copyToClipboard = async () => {
    try {
      await navigator.clipboard.writeText(shortUrl);
      toast.info('✅ Copied to clipboard!');
    } catch (err) {
      console.error('Failed to copy text: ', err);
    }
  };

  useEffect(() => {
    if (!originUrl) return;

    generateQRCode(originUrl);
    if (originUrl === lastProcessedUrl) {
      setDisabled(true);
    } else {
      setDisabled(false);
    }
  }, [originUrl, lastProcessedUrl]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onShortButtonClick();
  };

  return (
    <Box className={styles.container}>
      <Container className={styles.content}>
        <Box className={styles.header}>
          <MainHeading />
        </Box>

        <Card size={{sm: '3', md: '3', lg: '3'}}>
          <Box p={{sm: '1', md: '2', lg: '3'}}>
            <Heading as='h2' size='3' mb='4'>
              Paste long link here:
            </Heading>
            <form onSubmit={handleSubmit}>
              <Flex
                gap='3'
                height='auto'
                className={styles.inputContainer}
                direction={{initial: 'column', xs: 'column', sm: 'row', md: 'row', lg: 'row'}}
              >
                <Box className={styles.inputWrapper}>
                  <GlobeIcon className={styles.icon} />
                  <TextField.Root
                    type='url'
                    placeholder='https://example.com/very-loooong-url'
                    value={originUrl}
                    onChange={onInputChange}
                    required
                    className={styles.inputRoot}
                  >
                    <TextField.Slot side='right' color={symbolsLeft < 0 ? 'red' : undefined}>
                      {symbolsLeft}
                    </TextField.Slot>
                    <TextField.Slot side='right' style={{cursor: 'pointer'}} onClick={() => setOriginUrl('')}>
                      <Cross1Icon width={20} height={20} color='black' />
                    </TextField.Slot>
                  </TextField.Root>
                </Box>
                <Button type='submit' size='3' disabled={isDisabled}>
                  Short it!
                </Button>
              </Flex>
              {symbolsLeft < 0 && (
                <Text size='2' mt='2' color='red'>
                  The maximum number of characters has been exceeded.
                </Text>
              )}
            </form>
          </Box>
        </Card>

        {shortUrl && (
          <Card size={{sm: '3', md: '3', lg: '3'}} mt='4'>
            <Box p={{sm: '1', md: '2', lg: '3'}}>
              <Flex
                className={styles.resultControls}
                direction={{initial: 'column', xs: 'column', sm: 'row', md: 'row', lg: 'row'}}
              >
                <Box className={styles.resultContainer}>
                  <div className={styles.resultLink}>
                    <Link href={shortUrl} target='_blank'>
                      <Text className={styles.shortUrl} weight='bold' size='6'>
                        <Link2Icon width={20} height={20} fill='inherit' />
                        {trimmedShortUrl}
                      </Text>
                    </Link>
                  </div>

                  <Flex className={styles.resultMiniIcons}>
                    <Tooltip content='Copy to clipboard'>
                      <Button size='3' variant='soft' onClick={copyToClipboard}>
                        Copy
                      </Button>
                    </Tooltip>
                    <Tooltip content='Share URL'>
                      <ShareButton shortLink={shortUrl} disabled>
                        <IconButton size='3' variant='soft'>
                          <Share1Icon width={20} height={20} color='black' />
                        </IconButton>
                      </ShareButton>
                    </Tooltip>
                    <Tooltip content='Settings'>
                      <IconButton size='3' variant='soft' disabled>
                        <GearIcon width={20} height={20} color='black' />
                      </IconButton>
                    </Tooltip>
                    <Tooltip content='Customize QR Code'>
                      <IconButton size='3' variant='soft' disabled>
                        <Image
                          className={styles.qrCodeIcon}
                          src='/svg/palette-icon.svg'
                          alt='Customize QR code'
                          width={20}
                          height={20}
                        />
                      </IconButton>
                    </Tooltip>
                    <Tooltip content='Show QR Code'>
                      <IconButton size='3' variant='soft'>
                        <Image
                          className={styles.qrCodeIcon}
                          src='/svg/qr-code.svg'
                          alt='QR code icon'
                          width={20}
                          height={20}
                        />
                      </IconButton>
                    </Tooltip>
                  </Flex>
                </Box>
                {qrCode && (
                  <QRWithControls
                    shortUrl={shortUrl}
                    qrCodeImageSrc={qrCode}
                    qrCodeSvg={qrCodeSvg}
                    qrCodePngJpeg={qrCodePngJpeg}
                  />
                )}
              </Flex>
            </Box>
          </Card>
        )}
      </Container>
    </Box>
  );
}
