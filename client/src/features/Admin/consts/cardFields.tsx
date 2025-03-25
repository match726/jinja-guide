import { CardProps } from '@/features/Admin/types/cardProps';

export const CardFields: CardProps[] = [
  {
    title: "神社登録",
    furigana: "じんじゃとうろく",
    description: "神社名称と住所から神社を新規登録します。",
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
]