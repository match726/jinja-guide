import { Header } from '@/components/ui/header';
import { OrdinaryCard } from '@/components/ui/card/ordinary-card';

import '@/styles/global.css';

const Admin: React.FC = () => {

  const cardData = [
    {
      title: "神社登録",
      furigana: "じんじゃとうろく",
      description: "名称と住所から神社を新規登録します。",
      link: "admin/regist/shrine"
    },
    {
      title: "神社一括登録",
      furigana: "じんじゃいっかつとうろく",
      description: "神社一括登録テーブルから一括登録します。",
      link: "admin/bulk-regist/shrine"
    },
    {
      title: "神社詳細情報登録",
      furigana: "じんじゃしょうさいじょうほうとうろく",
      description: "PlusCodeから神社の詳細情報を新規登録します。",
      link: "admin/regist/shrine-details"
    },
    {
      title: "標準地域コード管理",
      furigana: "ひょうじゅんちいきこーどかんり",
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