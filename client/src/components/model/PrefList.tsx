import { useState, useEffect } from "react";
import axios from 'axios';

import { Header } from '@/components/ui/header';

import '@/styles/global.css';

const BACKEND_ENDPOINT=import.meta.env.VITE_BACKEND_ENDPOINT;

const Prefs = () => {

  const [hrchy, setHrchy] = useState([]);

  useEffect(() => {

    const options = {
      method: "GET",
      url: BACKEND_ENDPOINT + "/api/prefs",
      headers: {
        "Content-Type": "application/json"
      }
    };

    const fetchHrchyInfo = async () => {
      try {
        const resp = await axios(options);
        setHrchy(resp.data);
      } catch (error) {
        console.error("GETリクエスト失敗", error);
      }
    };

    fetchHrchyInfo();

  }, []);

  console.log(hrchy);

  return (
    <>
      <Header />
    </>
  );

}

export default Prefs;