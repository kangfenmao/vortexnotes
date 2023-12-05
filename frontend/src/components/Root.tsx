import { Outlet, useNavigate } from 'react-router-dom'
import Navbar from '@/components/Navbar.tsx'
import React, { useEffect } from 'react'
import { fetchConfig } from '@/utils/api.ts'

const Root: React.FC = () => {
  const navigate = useNavigate()

  useEffect(() => {
    const onLoad = async () => {
      if (!localStorage.vortexnotes_auth_scope) {
        navigate('/init', { replace: true })
      } else {
        await fetchConfig()
      }
    }

    window.addEventListener('load', onLoad)

    return () => window.removeEventListener('load', onLoad)
  }, [navigate])

  return (
    <div className="flex flex-col flex-1 w-full h-full">
      <Navbar />
      <Outlet />
    </div>
  )
}

export default Root
