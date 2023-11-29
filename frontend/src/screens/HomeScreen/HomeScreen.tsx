import React, { useState } from 'react'
import RecentlyNotes from '@/components/RecentlyNotes.tsx'
import Navbar from '@/components/Navbar.tsx'
import { useNavigate } from 'react-router-dom'
import { onSearch as search } from '@/utils'

const HomeScreen: React.FC = () => {
  const [keywords, setKeywords] = useState('')
  const navigate = useNavigate()

  const onSearch = search.bind(this, keywords, navigate)

  return (
    <main className="flex flex-1">
      <Navbar />
      <div className="flex flex-col mt-28 sm:m-auto items-center w-full px-5 transition-all duration-150 max-w-xl">
        <h1
          className="font-bold align-middle mb-8 text-4xl sm:text-5xl select-none"
          style={{ fontFamily: 'Major Mono Display' }}>
          <span className="text-red-500">V</span>
          <span className="text-violet-700">N</span>
          <span>OTE</span>
        </h1>
        <section className="relative mb-10 w-full border border-white border-opacity-30 rounded-md">
          <input
            type="text"
            name="keywords"
            value={keywords}
            onChange={e => setKeywords(e.target.value)}
            placeholder="Search Notes"
            className="px-4 py-3 w-full outline-none bg-white border border-gray-300 dark:bg-zinc-800 dark:border-transparent rounded-md"
            onKeyDown={e => e.key === 'Enter' && onSearch()}
            autoComplete="off"
            autoFocus
            required
          />
          <button
            type="button"
            onClick={onSearch}
            className="absolute top-0 bottom-0 right-0 w-12 flex flex-row justify-center items-center cursor-pointer">
            <i className="iconfont icon-search text-2xl opacity-50 hover:opacity-90 transition-opacity"></i>
          </button>
        </section>
        <RecentlyNotes />
      </div>
    </main>
  )
}

export default HomeScreen
