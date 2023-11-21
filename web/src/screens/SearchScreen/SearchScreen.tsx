import { Link, useSearchParams } from 'react-router-dom'
import React, { useEffect, useState } from 'react'
import { displayName, runAsyncFunction } from '@/utils'
import HighlightText from '@/screens/SearchScreen/HighlightText.tsx'
import Navbar from '@/components/Navbar.tsx'

const SearchScreen: React.FC = () => {
  const [searchParams] = useSearchParams()
  const [notes, setNotes] = useState<any>([])
  const [time, setTime] = useState(0)
  const keywords = searchParams.get('keywords') || ''

  useEffect(() => {
    runAsyncFunction(async () => {
      const res = await window.$http.get(`search?keywords=${keywords}`)
      setNotes(res.data.data)
      setTime(res.data.duration)
    })
  }, [])

  return (
    <main className="w-full">
      <Navbar />
      <div className="container mx-auto mt-24 px-5 max-w-lg sm:max-w-6xl">
        <div className="mb-5 text-sm" style={{ color: '#9aa0a6' }}>
          找到约 {notes.length} 条结果 (用时{time}秒)
        </div>
        {notes.map((note: any) => (
          <div className="mb-5" key={notes.id}>
            <Link to={`/notes/${note.id}`}>
              <h4 className="mb-2 text-blue-400 font-bold text-xl">{displayName(note.name)}</h4>
            </Link>
            <p className="text-md line-clamp-5 opacity-90">
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
