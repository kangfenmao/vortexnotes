import React from 'react'

interface Props
  extends React.DetailedHTMLProps<React.HTMLAttributes<HTMLDivElement>, HTMLDivElement> {
  title: string
}

const EmptyView: React.FC<Props> = ({ title, className, style }) => {
  return (
    <div className={'flex flex-col items-center justify-center' + ' ' + className} style={style}>
      <i className="iconfont icon-empty text-6xl opacity-70" />
      <div className="text-xl opacity-70 mt-2">{title}</div>
    </div>
  )
}

export default EmptyView
