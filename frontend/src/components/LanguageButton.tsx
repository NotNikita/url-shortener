'use client';

import {Button, DropdownMenu} from '@radix-ui/themes';
import {PersonIcon} from '@radix-ui/react-icons';
import {useState} from 'react';

enum LanguageButtonType {
  Short = 'Short',
  Full = 'Full',
}

export const LanguageButton = ({type = LanguageButtonType.Full}: {type?: LanguageButtonType}) => {
  // TODO: localization
  const supportedLanguages = [
    {
      code: 'en',
      flag: '🇬🇧',
      lan: 'English',
    },
    {
      code: 'pl',
      flag: '🇵🇱',
      lan: 'Polish',
    },

    {
      code: 'by',
      flag: '🇧🇾',
      lan: 'Belarusian',
    },
    {
      code: 'ru',
      flag: '🇷🇺',
      lan: 'Russian',
    },
  ];
  const [curLanguage, setCurLanguage] = useState(supportedLanguages[0]);

  return (
    <DropdownMenu.Root dir='rtl'>
      <DropdownMenu.Trigger>
        <Button
          variant='soft'
          size='3'
          style={{
            width: '100%',
          }}
        >
          {curLanguage.flag} {curLanguage.lan}
        </Button>
      </DropdownMenu.Trigger>

      <DropdownMenu.Content>
        <DropdownMenu.Content>
          {supportedLanguages.map(lang => {
            const {code, flag, lan} = lang;
            const onClick = () => setCurLanguage(lang);

            return (
              <DropdownMenu.Item key={code} onClick={onClick}>
                {type === LanguageButtonType.Full ? `${flag} ${lan}` : flag}
              </DropdownMenu.Item>
            );
          })}
        </DropdownMenu.Content>
      </DropdownMenu.Content>
    </DropdownMenu.Root>
  );
};
