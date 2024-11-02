import { Header } from '@/components/ui/header';
import { Button } from '@/components/ui/button';

import '@/styles/global.css';

const Home: React.FC = () => {
  return (
    <div>
      <Header />
      <h2>{import.meta.env.ENV_KEY}</h2>
      <Button>Click me</Button>
    </div>
  );
};

export default Home;