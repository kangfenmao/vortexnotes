import React from 'react'

const LoadingView: React.FC = () => {
  return (
    <div className="flex flex-row justify-center py-10">
      <i className="iconfont icon-loading text-3xl animate-spin opacity-70"></i>
    </div>
  )
}

export default LoadingView
