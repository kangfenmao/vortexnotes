import { Link, useSearchParams } from 'react-router-dom'
import React, { useEffect, useState } from 'react'
import { displayName } from '@/utils'
import Navbar from '@/components/Navbar.tsx'
import { NoteType } from '@/types'
import useRequest from '@/hooks/useRequest.ts'
import { useBottomScrollListener } from 'react-bottom-scroll-listener'
import './SearchScreen.css'

interface SearchResponse {
  data: NoteType[]
  duration: number
  page: number
  total_page: number
}

const SearchScreen: React.FC = () => {
  const [searchParams] = useSearchParams()
  const [time, setTime] = useState(0)
  const [keywords, setKeywords] = useState(searchParams.get('keywords' || ''))
  const [page, setPage] = useState(1)
  const [end, setEnd] = useState(false)
  const [notes, setNotes] = useState<NoteType[]>([])
  const limit = 20

  const { data, isLoading } = useRequest<SearchResponse>({
    method: 'GET',
    url: `search?keywords=${keywords}&page=${page}&limit=${limit}`
  })

  const nextPage = () => !end && setPage(page + 1)

  useBottomScrollListener(nextPage, {
    debounce: 1000,
    offset: 0
  })

  useEffect(() => {
    if (data) {
      const notesCount = data.data.length
      setNotes([...notes, ...data.data])
      notesCount < limit && setEnd(true)
      setTime(data.duration)
    }
  }, [data])

  useEffect(() => {
    setNotes([])
    setKeywords(searchParams.get('keywords') || '')
  }, [searchParams])

  return (
    <main className="w-full">
      <Navbar />
      <div className="container mx-auto mt-24 px-5 max-w-lg sm:max-w-6xl">
        <div className="mb-5 text-sm" style={{ color: '#9aa0a6' }}>
          找到约 {notes.length} 条结果 (用时{time}秒)
        </div>
        {isLoading && !notes.length && (
          <div className="flex flex-row justify-center py-10">
            <i className="iconfont icon-loading text-3xl animate-spin opacity-70"></i>
          </div>
        )}
        {notes.map((note, index) => (
          <div className="mb-5" key={note.id + '_' + index}>
            <Link to={`/notes/${note.id}`}>
              <h4 className="mb-2 text-blue-400 font-bold text-xl">{displayName(note.name)}</h4>
            </Link>
            <p
              className="text-sm font-medium text-white opacity-80 line-clamp-5 search-content"
              dangerouslySetInnerHTML={{ __html: note._formatted!.content }}></p>
          </div>
        ))}
        <div className="h-10"></div>
      </div>
    </main>
  )
}

export default SearchScreen
