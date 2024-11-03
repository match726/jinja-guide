import { useState, useEffect } from 'react';
import axios from 'axios';

import { Header } from '@/components/ui/header';

const BACKEND_ENDPOINT=import.meta.env.VITE_BACKEND_ENDPOINT

// 標準地域コードのデータ型を定義
type stdAreaCode = {
  StdAreaCode: string;
  PrefAreaCode: string;
  SubPrefAreaCode: string;
  MunicAreaCode1: string;
  MunicAreaCode2: string;
  PrefName: string;
  SubPrefName: string;
  MunicName1: string;
  MunicName2: string;
  CreatedAt: string;
  UpdatedAt: string;
}

const AdminStdAreaCode = () => {

  const [sacList, setSacList] = useState<stdAreaCode[]>([]);

  useEffect(() => {

    const fetchStdAreaCodeList = async () => {
      try {
        const resp = await axios.get(`${BACKEND_ENDPOINT}/api/admin/sac`, {
          headers: {
            "Content-Type": "application/json",
            "ShrGuide-Shrines-Authorization": "Test",
          },
        });
        setSacList(resp.data);
      } catch (err) {
        console.error("GETリクエスト失敗", err)
      }
    }

    fetchStdAreaCodeList();

  }, []);

  return (
    <div>
      <Header />
      <h1 className="text-[min(4vw,30px)] flex py-4 items-center justify-center">
        標準地域コード管理
      </h1>
      <div className="max-w-7xl mx-auto my-8 bg-white rounded-lg shadow-lg overflow-hidden">
        <header className="bg-zinc-800 text-white p-4 text-center">
          <h1 className="text-2xl font-bold">標準地域コード一覧</h1>
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
                <tr key={item.StdAreaCode} className={index % 2 === 0 ? 'bg-gray-100' : 'bg-white'}>
                  <td className="p-3 border-b">{item.StdAreaCode}</td>
                  <td className="p-3 border-b">{item.PrefAreaCode}</td>
                  <td className="p-3 border-b">{item.SubPrefAreaCode}</td>
                  <td className="p-3 border-b">{item.MunicAreaCode1}</td>
                  <td className="p-3 border-b">{item.MunicAreaCode2}</td>
                  <td className="p-3 border-b">{item.PrefName}</td>
                  <td className="p-3 border-b">{item.SubPrefName}</td>
                  <td className="p-3 border-b">{item.MunicName1}</td>
                  <td className="p-3 border-b">{item.MunicName2}</td>
                  <td className="p-3 border-b">{item.CreatedAt}</td>
                  <td className="p-3 border-b">{item.UpdatedAt}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
        <footer className="bg-zinc-800 text-white p-2 text-center text-sm">
          <p>&copy; 2024 標準地域コード一覧. All rights reserved.</p>
        </footer>
      </div>
    </div>
  );

};

export default AdminStdAreaCode;