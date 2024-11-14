import { Header } from '@/components/ui/header';

import '@/styles/global.css';

const Home: React.FC = () => {
  return (
    <div>
      <Header />
      <h2>環境名：{import.meta.env.VITE_ENV_KEY}</h2>
    </div>
  );
};

export default Home;