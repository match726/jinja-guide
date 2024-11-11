import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import { ChevronRight, ChevronDown } from 'lucide-react';

import { Header } from '@/components/ui/header';
import { Button } from '@/components/ui/button';
import { Collapsible, CollapsibleContent, CollapsibleTrigger, } from '@/components/ui/collapsible';

import '@/styles/global.css';

type SacRelationship = {
  name: string
  std_area_code: string
  sup_std_area_code: string
  kinds: string
  has_child: boolean
};

const FRONTEND_URL=import.meta.env.VITE_FRONTEND_URL + "/shrines";
const BACKEND_ENDPOINT=import.meta.env.VITE_BACKEND_ENDPOINT;

const Prefs = () => {

  const [sacrs, setSacrs] = useState<SacRelationship[]>([]);

  // 標準地域コードの関係性リストを取得
  useEffect(() => {

    const options = {
      method: "GET",
      url: BACKEND_ENDPOINT + "/api/prefs",
      headers: {
        "Content-Type": "application/json"
      }
    };

    const fetchSacRelationshipInfo = async () => {
      try {
        const resp = await axios(options);
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
          <Button asChild variant="ghost" className="w-full justify-start p-2 whitespace-nowrap">
            <>
              {sacr.has_child && (
                isOpen ? <ChevronDown className="mr-2 h-4 w-4" /> : <ChevronRight className="mr-2 h-4 w-4" />
              )}
              <Link to={FRONTEND_URL + "?kinds=" + sacr.kinds + "&code=" + sacr.std_area_code}>{sacr.name}</Link>
            </>
          </Button>
        </CollapsibleTrigger>
        {sacr.has_child && (
          <CollapsibleContent className="ml-4">
            {sacrs.filter((row: SacRelationship) => row.sup_std_area_code === sacr.std_area_code).map((elem: SacRelationship) => (
              <RenderPrefNode key={elem.std_area_code} sacr={elem} />
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
        エリアから検索
      </h1>
      <div className="w-full max-w-md mx-auto p-4 space-y-2 border rounded-lg shadow-sm">
        {sacrs.filter((row: SacRelationship) => row.kinds === "Pref").map((elem: SacRelationship) => (
          <RenderPrefNode key={elem.std_area_code} sacr={elem} />
        ))}
      </div>
    </>
  );

};

export default Prefs;