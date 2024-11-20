import { useState, useEffect } from 'react';
import { ChevronRight, ChevronDown } from 'lucide-react';
import axios from 'axios';

import { Header } from '@/components/ui/header';
import { Button } from '@/components/ui/button';
import { Collapsible, CollapsibleContent, CollapsibleTrigger, } from '@/components/ui/collapsible';

import '@/styles/global.css';

type SacRelationship = {
  name: string
  stdAreaCode: string
  supStdAreaCode: string
  kinds: string
  hasChild: boolean
};

const linkUrl = import.meta.env.VITE_FRONTEND_URL + "/shrines";
const backendEndpoint = import.meta.env.VITE_BACKEND_ENDPOINT;

const Prefs = () => {

  const [sacrs, setSacrs] = useState<SacRelationship[]>([]);

  // 標準地域コードの関係性リストを取得
  useEffect(() => {

    const reqOptions = {
      method: "GET",
      url: backendEndpoint + "/api/prefs",
      headers: {
        "Content-Type": "application/json"
      }
    };

    const fetchSacRelationshipInfo = async () => {
      try {
        const resp = await axios(reqOptions);
        setSacrs(resp.data);
      } catch (error) {
        console.error("GETリクエスト失敗", error);
      }
    };

    fetchSacRelationshipInfo();

  }, []);

  function RenderPrefNode({ sacr }: {sacr: SacRelationship}) {

    const [isOpen, setIsOpen] = useState(false);

    return (
      <Collapsible open={isOpen} onOpenChange={setIsOpen}>
        <CollapsibleTrigger asChild>
        <Button variant="ghost" className="w-full justify-start p-2 hover:bg-transparent">
          {isOpen ? <ChevronDown className="mr-2 h-4 w-4" /> : <ChevronRight className="mr-2 h-4 w-4" />}
          <a href={linkUrl + "?kinds=" + sacr.kinds + "&code=" + sacr.stdAreaCode}
            className="text-blue-600 hover:underline focus:outline-none focus:ring-2 focus:ring-blue-500 rounded px-1"
            onClick={(e) => e.stopPropagation()}
          >
            {sacr.name}
          </a>
        </Button>
        </CollapsibleTrigger>
        {sacr.hasChild && (
          <CollapsibleContent className="ml-4">
            {sacrs.filter((row: SacRelationship) => row.supStdAreaCode === sacr.stdAreaCode).map((elem: SacRelationship) => (
              <RenderPrefNode key={elem.stdAreaCode} sacr={elem} />
            ))}
          </CollapsibleContent>
        )}
      </Collapsible>
    )
  };

  return (
    <>
      <Header />
      <h1 className="text-[min(4vw,30px)] flex py-4 items-center justify-center">
        都道府県／市区町村検索
      </h1>
      <div className="w-full max-w-md mx-auto p-4 space-y-2 border rounded-lg shadow-sm">
        {sacrs.filter((row: SacRelationship) => row.kinds === "Pref").map((elem: SacRelationship) => (
          <RenderPrefNode key={elem.stdAreaCode} sacr={elem} />
        ))}
      </div>
    </>
  );

};

export default Prefs;