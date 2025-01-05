import React, {JSX} from 'react';
import {Button, Tooltip} from '@radix-ui/themes';

interface IconButtonProps {
  icon: JSX.Element;
  tooltipContent: string;
  onClick: () => void;
}

export function IconButton({icon: Icon, tooltipContent, onClick}: IconButtonProps) {
  return (
    <Tooltip content={tooltipContent}>
      <Button size='2' variant='soft' onClick={onClick}>
        {Icon}
      </Button>
    </Tooltip>
  );
}
