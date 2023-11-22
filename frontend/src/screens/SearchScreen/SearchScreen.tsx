import { Link, useSearchParams } from 'react-router-dom'
import React, { useEffect, useState } from 'react'
import { displayName } from '@/utils'
import HighlightText from '@/screens/SearchScreen/HighlightText.tsx'
import Navbar from '@/components/Navbar.tsx'
import { NoteType } from '@/types'
import useRequest from '@/hooks/useRequest.ts'

const SearchScreen: React.FC = () => {
  const [searchParams] = useSearchParams()
  const [notes, setNotes] = useState<NoteType[]>([])
  const [time, setTime] = useState(0)
  const keywords = searchParams.get('keywords') || ''
  const { data } = useRequest<{ data: NoteType[]; duration: number }>({
    method: 'GET',
    url: `search?keywords=${keywords}`
  })

  useEffect(() => {
    data && setNotes(data.data)
    data && setTime(data.duration)
  }, [data])

  return (
    <main className="w-full">
      <Navbar />
      <div className="container mx-auto mt-24 px-5 max-w-lg sm:max-w-6xl">
        <div className="mb-5 text-sm" style={{ color: '#9aa0a6' }}>
          找到约 {notes.length} 条结果 (用时{time}秒)
        </div>
        {notes.map(note => (
          <div className="mb-5" key={note.id}>
            <Link to={`/notes/${note.id}`} target="_blank">
              <h4 className="mb-2 text-blue-400 font-bold text-xl">{displayName(note.name)}</h4>
            </Link>
            <p className="text-sm font-medium text-white opacity-80 line-clamp-5">
              <HighlightText text={note.content} highlight={keywords} />
            </p>
          </div>
        ))}
        <div className="h-10"></div>
      </div>
    </main>
  )
}

export default SearchScreen
