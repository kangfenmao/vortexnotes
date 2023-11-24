import SearchIcon from '@/assets/images/search_icon.svg'
import React from 'react'
import RecentlyNotes from '@/components/RecentlyNotes.tsx'
import Navbar from '@/components/Navbar.tsx'

const HomeScreen: React.FC = () => {
  return (
    <main className="flex flex-1">
      <Navbar />
      <div className="flex flex-col m-auto items-center w-full px-5 transition-all duration-150 max-w-xl">
        <h1
          className="font-bold align-middle mb-8 text-4xl sm:text-5xl"
          style={{ fontFamily: 'Major Mono Display' }}>
          <span className="text-red-500">V</span>
          <span className="text-violet-700">N</span>
          <span>OTE</span>
        </h1>
        <form className="relative mb-10 w-full" method="get" action="/search">
          <input
            type="text"
            name="keywords"
            placeholder="Search"
            className="px-4 py-2 sm:py-3 w-full outline-none rounded-sm"
            style={{ backgroundColor: '#3b3b3b' }}
            autoFocus
            required
          />
          <button
            type="submit"
            className="absolute top-0 bottom-0 right-0 w-12 flex flex-row justify-center items-center cursor-pointer">
            <img src={SearchIcon} alt="" />
          </button>
        </form>
        <RecentlyNotes />
      </div>
    </main>
  )
}

export default HomeScreen
