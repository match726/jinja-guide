import { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import { ExternalLink } from 'lucide-react';
import axios from 'axios';

import { Header } from '@/components/ui/header';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';

import '@/styles/global.css';

const BACKEND_ENDPOINT=import.meta.env.VITE_BACKEND_ENDPOINT;

type ShrineDetails = {
  name: string
  furigana: string
  alt_name: string
  address: string
  image: string
  description: string
  tags: string[]
  founded_year: string
  object_of_worship: string[]
  shrine_rank: string[]
  has_goshuin: boolean
  website_url: string
  wikipedia_url: string
};

const ShrineInfo = () => {

  const [shrDetails, setShrDetails] = useState<ShrineDetails>({name: "", furigana: "", alt_name: "", address: "", image: "", description: "", tags: [""], founded_year: "", object_of_worship: [""], shrine_rank: [""], has_goshuin: false, website_url: "", wikipedia_url: ""});
  const search = useLocation().search;
  const query = new URLSearchParams(search);
  const payload = {plus_code: query.get('code')};

  useEffect(() => {

    const reqOptions = {
      method: "GET",
      url: BACKEND_ENDPOINT + "/api/shrine",
      headers: {
        "Content-Type": "application/json",
        "ShrGuide-Shrines-Authorization": JSON.stringify(payload),
      }
    };

    const fetchShrineInfo = async () => {
      try {
        const resp = await axios(reqOptions);
        setShrDetails(resp.data);
      } catch (error) {
        console.error("GETリクエスト失敗", error);
      }
    };

    fetchShrineInfo();

  }, []);

  return (
    <>
      <Header />
      <div className="min-h-screen bg-stone-100 flex items-center justify-center p-4">
        <Card className="w-full max-w-4xl bg-white shadow-lg rounded-lg overflow-hidden border-2 border-red-800">
          <CardHeader className="bg-red-800 text-white p-4">
            <h2 className="text-2xl font-bold text-center">{shrDetails.name}</h2>
            <p className="text-center text-gray-200">{shrDetails.furigana}</p>
          </CardHeader>
          <img src={shrDetails.image} alt={shrDetails.name} className="w-full h-80 object-cover" />
          <CardContent className="p-6">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <h3 className="text-lg font-semibold mb-2">別名称</h3>
                <p>{shrDetails.alt_name}</p>
              </div>
              <div>
                <h3 className="text-lg font-semibold mb-2">所在地</h3>
                <p>{shrDetails.address}</p>
              </div>
              <div>
                <h3 className="text-lg font-semibold mb-2">創建年</h3>
                <p>{shrDetails.founded_year}</p>
              </div>
              <div>
                <h3 className="text-lg font-semibold mb-2">社格</h3>
                  {shrDetails.shrine_rank.map((item, index) => (
                    <li key={index}>{item}</li>
                  ))}
              </div>
              <div>
                <h3 className="text-lg font-semibold mb-2">御祭神</h3>
                <ul className="list-disc list-inside">
                  {shrDetails.object_of_worship.map((deity, index) => (
                    <li key={index}>{deity}</li>
                  ))}
                </ul>
              </div>
              <div>
                <h3 className="text-lg font-semibold mb-2">御朱印</h3>
                <p>{shrDetails.has_goshuin ? "あり" : "なし"}</p>
              </div>
              <div className="md:col-span-2">
                <h3 className="text-lg font-semibold mb-2">説明</h3>
                <p className="text-gray-700">{shrDetails.description}</p>
              </div>
              <div className="md:col-span-2">
                <h3 className="text-lg font-semibold mb-2">関連ワード</h3>
                <div className="flex flex-wrap gap-2">
                  {shrDetails.tags.map((item, index) => (
                    <Badge key={index} variant="secondary">{item}</Badge>
                  ))}
                </div>
              </div>
              <div>
                <h3 className="text-lg font-semibold mb-2">公式HP</h3>
                <Button variant="link" className="p-0">
                  <a href={shrDetails.website_url} target="_blank" rel="noopener noreferrer" className="flex items-center">
                    公式サイトへ <ExternalLink className="ml-1 h-4 w-4" />
                  </a>
                </Button>
              </div>
              <div>
                <h3 className="text-lg font-semibold mb-2">Wikipedia</h3>
                <Button variant="link" className="p-0">
                  <a href={shrDetails.wikipedia_url} target="_blank" rel="noopener noreferrer" className="flex items-center">
                    Wikipediaへ <ExternalLink className="ml-1 h-4 w-4" />
                  </a>
                </Button>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </>
  );

};

export default ShrineInfo;