'use client';

import {DropdownMenu} from '@radix-ui/themes';
import React from 'react';

/**
 * ShareButton HOC. Which receives children as props, and make them a trigger for a dropdown menu.
 */
export function ShareButton({
  children,
  shortLink,
  disabled,
}: {
  children: React.ReactNode;
  shortLink?: string;
  disabled: boolean;
}) {
  return (
    <DropdownMenu.Root>
      <DropdownMenu.Trigger disabled={disabled}>{children}</DropdownMenu.Trigger>
      <DropdownMenu.Content>
        <DropdownMenu.Content>
          <DropdownMenu.Item>Facebook</DropdownMenu.Item>
          <DropdownMenu.Item>X</DropdownMenu.Item>
          <DropdownMenu.Separator />
          <DropdownMenu.Item>Telegram</DropdownMenu.Item>
          <DropdownMenu.Item>LinkedIn</DropdownMenu.Item>
        </DropdownMenu.Content>
      </DropdownMenu.Content>
    </DropdownMenu.Root>
  );
}
