import Image from 'next/image';
import styles from './page.module.css';
import {Flex, Text, Button} from '@radix-ui/themes';

export default function Home() {
  return (
    <div className={styles.page}>
      <main className={styles.main}>
        <Image className={styles.logo} src='/next.svg' alt='Next.js logo' width={180} height={38} priority />
        <div className={styles.ctas}>Hello Next.js</div>
        <Flex direction='column' align='center' gap='2'>
          <Text>Testing Theme panel</Text>
          <Button>Lets Go!</Button>
        </Flex>
      </main>
    </div>
  );
}
