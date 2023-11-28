import { TopView } from '@/components/TopView.tsx'
import React, { useEffect } from 'react'

let topViewKey = 0

const ToastContainer: React.FC<ToastProps> = ({ message }) => {
  useEffect(() => {
    setTimeout(() => TopView.hide(topViewKey), 1500)
  }, [])

  return (
    <div className="toast toast-top toast-center mt-20" style={{ zIndex: 1000 }}>
      <div className="alert flex justify-center">
        <span className="font-bold">{message}</span>
      </div>
    </div>
  )
}

interface ToastProps {
  message: string
}

const Toast = {
  show: (params: ToastProps) => {
    topViewKey = TopView.show(<ToastContainer {...params} />)
  }
}

export default Toast
