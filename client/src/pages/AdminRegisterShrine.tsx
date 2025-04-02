import { useState } from "react";
import axios from 'axios';

import { BACKEND_ENDPOINT } from '@/config/config';
import { FieldProps, FormProps } from '@/features/register/types/types';
import { Header } from '@/components/ui/header';
import { RegisterForm } from '@/features/register/components/formField';
import { RegisterShrineFormProps } from '@/features/register/consts/formContents';

import '@/styles/global.css';

const AdminRegisterShrine = () => {

  // フォームの初期状態を定義
  const [formSections, setFormSections] = useState<FormProps[]>(RegisterShrineFormProps)

  // 特定のセクションにフィールドを追加
  const handleAddField = (sectionName: string) => {
    setFormSections((prevSections) =>
      prevSections.map((section) => {
        if (section.name === sectionName) {
          const newId = section.fields.length > 0 ? Math.max(...section.fields.map((field) => field.id)) + 1 : 1
          const newField: FieldProps = {
            id: newId,
            seq: "",
            content1: "",
            content2: "",
            content3: "",
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

  // 特定のセクションの特定のフィールドの値を更新
  const handleValueChange = (sectionName: string, fieldId: number, target: string, value: string) => {
    setFormSections((prevSections) =>
      prevSections.map((section) => {
        if (section.name === sectionName) {
          return {
            ...section,
            fields: section.fields.map((field) => (field.id === fieldId ? { ...field, [target]: value } : field)),
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

    const reqMap = new Map<string, FieldProps[]>();
    formSections.map(section => {
      let effectiveFields = section.fields.filter(field => field.content1 !== "");
      if (effectiveFields.length > 0) {
        reqMap.set(section.name, effectiveFields)
      }
    });

    const options = {
      method: "POST",
      url: BACKEND_ENDPOINT + "/api/admin/regist/shrine",
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
    setFormSections(RegisterShrineFormProps);

  }

  return (
    <>
      <Header />
      <div className="bg-gradient-to-b from-red-50 to-white flex items-top justify-center p-8">
        <RegisterForm title="神社登録" sections={formSections} onValueChange={handleValueChange} onAddField={handleAddField} onRemoveField={handleRemoveField} onSubmit={handleSubmit} />
      </div>
    </>
  );

};

export default AdminRegisterShrine;