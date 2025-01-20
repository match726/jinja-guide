import { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import { Stamp } from 'lucide-react';
import axios from 'axios';

import { Header } from '@/components/ui/header';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';

import '@/styles/global.css';

const frontendUrl = import.meta.env.VITE_FRONTEND_URL;
const backendEndpoint = import.meta.env.VITE_BACKEND_ENDPOINT;

// 神社一覧の型定義
type Shrine = {
  name: string
  address: string
  plusCode: string
  placeId: string
  objectOfWorship: string[]
  hasGoshuin: boolean
};

const ShrineTagList = () => {

  const [shrines, setShrines] = useState<Shrine[]>([]);
  const search = useLocation().search;
  const query = new URLSearchParams(search);
  const payload = {tags: query.get('tags')};

  useEffect(() => {

    const reqOptions = {
      method: "GET",
      url: backendEndpoint + "/api/shrines/tag",
      headers: {
        "Content-Type": "application/json; charset=utf-8",
        "ShrGuide-Shrines-Authorization": JSON.stringify(payload),
      }
    };

    const fetchShrinesInfo = async () => {
      try {
        const resp = await axios(reqOptions);
        setShrines(resp.data);
      } catch (error) {
        console.error("GETリクエスト失敗", error);
      }
    };

    fetchShrinesInfo();

  }, []);

  return (
    <>
      <Header />
      <div className="container mx-auto p-4 max-w-4xl">
        <div className="bg-red-900 text-white p-4 rounded-t-lg shadow-lg">
          <h2 className="text-2xl font-bold text-center">神社一覧</h2>
        </div>
        <div className="overflow-x-auto">
          <Table className="w-full">
            <TableHeader>
              <TableRow className="bg-gray-200 text-gray-800">
                <TableHead className="text-left">名称</TableHead>
                <TableHead className="text-left">住所</TableHead>
                <TableHead className="text-left">主祭神</TableHead>
                <TableHead className="text-center">御朱印</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {shrines.map((shrine: Shrine, index) => (
                <TableRow 
                  key={shrine.plusCode}
                  className={`${index % 2 === 0 ? 'bg-red-50' : 'bg-white'} hover:bg-red-100 transition-colors`}
                >
                  <TableCell className="font-medium">
                    <a href={frontendUrl + "/shrine?code=" + shrine.plusCode}
                      className="text-blue-600 hover:underline focus:outline-none focus:ring-2 focus:ring-blue-500 rounded px-1"
                      target="_blank"
                    >
                      {shrine.name}
                    </a>
                  </TableCell>
                  <TableCell>
                    <a href={"https://www.google.com/maps/search/?api=1&query=" + shrine.name + "&query_place_id=" + shrine.placeId}
                      className="text-blue-600 hover:underline focus:outline-none focus:ring-2 focus:ring-blue-500 rounded px-1"
                      target="_blank"
                    >
                      {shrine.address}
                    </a>
                  </TableCell>
                  <TableCell>{shrine.objectOfWorship}</TableCell>
                  <TableCell className="text-center">
                    {shrine.hasGoshuin && (
                      <Stamp className="inline-block text-red-800" size={24} />
                    )}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </div>
        <div className="bg-red-900 h-4 rounded-b-lg shadow-lg" />
      </div>
    </>
  )
};

export default ShrineTagList;