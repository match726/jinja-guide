import { Header } from '@/components/ui/header';
import { OrdinaryCard } from '@/components/ui/card/ordinary-card';

import { CardFields } from '@/features/Admin/consts/cardFields';

import '@/styles/global.css';

const Admin: React.FC = () => {

  return (
    <>
      <Header />
      <div className="bg-gradient-to-b from-red-50 to-white">
        <h1 className="text-[min(4vw,30px)] flex py-4 items-center justify-center">
          管理者画面
        </h1>
        <section className="container flex grid grid-cols-2 gap-10 xl:grid-cols-3">
          {CardFields.map((data) => (
            <OrdinaryCard
              cardTitle={data.title}
              cardTitleRuby={data.furigana}
              cardDescription={data.description}
              cardLink={data.link}
            />
          ))}
        </section>
      </div>
    </>
  );
};

export default Admin;