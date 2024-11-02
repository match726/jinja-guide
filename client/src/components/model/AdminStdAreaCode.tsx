import { useEffect } from 'react';

import { Header } from '@/components/ui/header';

const BACKEND_ENDPOINT=import.meta.env.VITE_BACKEND_ENDPOINT

// テーブルの行のデータ型を定義
type stdAreaCode = {
  stdAreaCode: string;
  prefAreaCode: string;
  subprefAreaCode: string;
  municAreaCode1: string;
  municAreaCode2: string;
  prefName: string;
  subprefName: string;
  municName1: string;
  municName2: string;
  createdAt: string;
  updatedAt: string;
}

// サンプルデータ
const sampleData: stdAreaCode[] = [
  {
    stdAreaCode: '01101',
    prefAreaCode: '01',
    subprefAreaCode: '-',
    municAreaCode1: '-',
    municAreaCode2: '101',
    prefName: '北海道',
    subprefName: '-',
    municName1: '-',
    municName2: '札幌市中央区',
    createdAt: '2023-05-01 10:00:00',
    updatedAt: '2023-05-01 10:00:00',
  },
  {
    stdAreaCode: '13101',
    prefAreaCode: '13',
    subprefAreaCode: '-',
    municAreaCode1: '-',
    municAreaCode2: '101',
    prefName: '東京都',
    subprefName: '-',
    municName1: '-',
    municName2: '千代田区',
    createdAt: '2023-05-01 10:00:00',
    updatedAt: '2023-05-01 10:00:00',
  },
];

const AdminStdAreaCode = () => {

  useEffect(() => {

    const reqOptions = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      }
    };

    fetch(`${BACKEND_ENDPOINT}/api/admin/sac`, reqOptions
    ).then((resp) => resp.json()
    ).then((respJson) => {
      console.log(respJson);
    }).catch(() => {
      console.log("error");
    });

  }, []);

  return (
    <div>
      <Header />
      <h1 className="text-[min(4vw,30px)] flex py-4 items-center justify-center">
        標準地域コード管理
      </h1>
      <div className="max-w-7xl mx-auto my-8 bg-white rounded-lg shadow-lg overflow-hidden">
        <header className="bg-gray-900 text-white p-4 text-center">
          <h1 className="text-2xl font-bold">標準地域コード</h1>
        </header>
        <div className="overflow-x-auto">
          <table className="w-full" aria-label="標準地域コード">
            <thead>
              <tr className="bg-red-900 text-white">
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
              {sampleData.map((item, index) => (
                <tr key={item.stdAreaCode} className={index % 2 === 0 ? 'bg-gray-100' : 'bg-white'}>
                  <td className="p-3 border-b">{item.stdAreaCode}</td>
                  <td className="p-3 border-b">{item.prefAreaCode}</td>
                  <td className="p-3 border-b">{item.subprefAreaCode}</td>
                  <td className="p-3 border-b">{item.municAreaCode1}</td>
                  <td className="p-3 border-b">{item.municAreaCode2}</td>
                  <td className="p-3 border-b">{item.prefName}</td>
                  <td className="p-3 border-b">{item.subprefName}</td>
                  <td className="p-3 border-b">{item.municName1}</td>
                  <td className="p-3 border-b">{item.municName2}</td>
                  <td className="p-3 border-b">{item.createdAt}</td>
                  <td className="p-3 border-b">{item.updatedAt}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
        <footer className="bg-gray-900 text-white p-2 text-center text-sm">
          <p>&copy; 2023 神社情報テーブル. All rights reserved.</p>
        </footer>
      </div>
    </div>
  );

};

export default AdminStdAreaCode;