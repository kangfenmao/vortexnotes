import React from 'react'

interface Props {
  text: string
  highlight: string
}

const HighlightText: React.FC<Props> = ({ text, highlight }) => {
  const words = highlight.split(' ')
  const regex = highlight.includes(' ')
    ? new RegExp(`(${words.join('|')})`, 'gi')
    : new RegExp(`(${highlight})`, 'gi')
  const parts = text.split(regex)

  return (
    <span>
      {parts.map((part, index) => {
        return highlight.toLowerCase().includes(part.toLowerCase()) ? (
          <span key={index} className="text-red-400">
            {part}
          </span>
        ) : (
          part
        )
      })}
    </span>
  )
}

export default HighlightText
