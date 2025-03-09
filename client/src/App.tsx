import { Routes, Route } from 'react-router-dom';

import Home from '@/components/model/Home';
import PrefList from '@/components/model/PrefList';
import ShrineInfo from '@/components/model/ShrineInfo';
import ShrineSacList from '@/components/model/ShrineSacList';
import ShrineTagList from '@/components/model/ShrineTagList';
import Admin from '@/components/model/Admin';
import AdminRegisterShrine from '@/components/model/AdminRegisterShrine';
import AdminRegisterShrineDetails from '@/components/model/AdminRegisterShrineDetails';
import AdminBulkRegisterShrine from '@/components/model/AdminBulkRegisterShrine';
import AdminStdAreaCode from '@/components/model/AdminStdAreaCode';

import '@/styles/global.css';

function App() {
    return (
      <div>
      <Routes>
        <Route path="/" element={ <Home /> } />
        <Route path="/prefs" element={ <PrefList /> } />
        <Route path="/shrine" element={ <ShrineInfo /> } />
        <Route path="/shrines/sac" element={ <ShrineSacList /> } />
        <Route path="/shrines/tag" element={ <ShrineTagList /> } />
        <Route path="/admin" element={ <Admin /> } />
        <Route path="/admin/regist/shrine" element={ <AdminRegisterShrine /> } />
        <Route path="/admin/regist/shrine-details" element={ <AdminRegisterShrineDetails /> } />
        <Route path="/admin/bulk-regist/shrine" element={ <AdminBulkRegisterShrine /> } />
        <Route path="/admin/stdareacode" element={ <AdminStdAreaCode /> } />
      </Routes>
      </div>
    );
};

export default App;