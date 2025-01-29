import { useState, useEffect, useRef } from 'react';
import axios from 'axios';

import { Header } from '@/components/ui/header';
import { Button } from '@/components/ui/button';
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle} from '@/components/ui/dialog';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';

import '@/styles/global.css';

const backendEndpoint = import.meta.env.VITE_BACKEND_ENDPOINT;

const AdminRegisterShrine = () => {

  const [payload, setPayload] = useState({name: "", furigana: "", address: "", wikipediaUrl: ""});
  const [respStatus, setRespStatus] = useState(0);
  const [dialogOpen, setDialogOpen] = useState(false);
  // 初回レンダリングのリクエスト送信を無効化
  const isFirstRender = useRef(true);

  useEffect(() => {

    const options = {
      method: "POST",
      url: backendEndpoint + "/api/admin/regist/shrine",
      headers: {
        "Content-Type": "application/json"
      },
      data: JSON.stringify(payload)
    };

    if (isFirstRender.current) {
      isFirstRender.current = false;
      return
    } else {
      axios(options)
        .then((resp) => {
          setRespStatus(resp.status)
          setDialogOpen(true)
          console.log('POSTリクエストが成功しました', resp)
        })
        .catch((err) => console.error("POSTリクエスト失敗", err));
    }

  }, [payload]);

  const handleFromSubmit = (e: React.FormEvent<HTMLFormElement>) => {

    // ページ遷移を防ぐ（デフォルトでは、フォーム送信ボタンを押すとページが遷移してしまう）
    e.preventDefault();

    const form = e.currentTarget
    const formElements = form.elements as typeof form.elements & {
      shrineName: HTMLInputElement,
      shrineFurigana: HTMLInputElement,
      shrineAddress: HTMLInputElement,
      shrineWikiUrl: HTMLInputElement
    }
    setPayload({ name: formElements.shrineName.value, furigana: formElements.shrineFurigana.value, address: formElements.shrineAddress.value, wikipediaUrl: formElements.shrineWikiUrl.value });

    // フォームをクリア
    formElements.shrineName.value = "";
    formElements.shrineFurigana.value = "";
    formElements.shrineAddress.value = "";
    formElements.shrineWikiUrl.value = "";

  };

  return (
    <>
      <Header />
      <div className="bg-gradient-to-b from-red-50 to-white flex items-top justify-center p-8">
        <div className="w-full max-w-md bg-white rounded-lg shadow-xl overflow-hidden">
          <div className="bg-red-900 p-4 flex items-center justify-center">
            <h2 className="text-2xl font-bold text-white ml-2 font-serif">神社登録</h2>
          </div>
          <form onSubmit={handleFromSubmit} className="p-6 space-y-6">
            <div className="space-y-2">
              <Label htmlFor="shrineName" className="text-lg font-medium text-gray-700 font-serif">
                神社名称
              </Label>
              <Input
                id="shrineName"
                type="text"
                placeholder="例：伊勢神宮 内宮（皇大神宮）"
                className="w-full border-2 border-red-800 rounded-md p-2 font-serif"
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="shrineFurigana" className="text-lg font-medium text-gray-700 font-serif">
                神社名称（振り仮名）
              </Label>
              <Input
                id="shrineFurigana"
                type="text"
                placeholder="例：いせじんぐう ないくう（こうたいじんぐう）"
                className="w-full border-2 border-red-800 rounded-md p-2 font-serif"
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="shrineAddress" className="text-lg font-medium text-gray-700 font-serif">
                住所
              </Label>
              <Input
                id="shrineAddress"
                name="address"
                placeholder="例：三重県伊勢市宇治館町１"
                className="w-full border-2 border-red-900 rounded-md p-2 font-serif"
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="shrineWikiUrl" className="text-lg font-medium text-gray-700 font-serif">
                WikipediaURL
              </Label>
              <Input
                id="shrineWikiUrl"
                type="text"
                placeholder="例：https://ja.wikipedia.org/wiki/伊勢神宮"
                className="w-full border-2 border-red-800 rounded-md p-2 font-serif"
              />
            </div>
            <Button className="w-full bg-red-900 hover:bg-red-800 text-white font-bold py-2 px-4 rounded-md transition duration-300 ease-in-out transform hover:scale-105 font-serif">
              登録
            </Button>
          </form>
        </div>
        <Dialog open={dialogOpen}>
          <DialogContent className="sm:max-w-[425px] bg-red-50 border-2 border-red-900 rounded-none">
            <DialogHeader className="relative">
              <DialogTitle className="text-2xl font-bold text-red-900">登録結果</DialogTitle>
              <DialogDescription>神社情報の登録結果です</DialogDescription>
            </DialogHeader>
            <div className="bg-white bg-opacity-70 p-4 rounded shadow-inner">{respStatus}</div>
          </DialogContent>
        </Dialog>
      </div>
    </>

  );
  
};

export default AdminRegisterShrine;