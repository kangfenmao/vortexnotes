import React from 'react'
import _ from 'lodash'
import { NoteType } from '@/types'
import { displayName } from '@/utils'
import { Link } from 'react-router-dom'

interface Props {
  data: NoteType[]
}

const GroupedNotes: React.FC<Props> = ({ data }) => {
  const sortedNotes = _.sortBy(data, 'name')

  const groupedData = _.groupBy(sortedNotes, item => {
    const firstChar = item.name[0]
    // eslint-disable-next-line no-control-regex
    return /[^\u0000-\u00FF]/.test(firstChar) ? '#' : firstChar.toUpperCase()
  })

  return (
    <div>
      {Object.entries(groupedData).map(([letter, items]) => (
        <div key={letter} className="mb-6">
          <h2 className="text-3xl font-bold opacity-30 mt-2">{letter}</h2>
          <hr className="my-2 border-b-1 border-black border-opacity-10" />
          <ul>
            {items.map(note => (
              <div className="py-1" key={note.id}>
                <Link
                  to={`/notes/${note.id}`}
                  className="text-black hover:text-black dark:text-white dark:hover:text-white opacity-70 hover:opacity-90 line-clamp-1">
                  {displayName(note.name)}
                </Link>
              </div>
            ))}
          </ul>
        </div>
      ))}
    </div>
  )
}

export default GroupedNotes
