import React from 'react'
import { groupBy } from 'lodash-es'
import { NoteType } from '@/types'
import { displayName } from '@/utils'
import { Link } from 'react-router-dom'
import dayjs from 'dayjs'

interface Props {
  data: NoteType[]
}

const GroupedNotes: React.FC<Props> = ({ data }) => {
  const groupedData = groupBy(data, item => {
    return dayjs(item.created_at).format('MM/DD')
  })

  return (
    <div>
      {Object.entries(groupedData).map(([letter, items]) => (
        <div key={letter} className="mb-6">
          <h2 className="text-3xl font-bold opacity-30 mt-2">{letter}</h2>
          <hr className="my-2 border-b-1 border-black border-opacity-10 dark:border-white dark:border-opacity-10" />
          <ul>
            {items.map(note => (
              <div className="py-1 flex flex-row justify-between" key={note.id}>
                <Link
                  to={`/notes/${note.id}`}
                  className="text-black hover:text-black dark:text-white dark:hover:text-white opacity-70 hover:opacity-90 line-clamp-1">
                  {displayName(note.name)}
                </Link>
                <span className="text-gray-500">{dayjs(note.created_at).format('YYYY/MM/DD')}</span>
              </div>
            ))}
          </ul>
        </div>
      ))}
    </div>
  )
}

export default GroupedNotes
