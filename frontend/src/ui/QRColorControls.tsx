'use client';

import styles from '@/app/short/page.module.css';
import {Box, Flex, Text} from '@radix-ui/themes';
import {useState} from 'react';
import {TwitterPicker} from 'react-color';
import {Popover} from 'react-tiny-popover';

interface QRColorControlsProps {
  qrColor: string;
  qrBgColor: string;
  setQrColor: React.Dispatch<React.SetStateAction<string>>;
  setQrBgColor: React.Dispatch<React.SetStateAction<string>>;
}

export const QRColorControls: React.FC<QRColorControlsProps> = ({qrColor, qrBgColor, setQrColor, setQrBgColor}) => {
  const [showPalettePopover, setShowPalettePopover] = useState<'color' | 'bg' | undefined>();

  return (
    <Popover
      isOpen={!!showPalettePopover}
      positions={['right']}
      content={
        <TwitterPicker
          triangle='hide'
          colors={[
            '#FF6900',
            '#7BDCB5',
            '#00D084',
            '#8ED1FC',
            '#0693E3',
            '#EB144C',
            '#F78DA7',
            '#9900EF',
            '#FFF',
            '#000',
          ]}
          color={showPalettePopover === 'color' ? qrColor : qrBgColor}
          onChange={c => {
            if (showPalettePopover === 'color') {
              setQrColor(c.hex);
            } else {
              setQrBgColor(c.hex);
            }
          }}
        />
      }
      onClickOutside={() => setShowPalettePopover(undefined)}
    >
      <Box className={styles.paletteWrapper}>
        <Flex className={styles.paletteControl} gap='3' dir='row'>
          <Text className={styles.paletteText} size='2'>
            Code color
          </Text>
          <div
            className={styles.paletteColor}
            onClick={() => setShowPalettePopover(prev => (prev === 'color' ? undefined : 'color'))}
            style={{backgroundColor: qrColor ? qrColor : 'black'}}
          />
        </Flex>
        <Flex className={styles.paletteControl} gap='3' dir='row'>
          <Text className={styles.paletteText} size='2'>
            Background color
          </Text>
          <div
            className={styles.paletteColor}
            onClick={() => setShowPalettePopover(prev => (prev === 'bg' ? undefined : 'bg'))}
            style={{backgroundColor: qrBgColor ? qrBgColor : 'white'}}
          />
        </Flex>
      </Box>
    </Popover>
  );
};
