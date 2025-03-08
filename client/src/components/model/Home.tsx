import { useState, useEffect } from 'react';
import axios from 'axios';

import { Header } from '@/components/ui/header';
//import { Badge } from '@/components/ui/badge';
//import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
//import { ShrineCard } from '@/components/ui/card/shrine-card';
import { Tabs, TabsList, TabsTrigger, TabsContent } from '@/components/ui/tabs';


import '@/styles/global.css';

//const frontendUrl = import.meta.env.VITE_FRONTEND_URL;
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
  count: number
}

type HomeContents = {
  shrines: RandomShrines[],
  tags: AllTags[]
}

const Home: React.FC = () => {

  const [activeTab, setActiveTab] = useState<string>("0")
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

  console.log(contents)

  return (
    <div>
      <Header />
      <div className="relative h-screen w-full">
        <div
          className="absolute inset-0 bg-cover bg-center bg-no-repeat"
          style={{
            backgroundImage: "url('https://nhrje5lnk6nkethb.public.blob.vercel-storage.com/top.jpg')",
            backgroundPosition: "center",
          }}
        />
        <div className="bg-white bg-opacity-0 px-6 py-4">
          <Tabs value={activeTab} onValueChange={setActiveTab} className="border-b">
            <TabsList className="relative max-w-md mx-auto flex">
              <TabsTrigger value="0" className="flex-1 py-3 text-center transition-colors duration-300 hover:bg-muted/50">
                神社検索
              </TabsTrigger>
              <TabsTrigger value="1" className="flex-1 py-3 text-center transition-colors duration-300 hover:bg-muted/50">
                ランダム神社
              </TabsTrigger>
              <TabsTrigger value="2" className="flex-1 py-3 text-center transition-colors duration-300 hover:bg-muted/50">
                関連ワード
              </TabsTrigger>
            </TabsList>
            <TabsContent value="0" className="py-6">
              <div className="flex justify-center">
                <div className="relative w-full max-w-md">
                  <Input
                    type="search"
                    placeholder="Search..."
                    className="w-full rounded-full bg-background pl-4 pr-12 py-2 text-foreground focus:outline-none focus:ring-1 focus:ring-primary"
                  />
                  <div className="absolute inset-y-0 right-0 flex items-center pr-4">
                    <SearchIcon className="h-5 w-5 text-muted-foreground" />
                  </div>
                </div>
              </div>
            </TabsContent>
            <TabsContent value="1" className="py-6">
              <div className="space-y-4">
                <h3 className="text-lg font-medium">Tab 2 Content</h3>
                <p className="text-muted-foreground">
                  This is the content for Tab 2. It can contain any kind of information or components you need.
                </p>
              </div>
            </TabsContent>
            <TabsContent value="2" className="py-6">
              <div className="space-y-4">
                <h3 className="text-lg font-medium">Tab 3 Content</h3>
                <p className="text-muted-foreground">
                  This is the content for Tab 3. It can contain any kind of information or components you need.
                </p>
              </div>
            </TabsContent>
          </Tabs>
        </div>
      </div>
    </div>
  );
};

function SearchIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <circle cx="11" cy="11" r="8" />
      <path d="m21 21-4.3-4.3" />
    </svg>
  )
}

export default Home;