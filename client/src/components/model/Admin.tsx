import { Header } from '@/components/ui/header';
import { ShrineCard } from '@/components/ui/shrine-card';

import '@/styles/global.css';

const Admin: React.FC = () => {

  const cardData = [
    {
      title: "神社登録",
      description: "名称と住所から神社を新規登録します。",
      link: "admin/regist/shrine"
    },
    {
      title: "神社詳細情報登録",
      description: "PlusCodeから神社の詳細情報を新規登録します。",
      link: "admin/regist/details"
    },
    {
      title: "標準地域コード管理",
      description: "e-Statから最新の標準地域コードを取得、現在の登録の参照を行います。",
      link: "admin/stdareacode"
    }
  ];

  return (
    <>
      <Header />
      <div className="bg-gradient-to-b from-red-50 to-white">
        <h1 className="text-[min(4vw,30px)] flex py-4 items-center justify-center">
          管理者画面
        </h1>
        <section className="container flex grid grid-cols-2 gap-10 xl:grid-cols-3">
          {cardData.map((data) => (
            <ShrineCard
              cardTitle={data.title}
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