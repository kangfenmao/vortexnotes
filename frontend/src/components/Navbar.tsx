import { Link, useLocation, useNavigate, useSearchParams } from 'react-router-dom'
import React, { useState } from 'react'

interface Props {}

const Navbar: React.FC<Props> = () => {
  const location = useLocation()
  const navigate = useNavigate()
  const [searchParams] = useSearchParams()
  const keywords = searchParams.get('keywords') || ''
  const [input, setInput] = useState(keywords)
  const isHome = location.pathname === '/'

  const onGoBack = () => navigate('/')
  const onSearch = () => input.trim() && navigate(`/search?keywords=${input}`)

  return (
    <div
      className="flex flex-row h-16 items-center fixed top-0 left-0 right-0"
      style={{
        backgroundColor: isHome ? 'transparent' : '#020409',
        borderBottomColor: isHome ? 'transparent' : 'rgba(255,255,255,0.2)',
        borderBottomWidth: '0.5px',
        zIndex: 100
      }}>
      <div className="flex flex-row items-center justify-between m-auto w-full px-5 max-w-lg sm:max-w-6xl">
        {isHome && <div className="flex-1"></div>}
        {!isHome && (
          <h1
            className="text-2xl md:text-4xl font-bold text-white mr-5 cursor-pointer select-none"
            style={{ fontFamily: 'Major Mono Display' }}
            onClick={onGoBack}>
            <span className="text-red-500">V</span>
            <span className="text-violet-700">N</span>
            <span>OTE</span>
          </h1>
        )}
        {!isHome && (
          <div className="flex-1 hidden sm:flex">
            <section className="relative w-full md:w-1/2">
              <input
                type="text"
                name="keywords"
                value={input}
                onChange={e => setInput(e.target.value)}
                placeholder="Search"
                className="w-full px-4 py-2 outline-none rounded-md bg-transparent border-white border-opacity-30"
                style={{ borderWidth: 0.5 }}
                onKeyDown={e => e.key === 'Enter' && onSearch()}
                autoComplete="off"
                required
              />
              <button
                type="button"
                onClick={onSearch}
                className="absolute top-0 bottom-0 right-0 w-12 flex flex-row justify-center items-center cursor-pointer opacity-70">
                <i className="iconfont icon-search text-white text-1xl mr-1"></i>
              </button>
            </section>
          </div>
        )}
        <Link to="/new">
          <div className="flex flex-row items-center ml-5 opacity-60 hover:opacity-80 transition-opacity">
            <i className="iconfont icon-add-circle text-white text-2xl"></i>
            <button className="text-white ml-1 hidden sm:inline">New Note</button>
          </div>
        </Link>
      </div>
    </div>
  )
}

export default Navbar
