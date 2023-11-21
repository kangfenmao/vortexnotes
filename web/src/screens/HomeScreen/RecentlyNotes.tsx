import React, { useEffect, useState } from 'react'
import { displayName, runAsyncFunction } from '@/utils'
import { Link } from 'react-router-dom'
import { NoteType } from '@/types'

const RecentlyNotes: React.FC = () => {
  const [notes, setNotes] = useState<NoteType[]>([])

  useEffect(() => {
    runAsyncFunction(async () => {
      const res = await window.$http.get('notes')
      setNotes(res.data)
    })
  }, [])

  return (
    <div className="flex flex-col items-center" style={{ minHeight: '350px' }}>
      <h6 className="text-xs opacity-60 mb-2">RECENTLY MODIFIED</h6>
      <ul>
        {notes.map(note => (
          <li className="flex flex-col items-center">
            <Link
              to={`/notes/${note.id}`}
              className="py-1 px-3 text-white text-md opacity-50 hover:text-white hover:opacity-80 hover:bg-zinc-700 rounded-sm">
              {displayName(note.name)}
            </Link>
          </li>
        ))}
      </ul>
    </div>
  )
}

export default RecentlyNotes