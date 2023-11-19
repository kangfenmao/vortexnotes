import SearchIcon from '@/assets/images/search_icon.svg'
import { Link, useSearchParams } from 'react-router-dom'
import React, { useEffect, useState } from 'react'
import { runAsyncFunction } from '@/utils'
import HighlightText from '@/screens/SearchScreen/HighlightText.tsx'

const SearchScreen: React.FC = () => {
  const [searchParams] = useSearchParams()
  const [notes, setNotes] = useState<any>([])
  const keywords = searchParams.get('keywords') || ''

  useEffect(() => {
    runAsyncFunction(async () => {
      const res = await window.$http.get(`search?keywords=${keywords}`)
      setNotes(res.data)
    })
  }, [])

  return (
    <main className="container mx-auto max-w-3xl px-4">
      <div className="flex flex-row h-20 items-center">
        <Link to="/">
          <h1
            className="text-4xl font-bold text-white"
            style={{ fontFamily: 'Major Mono Display' }}>
            Vortex
          </h1>
        </Link>
      </div>
      <form className="relative mb-8" method="get" action="search">
        <input
          type="text"
          name="keywords"
          value={keywords}
          placeholder="Search"
          className="w-full px-6 py-2 border border-gray-300 rounded-full focus:outline-none focus:border-blue-500"
        />
        <button
          type="submit"
          className="absolute top-0 bottom-0 right-0 w-16 flex flex-row justify-center items-center rounded-full cursor-pointer">
          <img src={SearchIcon} alt="" />
        </button>
      </form>
      {notes.map((note: any) => (
        <div className="mb-5">
          <Link to="notes">
            <h6 className="mb-2 text-blue-500">{note.name}</h6>
          </Link>
          <p className="text-xs">
            <HighlightText text={note.content} highlight={keywords} />
          </p>
        </div>
      ))}
      <div className="h-10"></div>
    </main>
  )
}

export default SearchScreen
