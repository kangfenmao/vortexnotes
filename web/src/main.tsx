import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import '@/config/http.ts'
import App from '@/App.tsx'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
)
