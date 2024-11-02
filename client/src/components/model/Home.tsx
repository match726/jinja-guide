import { Header } from '@/components/ui/header';
import { Button } from '@/components/ui/button';

import '@/styles/global.css';

const Home: React.FC = () => {
  return (
    <div>
      <Header />
      <Button>Click me</Button>
    </div>
  );
};

export default Home;