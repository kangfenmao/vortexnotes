import { Link, useSearchParams } from 'react-router-dom'
import React, { useEffect, useState } from 'react'
import { displayName } from '@/utils'
import Navbar from '@/components/Navbar.tsx'
import { NoteType } from '@/types'
import useRequest from '@/hooks/useRequest.ts'
import './SearchScreen.css'
import LoadingView from '@/components/LoadingView.tsx'
import useDebouncedValue from '@/hooks/useDebouncedValue.ts'
import EmptyView from '@/components/EmptyView.tsx'
import { isEmpty } from 'lodash'

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
  const [notes, setNotes] = useState<NoteType[]>([])
  const [empty, setEmpty] = useState(false)

  const { data, isLoading } = useRequest<SearchResponse>({
    method: 'GET',
    url: `search?keywords=${keywords}`
  })

  const loading = useDebouncedValue(false, isLoading && !notes.length, 1000)

  useEffect(() => {
    if (data) {
      const _notes = [...notes, ...data.data]
      setNotes(_notes)
      setTime(data.duration)
      isEmpty(_notes) && setEmpty(true)
    }
  }, [data])

  useEffect(() => {
    setNotes([])
    setKeywords(searchParams.get('keywords') || '')
  }, [searchParams])

  return (
    <main className="w-full flex">
      <Navbar />
      <div className="container mx-auto mt-20 px-5 max-w-lg sm:max-w-6xl flex flex-col flex-1">
        <div className="mb-4 pt-2 text-sm" style={{ color: '#9aa0a6' }}>
          找到约 {notes.length} 条结果 (用时{time}秒)
        </div>
        {loading && <LoadingView />}
        {empty && (
          <EmptyView
            title={'No Search Results'}
            className="flex-1 font-thin"
            style={{ marginTop: '-10%' }}
          />
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
