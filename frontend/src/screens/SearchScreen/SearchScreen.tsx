import { Link, useNavigate, useSearchParams } from 'react-router-dom'
import React, { useEffect, useState } from 'react'
import { displayName, runAsyncFunction } from '@/utils'
import Navbar from '@/components/Navbar.tsx'
import { NoteType } from '@/types'
import './index.css'
import LoadingView from '@/components/LoadingView.tsx'
import useDebouncedValue from '@/hooks/useDebouncedValue.ts'
import EmptyView from '@/components/EmptyView.tsx'
import { isEmpty } from 'lodash-es'
import dayjs from 'dayjs'

interface SearchResponse {
  data: NoteType[]
  duration: number
  page: number
  total_page: number
}

const cachedSearchResult = new Map<string, NoteType[]>()
let prevSearchKeywords = ''

const SearchScreen: React.FC = () => {
  const [searchParams] = useSearchParams()
  const keywords = searchParams.get('keywords') || ''
  const [time, setTime] = useState(0)
  const [notes, setNotes] = useState<NoteType[]>(cachedSearchResult.get(keywords) || [])
  const [empty, setEmpty] = useState(false)
  const [loading, setLoading] = useState(false)
  const navigate = useNavigate()
  useEffect(() => {
    runAsyncFunction(async () => {
      setLoading(true)
      const res = await window.$http.get<SearchResponse>(`search?keywords=${keywords}`)
      const _notes = res.data.data
      setTime(res.data.duration)
      setEmpty(isEmpty(_notes))
      setNotes(_notes)
      setLoading(false)
    })
  }, [keywords])

  useEffect(() => {
    if (keywords !== prevSearchKeywords) {
      window.scrollTo({ top: 0 })
    }
    prevSearchKeywords = keywords
  }, [keywords])

  const isLoading = useDebouncedValue(false, loading && !notes.length, 1000)

  useEffect(() => {
    return () => {
      keywords && cachedSearchResult.set(keywords, notes)
    }
  }, [keywords, notes])

  return (
    <main className="w-full flex">
      <Navbar />
      <div className="container mx-auto mt-20 px-5 max-w-lg sm:max-w-6xl flex flex-col flex-1">
        <div className="mb-4 pt-2 text-sm" style={{ color: '#9aa0a6' }}>
          About {notes.length} results ({time.toFixed(6)} seconds)
        </div>
        {isLoading && <LoadingView />}
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
              <h4
                className="mb-2 font-bold text-opacity-70 text-xl"
                dangerouslySetInnerHTML={{ __html: displayName(note.name) }}></h4>
            </Link>
            <p
              className="font-medium opacity-80 line-clamp-5 search-content cursor-pointer"
              onClick={() => navigate(`/notes/${note.id}`)}
              dangerouslySetInnerHTML={{ __html: note.content }}></p>
            <span className="text-xs text-opacity-50">
              {dayjs(note.created_at).format('YYYY/MM/DD HH:mm')}
            </span>
          </div>
        ))}
        <footer className="py-5"></footer>
      </div>
    </main>
  )
}

export default SearchScreen
