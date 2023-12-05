import { Outlet } from 'react-router-dom'
import Navbar from '@/components/Navbar.tsx'
import React from 'react'

const Root: React.FC = () => {
  return (
    <div className="flex flex-col flex-1 w-full h-full">
      <Navbar />
      <Outlet />
    </div>
  )
}

export default Root
