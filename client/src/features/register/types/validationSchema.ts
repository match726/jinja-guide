import { z } from 'zod';

// 神社名称のフィールド定義
const nameSchema = z.object({
  seq: z
    .literal(""),
  content1: z
    .string({ required_error: "神社名称は必須入力項目です。" })
    .max(20, { message: "20文字以下で入力してください。" })
    .regex(/^[^\x01-\x7E\xA1-\xDF]+$/, { message: "全角で入力してください。" }),
  content2: z
    .literal(""),
  content3: z
    .literal("")
});

// 住所のフィールド定義
const addressSchema = z.object({
  seq: z
    .literal(""),
  content1: z
    .string({ required_error: "住所は必須入力項目です。" })
    .max(40, { message: "40文字以下で入力してください。" })
    .regex(/^[^\x01-\x7E\xA1-\xDF]+$/, { message: "全角で入力してください。" }),
  content2: z
    .literal(""),
  content3: z
    .literal("")
});

// 神社名称（振り仮名）のフィールド定義
const furiganaSchema = z.object({
  seq: z
    .literal(""),
  content1: z
    .string()
    .regex(/^[\u3041-\u3096]+$/, { message: "全角ひらがなで入力してください。" }),
  content2: z
    .literal(""),
  content3: z
    .literal("")
});

// WikipediaURLのフィールド定義
const wikipediaUrlSchema = z.object({
  seq: z
    .literal(""),
  content1: z
    .union([
      z.string().url({ message: "URL形式で入力してください。" }),
      z.literal("")
    ]),
  content2: z
    .literal(""),
  content3: z
    .literal("")
});

export const RegisterShrineSchema = z.object({
  name: z.array(nameSchema),
  furigana: z.array(furiganaSchema),
  address: z.array(addressSchema),
  wikipediaUrl: z.array(wikipediaUrlSchema)
});

export type RegisterShrine = z.infer<typeof RegisterShrineSchema>;