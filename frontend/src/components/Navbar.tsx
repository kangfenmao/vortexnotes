import { Link, useLocation, useNavigate, useSearchParams } from 'react-router-dom'
import SearchIcon from '@/assets/images/search_icon.svg'
import React, { useState } from 'react'
import AddIcon from '@/assets/images/add_icon.svg'

interface Props {}

const Navbar: React.FC<Props> = () => {
  const location = useLocation()
  const navigate = useNavigate()
  const [searchParams] = useSearchParams()
  const keywords = searchParams.get('keywords') || ''
  const [input, setInput] = useState(keywords)

  const isHome = location.pathname === '/'

  const onGoBack = () => navigate('/')

  return (
    <div
      className="flex flex-row h-16 items-center fixed top-0 left-0 right-0"
      style={{
        backgroundColor: isHome ? 'transparent' : '#020409',
        borderBottomColor: isHome ? 'transparent' : 'rgba(255,255,255,0.2)',
        borderBottomWidth: '0.5px',
        zIndex: 100
      }}>
      <div className="flex flex-row items-center m-auto w-full px-5 max-w-lg sm:max-w-6xl">
        {isHome && <div className="flex-1"></div>}
        {!isHome && (
          <h1
            className="text-2xl md:text-4xl font-bold text-white mr-5 cursor-pointer"
            style={{ fontFamily: 'Major Mono Display' }}
            onClick={onGoBack}>
            <span className="text-red-500">V</span>
            <span className="text-violet-700">N</span>
            <span>OTE</span>
          </h1>
        )}
        {!isHome && (
          <div className="flex flex-1">
            <form className="relative w-full md:w-1/2" method="get" action="/search">
              <input
                type="text"
                name="keywords"
                value={input}
                onChange={e => setInput(e.target.value)}
                placeholder="Search"
                className="w-full px-4 py-2 outline-none rounded-md bg-transparent border-white border-opacity-30"
                style={{ borderWidth: 0.5 }}
                required
              />
              <button
                type="submit"
                className="absolute top-0 bottom-0 right-0 w-12 flex flex-row justify-center items-center cursor-pointer opacity-70">
                <img src={SearchIcon} alt="" />
              </button>
            </form>
          </div>
        )}
        <Link to="/new">
          <div className="flex flex-row items-center ml-5 opacity-60 hover:opacity-80 transition-opacity">
            <img src={AddIcon} alt="" className="mr-2" />
            <button className="text-white">New Note</button>
          </div>
        </Link>
      </div>
    </div>
  )
}

export default Navbar
