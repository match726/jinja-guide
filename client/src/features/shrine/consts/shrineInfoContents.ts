import { ShrineDetails } from '@/features/shrine/types/types';

// 神社詳細情報画面のレスポンスデータ初期化
export const InitShrineDetails: ShrineDetails = {
  name: "",
  furigana: {id: 0, seq: "", content1: "", content2: "", content3: ""},
  altName: [],
  address: "",
  placeId: "",
  image: "",
  description: {id: 0, seq: "", content1: "", content2: "", content3: ""},
  tags: [],
  foundedYear: {id: 0, seq: "", content1: "", content2: "", content3: ""},
  objectOfWorship: [],
  shrineRank: [],
  hasGoshuin: {id: 0, seq: "", content1: "", content2: "", content3: ""},
  websiteUrl: {id: 0, seq: "", content1: "", content2: "", content3: ""},
  wikipediaUrl: {id: 0, seq: "", content1: "", content2: "", content3: ""}
};