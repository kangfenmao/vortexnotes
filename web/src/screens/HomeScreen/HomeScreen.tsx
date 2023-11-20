import SearchIcon from '@/assets/images/search_icon.svg'
import React from 'react'

const HomeScreen: React.FC = () => {
  return (
    <div className="flex flex-col m-auto items-center w-full px-5 transition-all duration-150 max-w-xl">
      <h1
        className="text-5xl font-bold align-middle mb-10 sm:text-6xl"
        style={{ fontFamily: 'Major Mono Display' }}>
        <span className="text-violet-700">V</span>
        <span className="text-red-500">o</span>rtex
      </h1>
      <form className="relative mb-40 w-full" method="get" action="/search">
        <input
          type="text"
          name="keywords"
          placeholder="Search"
          className="px-4 py-2 sm:py-3 w-full"
        />
        <button
          type="submit"
          className="absolute top-0 bottom-0 right-0 w-12 flex flex-row justify-center items-center cursor-pointer">
          <img src={SearchIcon} alt="" />
        </button>
      </form>
    </div>
  )
}

export default HomeScreen
