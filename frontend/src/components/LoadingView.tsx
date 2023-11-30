import React from 'react'

const LoadingView: React.FC = () => {
  return (
    <div className="flex flex-row justify-center py-10">
      <span className="loading loading-ring loading-lg"></span>
    </div>
  )
}

export default LoadingView
