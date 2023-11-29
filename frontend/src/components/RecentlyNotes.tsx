import React, { useEffect, useState } from 'react'
import { displayName } from '@/utils'
import { Link } from 'react-router-dom'
import { NoteType } from '@/types'
import useRequest from '@/hooks/useRequest.ts'

const RecentlyNotes: React.FC = () => {
  const [notes, setNotes] = useState<NoteType[]>([])
  const { data } = useRequest<NoteType[]>({
    method: 'GET',
    url: 'notes?page=1&limit=5&sort=updated_at:desc'
  })

  useEffect(() => {
    data && setNotes(data)
  }, [data])

  return (
    <div className="flex flex-col items-center" style={{ minHeight: '200px' }}>
      <h6 className="text-xs opacity-60 mb-2">RECENTLY MODIFIED</h6>
      <ul>
        {notes.map(note => (
          <li className="flex flex-col items-center" key={note.id}>
            <Link
              to={`/notes/${note.id}`}
              className="py-1 px-3 text-md transition-all rounded-sm opacity-50 text-black hover:text-black dark:text-white dark:hover:text-white">
              {displayName(note.name)}
            </Link>
          </li>
        ))}
      </ul>
    </div>
  )
}

export default RecentlyNotes
