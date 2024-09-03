import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import Analyzer, { DiskSelector, DiskAnalysis } from './screens/Analyzer';
import { HashRouter, Routes, Route } from "react-router-dom";
import './index.css'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <HashRouter basename={"/"}>
      <Routes>
        <Route index path="/" element={<App />} />
        <Route path="/analyzer" element={<Analyzer />}>
          <Route index element={<DiskSelector />} />
          <Route path=":disk" element={<DiskAnalysis />} />
        </Route>
      </Routes>
    </HashRouter>
  </React.StrictMode>
)
