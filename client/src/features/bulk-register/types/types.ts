// 神社一括登録テーブルの定義
export interface RegisterShrineProps {
  name: string
  address: string
  furigana: string
  altName: string[]
  tags: string[]
  foundedYear: string[]
  objectOfWorship: string[]
  shrineRank: string[]
  hasGoshuin: string
  websiteUrl: string
  wikipediaUrl: string
}