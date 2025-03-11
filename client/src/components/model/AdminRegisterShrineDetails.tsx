import { useState } from "react";
//import axios from 'axios';

import { Header } from '@/components/ui/header';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';

import '@/styles/global.css';

//const backendEndpoint = import.meta.env.VITE_BACKEND_ENDPOINT;

// フィールドの型定義
type Field = {
  id: number
  value: string
  type: "text" | "select + text"
  options?: string[]
}

// セクションの型定義
type FormSection = {
  id: string
  title: string
  placeHolder: string
  fields: Field[]
}

const AdminRegisterShrineDetails = () => {

  // 初回レンダリングのリクエスト送信を無効化
  //const isFirstRender = useRef(true);

  // フォームの初期状態を定義
  const [formSections, setFormSections] = useState<FormSection[]>([
    {
      id: "plusCode",
      title: "PlusCode",
      placeHolder: "例：8Q6RFP4G+255",
      fields: [{ id: 1, value: "", type: "text" }],
    },
    {
      id: "furigana",
      title: "神社名称（振り仮名）",
      placeHolder: "例：いせじんぐう ないくう（こうたいじんぐう）",
      fields: [{ id: 1, value: "", type: "text" }],
    },
    {
      id: "altName",
      title: "別名称",
      placeHolder: "例：伊勢神宮",
      fields: [{ id: 1, value: "", type: "text" }],
    },
    {
      id: "tag",
      title: "関連ワード",
      placeHolder: "例：お伊勢参り",
      fields: [{ id: 1, value: "", type: "text" }],
    },
    {
      id: "foundedYear",
      title: "創建年",
      placeHolder: "例：垂仁天皇26年",
      fields: [{ id: 1, value: "", type: "text" }],
    },
    {
      id: "objectOfWorship",
      title: "御祭神",
      placeHolder: "例：天照坐皇大御神",
      fields: [{ id: 1, value: "", type: "text" }],
    },
    {
      id: "shrineRank",
      title: "社格",
      placeHolder: "例：式内社",
      fields: [{ id: 1, value: "", type: "select + text", options: ["1: 延喜式内社", "2: 国史見在社", "3: 二十二社制度", "4: 一宮制度", "5: 総社（惣社）", "6: 近代社格制度", "7: 別表神社"] }],
    },
    {
      id: "hasGoshuin",
      title: "御朱印",
      placeHolder: "例：あり",
      fields: [{ id: 1, value: "", type: "text" }],
    },
    {
      id: "websiteUrl",
      title: "公式サイトURL",
      placeHolder: "例：https://www.isejingu.or.jp/",
      fields: [{ id: 1, value: "", type: "text" }],
    },
    {
      id: "wikipediaUrl",
      title: "WikipediaURL",
      placeHolder: "例：https://ja.wikipedia.org/wiki/伊勢神宮",
      fields: [{ id: 1, value: "", type: "text" }],
    }
  ])

  // 特定のセクションにフィールドを追加
  const handleAddField = (sectionId: string) => {
    setFormSections((prevSections) =>
      prevSections.map((section) => {
        if (section.id === sectionId) {
          const newId = section.fields.length > 0 ? Math.max(...section.fields.map((field) => field.id)) + 1 : 1
          const newField: Field = {
            id: newId,
            value: "",
            type: section.fields[0].type, // 最初のフィールドと同じタイプを使用
            options: section.fields[0].options, // オプションがある場合はそれも複製
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
  const handleRemoveField = (sectionId: string, fieldId: number) => {
    setFormSections((prevSections) =>
      prevSections.map((section) => {
        if (section.id === sectionId) {
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
  
  // useEffect(() => {

  //   const options = {
  //     method: "POST",
  //     url: backendEndpoint + "/api/admin/regist/shrine-details",
  //     headers: {
  //       "Content-Type": "application/json"
  //     },
  //     data: JSON.stringify(formSections)
  //   };

  //   // if (isFirstRender.current) {
  //   //   isFirstRender.current = false;
  //   //   return
  //   // } else {
  //   //   axios(options)
  //   //     .then((resp) => {
  //   //       console.log('POSTリクエストが成功しました', resp)
  //   //     })
  //   //     .catch((err) => console.error("POSTリクエスト失敗", err));
  //   // }

  // }, [formSections]);

  // 特定のセクションの特定のフィールドの値を更新
  const handleInputChange = (sectionId: string, fieldId: number, value: string) => {
    setFormSections((prevSections) =>
      prevSections.map((section) => {
        if (section.id === sectionId) {
          return {
            ...section,
            fields: section.fields.map((field) => (field.id === fieldId ? { ...field, value } : field)),
          }
        }
        return section
      }),
    )
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    console.log("Form submitted with values:", formSections)
    // ここでバックエンドにデータを送信する処理を行う
  }

  // フィールドのレンダリング関数
  const renderField = (section: FormSection, field: Field) => {
    switch (field.type) {
      case "text":
        return (
          <div>
            <Input
              id={`${section.id}-${field.id}`}
              value={field.value}
              onChange={(e) => handleInputChange(section.id, field.id, e.target.value)}
              placeholder={section.placeHolder}
              className="w-full border-2 border-red-800 rounded-md p-2 font-serif"
            />
            <PlusIcon onClick={() => handleAddField(section.id)} className="h-5 w-5 text-muted-foreground" />
            <PlusIcon onClick={() => handleRemoveField(section.id, field.id)} className="h-4 w-4" />
          </div>
        )
      case "select + text":
        return (
          <div>
            <Select value={field.value} onValueChange={(value) => handleInputChange(section.id, field.id, value)}>
              <SelectTrigger className="w-full w-max-md border-2 border-red-800 rounded-md p-2 font-serif">
                <SelectValue placeholder={`${section.title}を選択`} />
              </SelectTrigger>
              <SelectContent>
                {field.options?.map((option) => (
                  <SelectItem key={option} value={option}>
                    {option}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            <PlusIcon onClick={() => handleAddField(section.id)} className="h-4 w-4" />
            <PlusIcon onClick={() => handleRemoveField(section.id, field.id)} className="h-4 w-4" />
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
              神社詳細情報登録
            </h2>
          </div>
          <form onSubmit={handleSubmit} className="p-6 space-y-6">
            {formSections.map((section) => (
              <div key={section.id}>
                <Label htmlFor={section.id} className="text-lg font-medium text-gray-700 font-serif">
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

function PlusIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      x="0px"
      y="0px"
      width="100"
      height="100"
      viewBox="0 0 50 50"
    >
      <path d="M 25 2 C 12.264481 2 2 12.264481 2 25 C 2 37.735519 12.264481 48 25 48 C 37.735519 48 48 37.735519 48 25 C 48 12.264481 37.735519 2 25 2 z M 25 4 C 36.664481 4 46 13.335519 46 25 C 46 36.664481 36.664481 46 25 46 C 13.335519 46 4 36.664481 4 25 C 4 13.335519 13.335519 4 25 4 z M 24 13 L 24 24 L 13 24 L 13 26 L 24 26 L 24 37 L 26 37 L 26 26 L 37 26 L 37 24 L 26 24 L 26 13 L 24 13 z" />
    </svg>
  )
}

export default AdminRegisterShrineDetails;