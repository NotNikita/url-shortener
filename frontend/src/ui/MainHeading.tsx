import {Box, Heading, Strong, Text} from '@radix-ui/themes';

export const MainHeading = () => {
  return (
    <Box>
      <Heading as='h1' size={{xs: '4', sm: '5', md: '7', lg: '7'}} mb='2'>
        Make your link compact
      </Heading>
      <Text as='span' size={{xs: '1', sm: '2', md: '3', lg: '3'}}>
        <Strong>The fastest tool</Strong> for transforming <Strong>Looooong URLs</Strong> into short and working links!
      </Text>
    </Box>
  );
};
