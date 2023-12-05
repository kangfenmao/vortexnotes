import { useEffect } from 'react'
import { hasPermission } from '@/utils'
import { useNavigate } from 'react-router-dom'

export default function usePermissionCheckEffect(scope: string) {
  const navigate = useNavigate()

  useEffect(() => {
    if (!hasPermission(scope)) {
      navigate('/auth')
    }
  }, [navigate, scope])
}
