import { useState, useEffect, useRef } from "react";
import axios from 'axios';

import { Header } from '@/components/ui/header';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';

import '@/styles/global.css';

const backendEndpoint = import.meta.env.VITE_BACKEND_ENDPOINT;

const AdminRegisterShrineDetails = () => {

  const [payload, setPayload] = useState({plusCode: "", furigana: "", altName: "", tag: "", foundedYear: "", hasGoshuin: "", objectOfWorship: "", websiteURL: "", wikipediaUrl: ""});
  // 初回レンダリングのリクエスト送信を無効化
  const isFirstRender = useRef(true);

  useEffect(() => {

    const options = {
      method: "POST",
      url: backendEndpoint + "/api/admin/regist/details",
      headers: {
        "Content-Type": "application/json"
      },
      data: JSON.stringify(payload)
    };

    console.log(payload)
    if (isFirstRender.current) {
      isFirstRender.current = false;
      return
    } else {
      axios(options)
        .then((resp) => {
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
      shrinePlusCode: HTMLInputElement,
      shrineFurigana: HTMLInputElement,
      shrineAltName: HTMLInputElement,
      shrineTag: HTMLInputElement,
      shrineFoundedYear: HTMLInputElement,
      shrineHasGoshuin: HTMLInputElement,
      shrineObjectOfWorship: HTMLInputElement,
      shrineWebsiteURL: HTMLInputElement,
      shrineWikiUrl: HTMLInputElement
    }
    setPayload({ plusCode: formElements.shrinePlusCode.value, furigana: formElements.shrineFurigana.value, altName: formElements.shrineAltName.value, tag: formElements.shrineTag.value, foundedYear: formElements.shrineFoundedYear.value, hasGoshuin: formElements.shrineHasGoshuin.value, objectOfWorship: formElements.shrineObjectOfWorship.value, websiteURL: formElements.shrineWebsiteURL.value, wikipediaUrl: formElements.shrineWikiUrl.value });

    // フォームをクリア
    formElements.shrinePlusCode.value = "";
    formElements.shrineFurigana.value = "";
    formElements.shrineAltName.value = "";
    formElements.shrineTag.value = "";
    formElements.shrineFoundedYear.value = "";
    formElements.shrineHasGoshuin.value = "";
    formElements.shrineObjectOfWorship.value = "";
    formElements.shrineWebsiteURL.value = "";
    formElements.shrineWikiUrl.value = "";

  };

  return (
    <>
      <Header />
      <div className="bg-gradient-to-b from-red-50 to-white flex items-top justify-center p-8">
        <div className="w-full max-w-md bg-white rounded-lg shadow-xl overflow-hidden">
          <div className="bg-red-900 p-4 flex items-center justify-center">
            <h2 className="text-2xl font-bold text-white ml-2 font-serif">神社詳細情報登録</h2>
          </div>
          <form onSubmit={handleFromSubmit} className="p-6 space-y-6">
            <div className="space-y-2">
              <Label htmlFor="shrinePlusCode" className="text-lg font-medium text-gray-700 font-serif">
                PlusCode
              </Label>
              <Input
                id="shrinePlusCode"
                type="text"
                placeholder="例：8Q6RFP4G+255"
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
              <Label htmlFor="shrineAltName" className="text-lg font-medium text-gray-700 font-serif">
                別名称
              </Label>
              <Input
                id="shrineAltName"
                type="text"
                placeholder="例：伊勢神宮"
                className="w-full border-2 border-red-800 rounded-md p-2 font-serif"
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="shrineTag" className="text-lg font-medium text-gray-700 font-serif">
                関連ワード
              </Label>
              <Input
                id="shrineTag"
                type="text"
                placeholder="例：お伊勢参り"
                className="w-full border-2 border-red-800 rounded-md p-2 font-serif"
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="shrineFoundedYear" className="text-lg font-medium text-gray-700 font-serif">
                創建年
              </Label>
              <Input
                id="shrineFoundedYear"
                type="text"
                placeholder="例：垂仁天皇26年"
                className="w-full border-2 border-red-800 rounded-md p-2 font-serif"
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="shrineObjectOfWorship" className="text-lg font-medium text-gray-700 font-serif">
                御祭神
              </Label>
              <Input
                id="shrineObjectOfWorship"
                type="text"
                placeholder="例：天照坐皇大御神"
                className="w-full border-2 border-red-800 rounded-md p-2 font-serif"
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="shrineHasGoshuin" className="text-lg font-medium text-gray-700 font-serif">
                御朱印
              </Label>
              <Input
                id="shrineHasGoshuin"
                type="text"
                placeholder="例：あり"
                className="w-full border-2 border-red-800 rounded-md p-2 font-serif"
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="shrineWebsiteURL" className="text-lg font-medium text-gray-700 font-serif">
              公式サイトURL
              </Label>
              <Input
                id="shrineWebsiteURL"
                type="text"
                placeholder="例：https://www.isejingu.or.jp/"
                className="w-full border-2 border-red-800 rounded-md p-2 font-serif"
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
      </div>
    </>

  );
  
};

export default AdminRegisterShrineDetails;