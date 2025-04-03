import { FieldProps } from '@/types/types';

// 神社詳細情報画面のレスポンスデータ定義
export type ShrineDetails = {
  name: string
  furigana: FieldProps
  image: string
  altName: FieldProps[]
  address: string
  placeId: string
  description: FieldProps
  tags: FieldProps[]
  foundedYear: FieldProps
  objectOfWorship: FieldProps[]
  shrineRank: FieldProps[]
  hasGoshuin: FieldProps
  websiteUrl: FieldProps
  wikipediaUrl: FieldProps
};
