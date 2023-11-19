import React from 'react'

interface Props {
  text: string
  highlight: string
}

const HighlightText: React.FC<Props> = ({ text, highlight }) => {
  const words = highlight.split(' ')
  const regex = new RegExp(`(${words.join('|')})`, 'gi')
  const parts = text.split(regex)

  return (
    <span>
      {parts.map((part, index) =>
        words.includes(part.toLowerCase()) ? (
          <span key={index} className="text-red-500">
            {part}
          </span>
        ) : (
          part
        )
      )}
    </span>
  )
}

export default HighlightText
