import { FormProps } from '@/features/register/types/types';

// 神社登録フォーム表示内容の初期状態
export const RegisterShrineFormProps: FormProps[] = [
  {
    name: "name",
    title: "神社名称",
    placeHolder1: "例：伊勢神宮 内宮（皇大神宮）",
    type: "text",
    isMultiple: false,
    fields: [{ id: 1, seq: "", content1: "", content2: "", content3: "" }],
  },
  {
    name: "furigana",
    title: "神社名称（振り仮名）",
    placeHolder1: "例：いせじんぐう ないくう（こうたいじんぐう）",
    type: "text",
    isMultiple: false,
    fields: [{ id: 1, seq: "", content1: "", content2: "", content3: "" }],
  },
  {
    name: "address",
    title: "住所",
    placeHolder1: "例：三重県伊勢市宇治館町１",
    type: "text",
    isMultiple: false,
    fields: [{ id: 1, seq: "", content1: "", content2: "", content3: "" }],
  },
  {
    name: "wikipediaUrl",
    title: "WikipediaURL",
    placeHolder1: "例：https://ja.wikipedia.org/wiki/伊勢神宮",
    type: "text",
    isMultiple: false,
    fields: [{ id: 1, seq: "", content1: "", content2: "", content3: "" }],
  }
]

// 神社詳細情報登録フォーム表示内容の初期状態
export const RegisterShrineDetailsFormProps: FormProps[] = [
  {
    name: "plusCode",
      title: "PlusCode",
    placeHolder1: "例：8Q6RFP4G+255",
    type: "text",
    isMultiple: false,
    fields: [{ id: 1, seq: "", content1: "", content2: "", content3: "" }],
  },
  {
    name: "furigana",
      title: "神社名称（振り仮名）",
    placeHolder1: "例：いせじんぐう ないくう（こうたいじんぐう）",
    type: "text",
    isMultiple: false,
    fields: [{ id: 1, seq: "", content1: "", content2: "", content3: "" }],
  },
  {
    name: "altName",
      title: "別名称",
    placeHolder1: "例：伊勢神宮",
    type: "text",
    isMultiple: true,
    fields: [{ id: 1, seq: "", content1: "", content2: "", content3: "" }],
  },
  {
    name: "tag",
      title: "関連ワード",
    placeHolder1: "例：お伊勢参り",
    type: "text",
    isMultiple: true,
    fields: [{ id: 1, seq: "", content1: "", content2: "", content3: "" }],
  },
  {
    name: "foundedYear",
      title: "創建年",
    placeHolder1: "例：垂仁天皇26年",
    placeHolder2: "例：伝",
    type: "text2",
    isMultiple: false,
    fields: [{ id: 1, seq: "", content1: "", content2: "", content3: "" }],
  },
  {
    name: "objectOfWorship",
    title: "御祭神",
    placeHolder1: "例：天照坐皇大御神",
    type: "text",
    isMultiple: true,
    fields: [{ id: 1, seq: "", content1: "", content2: "", content3: "" }],
  },
  {
    name: "shrineRank",
    title: "社格",
    placeHolder1: "例：並大社",
    placeHolder2: "例：論社",
    type: "select + text2",
    options: ["1: 延喜式内社", "2: 国史見在社", "3: 二十二社制度", "4: 一宮制度", "5: 総社（惣社）", "6: 近代社格制度", "7: 別表神社"],
    isMultiple: true,
    fields: [{ id: 1, seq: "", content1: "", content2: "", content3: "" }],
  },
  {
    name: "hasGoshuin",
      title: "御朱印",
    placeHolder1: "例：あり",
    type: "text",
    isMultiple: false,
    fields: [{ id: 1, seq: "", content1: "", content2: "", content3: "" }],
  },
  {
    name: "websiteUrl",
      title: "公式サイトURL",
    placeHolder1: "例：https://www.isejingu.or.jp/",
    type: "text",
    isMultiple: false,
    fields: [{ id: 1, seq: "", content1: "", content2: "", content3: "" }],
  },
  {
    name: "wikipediaUrl",
      title: "WikipediaURL",
    placeHolder1: "例：https://ja.wikipedia.org/wiki/伊勢神宮",
    type: "text",
    isMultiple: false,
    fields: [{ id: 1, seq: "", content1: "", content2: "", content3: "" }],
  }
]
