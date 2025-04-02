import { FieldProps } from '@/types/types';

// フォーム属性の定義
export interface FormProps {
  name: string
  title: string
  placeHolder1: string
  placeHolder2?: string
  placeHolder3?: string
  type: "text" | "text2" | "select + text2"
  options?: string[]
  isMultiple: boolean
  fields: FieldProps[]
}
