import { useState, useEffect, useRef } from 'react';
import axios from 'axios';

import { Header } from '@/components/ui/header';
import { Button } from '@/components/ui/button';

import '@/styles/global.css';

const backendEndpoint = import.meta.env.VITE_BACKEND_ENDPOINT;

const AdminBulkRegisterShrine = () => {

  const [submit, setSubmit] = useState(false);
  // 初回レンダリングのリクエスト送信を無効化
  const isFirstRender = useRef(true);

  useEffect(() => {

    const options = {
      method: "POST",
      url: backendEndpoint + "/api/admin/bulk-regist/shrine",
      headers: {
        "Content-Type": "application/json"
      }
    };

    if (isFirstRender.current) {
      isFirstRender.current = false;
      return
    } else {
      axios(options)
        .then((resp) => {
          console.log('POSTリクエストが成功しました', resp)
        })
        .catch((err) => console.error("POSTリクエスト失敗", err));
    }

  }, [submit]);

  const handleFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {

    // ページ遷移を防ぐ（デフォルトでは、フォーム送信ボタンを押すとページが遷移してしまう）
    e.preventDefault();

    setSubmit(true);

  };

  return (
    <>
      <Header />
      <div className="bg-gradient-to-b from-red-50 to-white flex items-top justify-center p-8">
        <div className="w-full max-w-md bg-white rounded-lg shadow-xl overflow-hidden">
          <div className="bg-red-900 p-4 flex items-center justify-center">
            <h2 className="text-2xl font-bold text-white ml-2 font-serif">神社一括登録</h2>
          </div>
          <form onSubmit={handleFormSubmit} className="p-6 space-y-6">
            <Button className="w-full bg-red-900 hover:bg-red-800 text-white font-bold py-2 px-4 rounded-md transition duration-300 ease-in-out transform hover:scale-105 font-serif">
              登録
            </Button>
          </form>
        </div>
      </div>
    </>

  );
  
};

export default AdminBulkRegisterShrine;