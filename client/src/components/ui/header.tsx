import { Menu } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger} from '@/components/ui/dropdown-menu';

const FRONTEND_URL=import.meta.env.VITE_FRONTEND_URL

const Header: React.FC = () => {
  return (
    <header className="bg-red-900 text-white">
      <div className="container mx-auto px-4 py-6">
        <div className="flex items-center justify-between">
          <div className="flex items-center space-x-2">
            <svg
              className="w-8 h-8"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M3 21h18M3 10h18M3 7l9-4 9 4M4 10h16v11H4V10z"
              />
            </svg>
            <a href="/" className="text-2xl font-semibold">神社ガイド</a>
            {import.meta.env.VITE_ENV_KEY === "DEVELOP"
              ? ( <p className="align-items flex-end">開発環境</p> )
              : null
            }
          </div>
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost" aria-label="メニュー">
                <Menu className="h-6 w-6" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuItem>
              <a href={FRONTEND_URL + "/prefs"}>都道府県／市区町村検索</a>
              </DropdownMenuItem>
              <DropdownMenuItem>
                <a href={FRONTEND_URL + "/admin"}>管理者画面</a>
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      </div>
    </header>
  )
};

export {Header};