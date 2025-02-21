import QRCode from 'qrcode';
import {useState} from 'react';

export default function useQRCode() {
  const [qrCode, setQRCode] = useState<string>('');
  const [qrCodeSvg, setQRCodeSvg] = useState<string>('');
  const [qrCodePngJpeg, setQRCodePngJpeg] = useState<string>('');

  const generateQRCode = async (text: string) => {
    try {
      const dataUrl = await QRCode.toDataURL(text, {
        errorCorrectionLevel: 'M',
        width: 200,
      });
      setQRCode(dataUrl);
      setQRCodePngJpeg(dataUrl.split(',')[1]);

      const svgPath = await QRCode.toString(text, {type: 'svg'});
      setQRCodeSvg(btoa(svgPath));
    } catch (err) {
      console.error(err);
    }
  };

  console.log(qrCode);

  return {
    qrCode,
    qrCodeSvg,
    qrCodePngJpeg,
    generateQRCode,
  };
}
