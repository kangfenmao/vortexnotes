import React, { useState } from 'react'
import RecentlyNotes from '@/components/RecentlyNotes.tsx'
import { useNavigate } from 'react-router-dom'
import { onSearch as search } from '@/utils'

const HomeScreen: React.FC = () => {
  const [keywords, setKeywords] = useState('')
  const navigate = useNavigate()

  const onSearch = search.bind(this, keywords, navigate)

  return (
    <main className="flex flex-1">
      <div className="flex flex-col mt-28 sm:m-auto items-center w-full px-5 transition-all duration-150 max-w-xl">
        <div className="flex flex-col items-center">
          <img src="/public/icon-200x200.png" className="w-28" alt="" />
          <span style={{ fontFamily: 'Major Mono Display' }} className="text-sm mb-5">
            vortex notes
          </span>
        </div>
        <section className="relative mb-10 w-full border border-black border-opacity-20 dark:border-white dark:border-opacity-30 rounded-md">
          <input
            type="text"
            name="keywords"
            value={keywords}
            onChange={e => setKeywords(e.target.value)}
            placeholder="Search Notes"
            className="px-4 py-3 w-full outline-none bg-white dark:bg-zinc-800 rounded-md"
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
