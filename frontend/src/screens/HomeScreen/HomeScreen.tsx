import React from 'react'
import RecentlyNotes from '@/components/RecentlyNotes.tsx'
import Navbar from '@/components/Navbar.tsx'

const HomeScreen: React.FC = () => {
  return (
    <main className="flex flex-1">
      <Navbar />
      <div className="flex flex-col m-auto items-center w-full px-5 transition-all duration-150 max-w-xl">
        <h1
          className="font-bold align-middle mb-8 text-4xl sm:text-5xl select-none"
          style={{ fontFamily: 'Major Mono Display' }}>
          <span className="text-red-500">V</span>
          <span className="text-violet-700">N</span>
          <span>OTE</span>
        </h1>
        <form className="relative mb-10 w-full" method="get" action="/search">
          <input
            type="text"
            name="keywords"
            placeholder="Search Notes"
            className="px-4 py-2 sm:py-3 w-full outline-none"
            style={{ backgroundColor: '#1f2227' }}
            autoFocus
            required
          />
          <button
            type="submit"
            className="absolute top-0 bottom-0 right-0 w-12 flex flex-row justify-center items-center cursor-pointer">
            <i className="iconfont icon-search text-white text-2xl opacity-70 hover:opacity-100 transition-opacity"></i>
          </button>
        </form>
        <RecentlyNotes />
      </div>
    </main>
  )
}

export default HomeScreen
