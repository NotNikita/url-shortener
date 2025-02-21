import {Box, Flex, Button} from '@radix-ui/themes';
import Link from 'next/link';
import Image from 'next/image';

interface QRWithControlsProps {
  shortUrl: string;
  qrCodeImageSrc: string;
  qrCodeSvg: string;
  qrCodePngJpeg: string;
}

export const QRWithControls: React.FC<QRWithControlsProps> = ({shortUrl, qrCodeImageSrc, qrCodeSvg, qrCodePngJpeg}) => {
  return (
    <Box dir='column'>
      <Link href={shortUrl} target='_blank'>
        <Image src={qrCodeImageSrc} alt='Shortening service QR example' width={200} height={200} />
      </Link>
      <Flex dir='row' gap='3'>
        <a download='MyQRCode.svg' href={`data:image/svg+xml;base64,${qrCodeSvg}`}>
          <Button size='2' variant='soft' onClick={() => {}}>
            SVG
          </Button>
        </a>
        <a download='MyQRCode.png' href={qrCodeImageSrc}>
          <Button size='2' variant='soft' onClick={() => {}}>
            PNG
          </Button>
        </a>
        <a download='MyQRCode.jpeg' href={`data:image/jpeg;base64,${qrCodePngJpeg}`}>
          <Button size='2' variant='soft' onClick={() => {}}>
            JPEG
          </Button>
        </a>
      </Flex>
    </Box>
  );
};
