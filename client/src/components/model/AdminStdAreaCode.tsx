import { useState, useEffect } from 'react';
import axios from 'axios';

import { Header } from '@/components/ui/header';
import { Button } from '@/components/ui/button';

const backendEndpoint = import.meta.env.VITE_BACKEND_ENDPOINT

// 標準地域コードのデータ型を定義
type stdAreaCode = {
  stdAreaCode: string;
  prefAreaCode: string;
  subPrefAreaCode: string;
  municAreaCode1: string;
  municAreaCode2: string;
  prefName: string;
  subPrefName: string;
  municName1: string;
  municName2: string;
  createdAt: string;
  updatedAt: string;
}

const AdminStdAreaCode: React.FC = () => {

  const [sacList, setSacList] = useState<stdAreaCode[]>([]);

  useEffect(() => {

    const options1 = {
      method: "GET",
      url: backendEndpoint + "/api/admin/sac",
      headers: {
        "Content-Type": "application/json"
      }
    };

    const fetchStdAreaCodeList = async () => {
      try {
        const resp = await axios(options1);
        setSacList(resp.data);
        console.log('GETリクエストが成功しました', resp.data)
      } catch (err) {
        console.error("GETリクエスト失敗", err)
      }
    }

    fetchStdAreaCodeList();

  }, []);

  const handleClick = (e: React.MouseEvent<HTMLButtonElement>) => {
 
    // ページ遷移を防ぐ（デフォルトでは、フォーム送信ボタンを押すとページが遷移してしまう）
    e.preventDefault();

    const options2 = {
      method: "PUT",
      url: backendEndpoint + "/api/admin/sac",
      headers: {
        "Content-Type": "application/json"
      }
    };

    axios(options2)
      .then((resp) => console.log('PUTリクエストが成功しました', resp.data))
      .catch((err) => console.error("PUTリクエスト失敗", err));

  };

  return (
    <div>
      <Header />
      <h1 className="text-[min(4vw,30px)] flex py-4 items-center justify-center">
        標準地域コード管理
      </h1>
      <div className="max-w-7xl mx-auto my-8 bg-gradient-to-b from-red-50 to-white rounded-lg shadow-lg overflow-hidden">
        <header className="bg-zinc-800 text-white p-4 text-center">
          <h2 className="text-2xl font-bold">標準地域コード一覧</h2>
        </header>
        <div className="overflow-x-auto">
          <table className="w-full" aria-label="標準地域コード一覧">
            <thead>
              <tr className="bg-zinc-800 text-white">
                <th className="p-3 text-left">標準地域コード</th>
                <th className="p-3 text-left">都道府県</th>
                <th className="p-3 text-left">振興局・支庁</th>
                <th className="p-3 text-left">市郡</th>
                <th className="p-3 text-left">区町村</th>
                <th className="p-3 text-left">都道府県名</th>
                <th className="p-3 text-left">振興局・支庁名</th>
                <th className="p-3 text-left">市郡名</th>
                <th className="p-3 text-left">区町村名</th>
                <th className="p-3 text-left">作成日時</th>
                <th className="p-3 text-left">更新日時</th>
              </tr>
            </thead>
            <tbody>
              {sacList.map((item, index) => (
                <tr key={item.stdAreaCode} className={index % 2 === 0 ? 'bg-gray-100' : 'bg-white'}>
                  <td className="p-3 border-b">{item.stdAreaCode}</td>
                  <td className="p-3 border-b">{item.prefAreaCode}</td>
                  <td className="p-3 border-b">{item.subPrefAreaCode}</td>
                  <td className="p-3 border-b">{item.municAreaCode1}</td>
                  <td className="p-3 border-b">{item.municAreaCode2}</td>
                  <td className="p-3 border-b">{item.prefName}</td>
                  <td className="p-3 border-b">{item.subPrefName}</td>
                  <td className="p-3 border-b">{item.municName1}</td>
                  <td className="p-3 border-b">{item.municName2}</td>
                  <td className="p-3 border-b">{item.createdAt}</td>
                  <td className="p-3 border-b">{item.updatedAt}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
        <footer className="bg-zinc-800 text-white p-2 text-center text-sm">
          <p>&copy; 2024 標準地域コード一覧. All rights reserved.</p>
        </footer>
      </div>
      <Button onClick={handleClick} className="w-full bg-red-900 hover:bg-red-800 text-white font-bold py-2 px-4 rounded-md transition duration-300 ease-in-out transform hover:scale-105 font-serif">
      最新化
      </Button>
    </div>
  );

};

export default AdminStdAreaCode;