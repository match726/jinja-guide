import { useState } from "react";
import { SlPlus, SlMinus } from 'react-icons/sl';
import axios from 'axios';

import { Header } from '@/components/ui/header';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';

import '@/styles/global.css';

const backendEndpoint = import.meta.env.VITE_BACKEND_ENDPOINT;

// フィールドの型定義
type Field = {
  id: number
  seq: string
  value: string
}

// セクションの型定義
type FormSection = {
  name: string
  title: string
  placeHolder: string
  type: "text" | "select + text"
  options?: string[]
  isMultiple: boolean
  fields: Field[]
}

// セクションの初期状態
const defaultFormSections: FormSection[] = [
  {
    name: "name",
    title: "神社名称",
    placeHolder: "例：伊勢神宮 内宮（皇大神宮）",
    type: "text",
    isMultiple: false,
    fields: [{ id: 1, seq: "", value: "" }],
  },
  {
    name: "furigana",
    title: "神社名称（振り仮名）",
    placeHolder: "例：いせじんぐう ないくう（こうたいじんぐう）",
    type: "text",
    isMultiple: false,
    fields: [{ id: 1, seq: "", value: "" }],
  },
  {
    name: "address",
    title: "住所",
    placeHolder: "例：三重県伊勢市宇治館町１",
    type: "text",
    isMultiple: false,
    fields: [{ id: 1, seq: "", value: "" }],
  },
  {
    name: "wikipediaUrl",
    title: "WikipediaURL",
    placeHolder: "例：https://ja.wikipedia.org/wiki/伊勢神宮",
    type: "text",
    isMultiple: false,
    fields: [{ id: 1, seq: "", value: "" }],
  }
]

const AdminRegisterShrine = () => {

  // フォームの初期状態を定義
  const [formSections, setFormSections] = useState<FormSection[]>(defaultFormSections)

  // 特定のセクションにフィールドを追加
  const handleAddField = (sectionName: string) => {
    setFormSections((prevSections) =>
      prevSections.map((section) => {
        if (section.name === sectionName) {
          const newId = section.fields.length > 0 ? Math.max(...section.fields.map((field) => field.id)) + 1 : 1
          const newField: Field = {
            id: newId,
            seq: "",
            value: "",
          }
          return {
            ...section,
            fields: [...section.fields, newField],
          }
        }
        return section
      }),
    )
  }

  // 特定のセクションから特定のフィールドを削除
  const handleRemoveField = (sectionName: string, fieldId: number) => {
    setFormSections((prevSections) =>
      prevSections.map((section) => {
        if (section.name === sectionName) {
          // 最後の1つは削除しない
          if (section.fields.length === 1) return section

          return {
            ...section,
            fields: section.fields.filter((field) => field.id !== fieldId),
          }
        }
        return section
      }),
    )
  }

  // 特定のセクションの特定のフィールドの値（value）を更新
  const handleValueChange = (sectionName: string, fieldId: number, value: string) => {
    setFormSections((prevSections) =>
      prevSections.map((section) => {
        if (section.name === sectionName) {
          return {
            ...section,
            fields: section.fields.map((field) => (field.id === fieldId ? { ...field, value } : field)),
          }
        }
        return section
      }),
    )
  }

  // 特定のセクションの特定のフィールドの値（seq）を更新
  const handleSeqChange = (sectionName: string, fieldId: number, seq: string) => {
    setFormSections((prevSections) =>
      prevSections.map((section) => {
        if (section.name === sectionName) {
          return {
            ...section,
            fields: section.fields.map((field) => (field.id === fieldId ? { ...field, seq } : field)),
          }
        }
        return section
      }),
    )
  }

  // 登録ボタン押下時の処理
  const handleSubmit = (e: React.FormEvent) => {

    // ページ遷移を防ぐ（デフォルトでは、フォーム送信ボタンを押すとページが遷移してしまう）
    e.preventDefault()

    const reqMap = new Map<string, Field[]>();
    formSections.map(section => {
      let effectiveFields = section.fields.filter(field => field.value !== "");
      if (effectiveFields.length > 0) {
        reqMap.set(section.name, effectiveFields)
      }
    });

    const options = {
      method: "POST",
      url: backendEndpoint + "/api/admin/regist/shrine",
      headers: {
        "Content-Type": "application/json"
      },
      data: JSON.stringify(Object.fromEntries(reqMap))
    };

    axios(options)
      .then((resp) => {
        console.log('POSTリクエストが成功しました', resp)
      })
      .catch((err) => console.error("POSTリクエスト失敗", err));

    // フォームを初期状態に戻す
    setFormSections(defaultFormSections);

  }

  // フィールドのレンダリング関数
  const renderField = (section: FormSection, field: Field) => {
    switch (section.type) {
      case "text":
        return (
          <div className="flex justify-center items-center gap-2">
            <Input
              id={`${section.name}-${field.id}`}
              value={field.value}
              onChange={(e) => handleValueChange(section.name, field.id, e.target.value)}
              placeholder={section.placeHolder}
              className="w-full border-2 border-red-800 rounded-md p-2 font-serif"
            />
            {section.isMultiple == true
              ? <div>
                  <SlPlus onClick={() => handleAddField(section.name)} className="h-5 w-5" />
                  <SlMinus onClick={() => handleRemoveField(section.name, field.id)} className="h-5 w-5" />
                </div>
              : null
            }
          </div>
        )
      case "select + text":
        return (
          <div className="flex justify-center items-center gap-2">
            <Select value={field.seq} onValueChange={(value) => handleSeqChange(section.name, field.id, value)}>
              <SelectTrigger className="w-full w-max-md border-2 border-red-800 rounded-md p-2 font-serif">
                <SelectValue placeholder={`${section.title}を選択`} />
              </SelectTrigger>
              <SelectContent>
                {section.options?.map((option) => (
                  <SelectItem key={option} value={option.slice(0,option.indexOf(":"))}>
                    {option}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            <Input
              id={`${section.name}-${field.id}`}
              value={field.value}
              onChange={(e) => handleValueChange(section.name, field.id, e.target.value)}
              placeholder={section.placeHolder}
              className="w-full w-max-md border-2 border-red-800 rounded-md p-2 font-serif"
            />
            {section.isMultiple == true
              ? <div>
                  <SlPlus onClick={() => handleAddField(section.name)} className="h-5 w-5" />
                  <SlMinus onClick={() => handleRemoveField(section.name, field.id)} className="h-5 w-5" />
                </div>
              : null
            }
          </div>
        )
      default:
        return null
    }
  }

  return (
    <>
      <Header />
      <div className="bg-gradient-to-b from-red-50 to-white flex items-top justify-center p-8">
        <div className="w-full max-w-lg bg-white rounded-lg shadow-xl overflow-hidden">
          <div className="bg-red-900 p-4 flex items-center justify-center">
            <h2 className="text-2xl font-bold text-white ml-2 font-serif">
              神社登録
            </h2>
          </div>
          <form onSubmit={handleSubmit} className="p-6 space-y-6">
            {formSections.map((section) => (
              <div key={section.name}>
                <Label htmlFor={section.name} className="text-lg font-medium text-gray-700 font-serif">
                  {section.title}
                </Label>
                {section.fields.map((field) => (
                  <div key={field.id} className="flex items-end gap-2">
                    <div className="flex-1">
                      {renderField(section, field)}
                    </div>
                  </div>
                ))}
              </div>
            ))}
            <Button className="w-full bg-red-900 hover:bg-red-800 text-white font-bold py-2 px-4 rounded-md transition duration-300 ease-in-out transform hover:scale-105 font-serif">
              登録
            </Button>
          </form>
        </div>
      </div>
    </>
  );
  
};

export default AdminRegisterShrine;