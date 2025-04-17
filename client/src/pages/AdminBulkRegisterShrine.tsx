import { useState, useEffect } from 'react';
import axios from 'axios';

import { BACKEND_ENDPOINT } from '@/config/config';
import { Header } from '@/components/ui/header';
import { Button } from '@/components/ui/button';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { RegisterShrineProps } from '@/features/bulk-register/types/types';

import '@/styles/global.css';

const AdminBulkRegisterShrine = () => {

  //const [submit, setSubmit] = useState(false);
  const [shrines, setShrines] = useState<RegisterShrineProps[]>([]);

  // ページネーション関連
  const [currentPage, setCurrentPage] = useState(1)
  const itemsPerPage = 10
  const totalPages = Math.ceil(shrines.length / itemsPerPage)

  const startIndex = (currentPage - 1) * itemsPerPage
  const endIndex = startIndex + itemsPerPage
  const currentShrines = shrines.slice(startIndex, endIndex)

  useEffect(() => {

    const options = {
      method: "GET",
      url: BACKEND_ENDPOINT + "/api/admin/bulk-regist/shrine",
      headers: {
        "Content-Type": "application/json"
      }
    };

    const fetchShrinesList = async () => {
      try {
        const resp = await axios(options);
        setShrines(resp.data);
      } catch (error) {
        console.error("GETリクエスト失敗", error);
      }
    };

    fetchShrinesList();

  }, []);

  const handleFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {

    // ページ遷移を防ぐ（デフォルトでは、フォーム送信ボタンを押すとページが遷移してしまう）
    e.preventDefault();

  };

  console.log(shrines);

  return (
    <>
      <Header />
      <div className="bg-gradient-to-b from-red-50 to-white flex items-top justify-center p-8">
        <div className="w-full bg-white rounded-lg shadow-xl overflow-hidden">
          <div className="bg-red-900 p-4 flex items-center justify-center">
            <h2 className="text-2xl font-bold text-white ml-2 font-serif">神社一括登録</h2>
          </div>
          <form onSubmit={handleFormSubmit} className="p-6 space-y-6">
            <div className="container mx-auto p-4 max-w-4xl">
              <div className="bg-red-900 text-white p-4 rounded-t-lg shadow-lg">
                <h3 className="text-xl font-bold text-center">神社一括登録テーブル</h3>
              </div>
              <div className="overflow-x-auto">
                <Table className="w-full">
                  <TableHeader>
                    <TableRow className="bg-gray-200 text-gray-800">
                      <TableHead className="text-left">名称</TableHead>
                      <TableHead className="text-left">住所</TableHead>
                      <TableHead className="text-left">振り仮名</TableHead>
                      <TableHead className="text-left">別名称</TableHead>
                      <TableHead className="text-left">関連ワード</TableHead>
                      <TableHead className="text-left">創建年</TableHead>
                      <TableHead className="text-left">御祭神</TableHead>
                      <TableHead className="text-left">社格</TableHead>
                      <TableHead className="text-left">御朱印</TableHead>
                      <TableHead className="text-left">公式サイトURL</TableHead>                    
                      <TableHead className="text-left">WikipediaURL</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {currentShrines.map((shrine: RegisterShrineProps, index) => (
                      <TableRow
                        key={index}
                        className={`${index % 2 === 0 ? 'bg-red-50' : 'bg-white'} hover:bg-red-100 transition-colors`}
                      >
                        <TableCell className="font-medium">{shrine.name}</TableCell>
                        <TableCell>{shrine.address}</TableCell>
                        <TableCell>{shrine.furigana}</TableCell>
                        <TableCell>
                          {typeof shrine.altName === 'undefined'
                            ? null
                            : shrine.altName.join(",")
                          }
                        </TableCell>
                        <TableCell>
                          {typeof shrine.tags === 'undefined'
                            ? null
                            : shrine.tags.join(",")
                          }
                        </TableCell>
                        <TableCell>
                          {typeof shrine.foundedYear === 'undefined'
                            ? null
                            : shrine.foundedYear.join(",")
                          }
                        </TableCell>
                        <TableCell>
                          {typeof shrine.objectOfWorship === 'undefined'
                            ? null
                            : shrine.objectOfWorship.join(",")
                          }
                        </TableCell>
                        <TableCell>
                          {typeof shrine.shrineRank === 'undefined'
                            ? null
                            : shrine.shrineRank.join(",")
                          }
                        </TableCell>
                        <TableCell>{shrine.hasGoshuin}</TableCell>
                        <TableCell>{shrine.websiteUrl}</TableCell>
                        <TableCell>{shrine.wikipediaUrl}</TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </div>
              <div className="bg-red-50 flex justify-between items-center">
                <Button
                  onClick={() => setCurrentPage((prev) => Math.max(prev - 1, 1))}
                  disabled={currentPage === 1}
                  className="bg-red-900 text-white hover:bg-red-700"
                >
                  前へ
                </Button>
                <span className="text-white bg-red-900 px-4 py-2 rounded-t-lg font-semibold">
                  {currentPage} / {totalPages} 頁 ({shrines.length})
                </span>
                <Button
                  onClick={() => setCurrentPage((prev) => Math.min(prev + 1, totalPages))}
                  disabled={currentPage === totalPages}
                  className="bg-red-900 text-white hover:bg-red-700"
                >
                  次へ
                </Button>
              </div>
              <div className="bg-red-900 h-4 rounded-b-lg shadow-lg" />
            </div>
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