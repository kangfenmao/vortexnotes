import LoadingView from '@/components/LoadingView.tsx'
import React, { useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { fetchConfig } from '@/utils/api.ts'

const InitScreen: React.FC = () => {
  const navigate = useNavigate()

  useEffect(() => {
    fetchConfig().then(() => navigate('/', { replace: true }))
  }, [navigate])

  return (
    <main className="flex flex-1 justify-center items-center">
      <LoadingView />
    </main>
  )
}

export default InitScreen
