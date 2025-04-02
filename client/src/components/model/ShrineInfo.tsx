import { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import { ExternalLink } from 'lucide-react';
import axios from 'axios';

import { FRONTEND_URL, BACKEND_ENDPOINT } from '@/config/config';
import { FieldProps } from '@/types/types';
import { Header } from '@/components/ui/header';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';

import '@/styles/global.css';

type ShrineDetails = {
  name: string
  furigana: string
  image: string
  altName: FieldProps[]
  address: string
  placeId: string
  description: string
  tags: FieldProps[]
  foundedYear: FieldProps
  objectOfWorship: FieldProps[]
  shrineRank: FieldProps[]
  hasGoshuin: boolean
  websiteUrl: string
  wikipediaUrl: string
};

const ShrineInfo = () => {

  const [shrDetails, setShrDetails] = useState<ShrineDetails>({name: "", furigana: "", altName: [], address: "", placeId: "", image: "", description: "", tags: [], foundedYear: {id: 0, seq: "", content1: "", content2: "", content3: ""}, objectOfWorship: [], shrineRank: [], hasGoshuin: false, websiteUrl: "", wikipediaUrl: ""});
  const search = useLocation().search;
  // プラス記号が空白として解釈されるため、置換する
  const query = new URLSearchParams(search.replace("+", "%2B"));
  const payload = {plusCode: query.get('code')};

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
        console.log("HTTPレスポンス: ", resp.data)
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
        <Card className="w-full max-w-4xl bg-white shadow-lg rounded-lg overflow-hidden border-2 border-red-900">
          <CardHeader className="bg-red-900 text-white p-4">
            <h2 className="text-2xl font-bold text-center">{shrDetails.name}</h2>
            <p className="text-center text-gray-200">{shrDetails.furigana}</p>
          </CardHeader>
          <img src={shrDetails.image} alt={shrDetails.name} className="aspect-auto w-full object-cover" />
          <CardContent className="p-6">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <h3 className="text-lg font-semibold mb-2">別名称</h3>
                  {shrDetails.altName.map((item, index) => (
                    <li key={index}>{item.content1}</li>
                  ))}
              </div>
              <div>
                <h3 className="text-lg font-semibold mb-2">所在地</h3>
                <a href={"https://www.google.com/maps/search/?api=1&query=" + shrDetails.name + "&query_place_id=" + shrDetails.placeId}
                  className="text-blue-600 hover:underline focus:outline-none focus:ring-2 focus:ring-blue-500 rounded px-1"
                  target="_blank"
                >
                  {shrDetails.address}
                </a>
              </div>
              <div>
                <h3 className="text-lg font-semibold mb-2">創建年</h3>
                {shrDetails.foundedYear.content2 === "伝"
                 ? <p>"（" + {shrDetails.foundedYear.content2} + "）" + {shrDetails.foundedYear.content1}</p>
                 : <p>{shrDetails.foundedYear.content1}</p>
                }
              </div>
              <div>
                <h3 className="text-lg font-semibold mb-2">社格</h3>
                  {shrDetails.shrineRank.map((item, index) => (
                    <li key={index}>{item.content1}</li>
                  ))}
              </div>
              <div>
                <h3 className="text-lg font-semibold mb-2">御祭神</h3>
                <ul className="list-disc list-inside">
                  {shrDetails.objectOfWorship.map((deity, index) => (
                    <li key={index}>{deity.content1}</li>
                  ))}
                </ul>
              </div>
              <div>
                <h3 className="text-lg font-semibold mb-2">御朱印</h3>
                <p>{shrDetails.hasGoshuin ? "あり" : "なし"}</p>
              </div>
              <div className="md:col-span-2">
                <h3 className="text-lg font-semibold mb-2">説明</h3>
                <p className="text-gray-700">{shrDetails.description}</p>
              </div>
              <div className="md:col-span-2">
                <h3 className="text-lg font-semibold mb-2">関連ワード</h3>
                <div className="flex flex-wrap gap-2">
                  {shrDetails.tags.map((item, index) => (
                    item.content1 === "登録なし"
                      ? <Badge key={index} variant="secondary">
                          {item.content1}
                        </Badge>
                      : <Badge key={index} variant="secondary" className="cursor-pointer hover:bg-primary/80">
                          <a href={FRONTEND_URL + "/shrines/tag?tag=" + encodeURIComponent(item.content1)} rel="noopener noreferrer" className="flex items-center">
                            {item.content1}
                          </a>
                        </Badge>
                  ))}
                </div>
              </div>
              <div>
                <h3 className="text-lg font-semibold mb-2">公式HP</h3>
                {shrDetails.websiteUrl === ""
                  ? <Button variant="link" className="p-0 cursor-not-allowed">
                      登録なし
                    </Button>
                  : <Button variant="link" className="p-0">
                      <a href={shrDetails.websiteUrl} target="_blank" rel="noopener noreferrer" className="flex items-center">
                        公式サイトへ <ExternalLink className="ml-1 h-4 w-4" />
                      </a>
                    </Button>
                }
              </div>
              <div>
                <h3 className="text-lg font-semibold mb-2">Wikipedia</h3>
                {shrDetails.wikipediaUrl === ""
                  ? <Button variant="link" className="p-0 cursor-not-allowed">
                      登録なし
                    </Button>
                  : <Button variant="link" className="p-0">
                      <a href={shrDetails.wikipediaUrl} target="_blank" rel="noopener noreferrer" className="flex items-center">
                        Wikipediaへ <ExternalLink className="ml-1 h-4 w-4" />
                      </a>
                    </Button>
                }
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </>
  );

};

export default ShrineInfo;