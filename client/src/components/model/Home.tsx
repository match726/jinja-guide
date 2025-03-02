import { useState, useEffect } from 'react';
import axios from 'axios';

import { Header } from '@/components/ui/header';
import { ShrineCard } from '@/components/ui/card/shrine-card';
import { Badge } from '@/components/ui/badge';

import '@/styles/global.css';

const backendEndpoint = import.meta.env.VITE_BACKEND_ENDPOINT;

interface RandomShrines {
  name: string
  furigana: string
  address: string
  plusCode: string
  placeId: string
  objectOfWorship: string[]
  description: string
}

interface AllTags {
  name: string
}

type HomeContents = {
  shrines: RandomShrines[],
  tags: AllTags[]
}

const frontendUrl = import.meta.env.VITE_FRONTEND_URL;

const Home: React.FC = () => {

  const [contents, setContents] = useState<HomeContents>({shrines: [], tags: []});

  useEffect(() => {

    const reqOptions = {
      method: "GET",
      url: backendEndpoint + "/api/home",
      headers: {
        "Content-Type": "application/json",
      }
    };

    const fetchHomeContents = async () => {
      try {
        const resp = await axios(reqOptions);
        console.log("HTTPレスポンス: ", resp.data)
        setContents(resp.data);
      } catch (error) {
        console.error("GETリクエスト失敗", error);
      }
    };

    fetchHomeContents();

  }, []);

  return (
    <div>
      <Header />
      <div className="bg-gradient-to-b from-red-50 to-white">
        <p className="text-[min(4vw,30px)] flex py-4 items-center justify-center">
          神社（ランダム表示）
        </p>
        <section className="container flex grid grid-cols-2 gap-10 xl:grid-cols-3">
          {contents && contents.shrines.map((data) => (
            <ShrineCard
              cardTitle={data.name}
              cardTitleRuby={data.furigana}
              cardAddress={data.address}
              cardObjectOfWorship={data.objectOfWorship}
              cardDescription={data.description}
              cardLink={frontendUrl + "/shrine?code=" + data.plusCode}
            />         
          ))}
        </section>
        <p className="text-[min(4vw,30px)] flex py-4 items-center justify-center">
          関連ワード一覧
        </p>
        <section className="container flex grid grid-cols-2 gap-10 xl:grid-cols-3">
          <div className="flex flex-wrap gap-2">
          {contents && contents.tags.map((item, index) => (
            <Badge key={index} variant="secondary" className="cursor-pointer hover:bg-primary/80">
              <a href={frontendUrl + "/shrines/tag?tag=" + encodeURIComponent(item.name)} rel="noopener noreferrer" className="flex items-center">
                {item.name}
              </a>
            </Badge>
          ))}
        </div>
        </section>
      </div>
    </div>
  );
};

export default Home;