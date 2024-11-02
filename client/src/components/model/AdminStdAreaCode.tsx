import { useState, useEffect } from 'react';

import { Header } from '@/components/ui/header';

const BACKEND_ENDPOINT="https://jinja-guide-server.vercel.app";

const AdminStdAreaCode = () => {

  useEffect(() => {

    const reqOptions = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      }
    };

    fetch(`${BACKEND_ENDPOINT}/api/admin/sac`, reqOptions
    ).then((resp) => resp.json()
    ).then((respJson) => {
      console.log("Name: " + respJson.Name + ", Address: " + respJson.Address + ", StdAreaCode: " + respJson.StdAreaCode + ", OpenLocnCode: " + respJson.OpenLocnCode + ", PlaceID: " + respJson.PlaceID);
    }).catch(() => {
      console.log("error");
    });

  }, []);

  return (
    <div>
      <Header />
    </div>
  );

};

export default AdminStdAreaCode;