import styles from './page.module.css';
import {Flex} from '@radix-ui/themes';
import {MainHeading} from '@/ui/MainHeading';

export default function HomePage() {
  return (
    <div className={styles.page}>
      <main className={styles.main}>
        <Flex direction='column' align='center' gap='2'>
          <MainHeading />
        </Flex>
      </main>
    </div>
  );
}
