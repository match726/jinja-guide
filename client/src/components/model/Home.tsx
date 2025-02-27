import { useState, useEffect } from 'react';
import axios from 'axios';

import { Header } from '@/components/ui/header';
import { AdminCard } from '@/components/ui/card/admin-card';

import '@/styles/global.css';

const backendEndpoint = import.meta.env.VITE_BACKEND_ENDPOINT;

type RandomShrines = {
  name: string,
  furigana: string,
  address: string,
  plusCode: string,
  placeId: string,
  objectOfWorship: string[],
  description: string
}

const frontendUrl = import.meta.env.VITE_FRONTEND_URL;

const Home: React.FC = () => {

  const [shrines, setShrines] = useState<RandomShrines[]>();

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
        setShrines(resp.data.RandomShrine);
      } catch (error) {
        console.error("GETリクエスト失敗", error);
      }
    };

    fetchHomeContents();

  }, []);

  console.log(shrines)

  return (
    <div>
      <Header />
      <div className="bg-gradient-to-b from-red-50 to-white">
        <h1 className="text-[min(4vw,30px)] flex py-4 items-center justify-center">
          神社（ランダム表示）
        </h1>
        <section className="container flex grid grid-cols-2 gap-10 xl:grid-cols-3">
          {shrines && shrines.map((data) => (
            <AdminCard
              cardTitle={data.name}
              cardTitleRuby={data.furigana}
              cardDescription={data.description}
              cardLink={frontendUrl + "/shrine?code=" + data.plusCode}
            />         
          ))}
        </section>
      </div>
    </div>
  );
};

export default Home;