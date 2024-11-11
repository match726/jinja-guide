import { useState, useEffect } from "react";
import axios from 'axios';

import { Header } from '@/components/ui/header';

import '@/styles/global.css';

type SacRelationship = {
  name: string
  std_area_code: string
  sup_std_area_code: string
  kinds: string
  has_child: boolean
};

//const FRONTEND_URL=import.meta.env.VITE_FRONTEND_URL;
const BACKEND_ENDPOINT=import.meta.env.VITE_BACKEND_ENDPOINT;

const Prefs = () => {

  const [sacr, setSacr] = useState<SacRelationship[]>([]);

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
        setSacr(resp.data);
      } catch (error) {
        console.error("GETリクエスト失敗", error);
      }
    };

    fetchSacRelationshipInfo();

  }, []);

  console.log(sacr);

  return (
    <>
      <Header />
    </>
  );

}

export default Prefs;