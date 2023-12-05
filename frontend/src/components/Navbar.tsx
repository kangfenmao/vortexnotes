import { Link, useLocation, useNavigate, useSearchParams } from 'react-router-dom'
import React, { useEffect, useState } from 'react'
import useTheme from '@/hooks/useTheme.ts'
import { hasPasscode, hasPermission, onSearch as search } from '@/utils'

interface Props {}

const Navbar: React.FC<Props> = () => {
  const location = useLocation()
  const navigate = useNavigate()
  const [searchParams] = useSearchParams()
  const keywords = searchParams.get('keywords') || ''
  const [input, setInput] = useState(keywords)
  const isHome = location.pathname === '/'
  const [theme, setTheme] = useTheme()

  const onSearch = search.bind(this, input, navigate)

  useEffect(() => {
    setInput(keywords)
  }, [keywords])

  const onLogout = () => {
    localStorage.removeItem('vortexnotes_passcode')
  }

  const navbarBg = isHome ? '' : 'bg-white dark:bg-black dark:bg-transparent-20'
  const navbarBorder = isHome
    ? 'border-b border-transparent'
    : 'border-b border-gray-200 dark:border-white dark:border-opacity-10'

  return (
    <div
      className={`flex flex-row h-16 items-center fixed top-0 left-0 right-0 ${navbarBorder} ${navbarBg}`}
      style={{
        zIndex: 100,
        backgroundColor: isHome ? 'transparent' : 'var(--theme-navbar-color)'
      }}>
      <div className="flex flex-row items-center justify-between m-auto w-full px-5 ">
        {isHome && <div className="flex-1"></div>}
        {!isHome && (
          <Link to="/" className="flex flex-row items-center">
            <img src="/public/icon-200x200.png" className="w-10 mr-2 hover:animate-spin" alt="" />
            <span
              style={{ fontFamily: 'Major Mono Display' }}
              className="text-xl text-black dark:text-white">
              VORTEX
            </span>
          </Link>
        )}
        {!isHome && (
          <div className="flex-1 hidden sm:flex mx-8">
            <section className="relative w-full rounded-full border border-black border-opacity-10 dark:border-white dark:border-opacity-20">
              <input
                type="text"
                name="keywords"
                value={input}
                onChange={e => setInput(e.target.value)}
                placeholder="Search"
                className="w-full px-4 py-2 outline-none rounded-md bg-transparent"
                onKeyDown={e => e.key === 'Enter' && onSearch()}
                autoComplete="off"
                required
              />
              <button
                type="button"
                onClick={onSearch}
                className="absolute top-0 bottom-0 right-0 w-12 flex flex-row justify-center items-center cursor-pointer opacity-70">
                <i className="iconfont icon-search opacity-70 text-black dark:text-white text-1xl mr-1"></i>
              </button>
            </section>
          </div>
        )}
        <div className="flex flex-row items-center">
          <label className="swap swap-rotate opacity-60 hover:opacity-80 transition-opacity">
            {/* this hidden checkbox controls the state */}
            <input
              type="checkbox"
              className="theme-controller"
              value="synthwave"
              checked={theme === 'dark'}
              onChange={() => setTheme(theme === 'dark' ? 'light' : 'dark')}
            />
            {/* sun icon */}
            <svg
              className="swap-on fill-current w-6 h-6"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24">
              <path d="M5.64,17l-.71.71a1,1,0,0,0,0,1.41,1,1,0,0,0,1.41,0l.71-.71A1,1,0,0,0,5.64,17ZM5,12a1,1,0,0,0-1-1H3a1,1,0,0,0,0,2H4A1,1,0,0,0,5,12Zm7-7a1,1,0,0,0,1-1V3a1,1,0,0,0-2,0V4A1,1,0,0,0,12,5ZM5.64,7.05a1,1,0,0,0,.7.29,1,1,0,0,0,.71-.29,1,1,0,0,0,0-1.41l-.71-.71A1,1,0,0,0,4.93,6.34Zm12,.29a1,1,0,0,0,.7-.29l.71-.71a1,1,0,1,0-1.41-1.41L17,5.64a1,1,0,0,0,0,1.41A1,1,0,0,0,17.66,7.34ZM21,11H20a1,1,0,0,0,0,2h1a1,1,0,0,0,0-2Zm-9,8a1,1,0,0,0-1,1v1a1,1,0,0,0,2,0V20A1,1,0,0,0,12,19ZM18.36,17A1,1,0,0,0,17,18.36l.71.71a1,1,0,0,0,1.41,0,1,1,0,0,0,0-1.41ZM12,6.5A5.5,5.5,0,1,0,17.5,12,5.51,5.51,0,0,0,12,6.5Zm0,9A3.5,3.5,0,1,1,15.5,12,3.5,3.5,0,0,1,12,15.5Z" />
            </svg>
            {/* moon icon */}
            <svg
              className="swap-off fill-current w-6 h-6"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24">
              <path d="M21.64,13a1,1,0,0,0-1.05-.14,8.05,8.05,0,0,1-3.37.73A8.15,8.15,0,0,1,9.08,5.49a8.59,8.59,0,0,1,.25-2A1,1,0,0,0,8,2.36,10.14,10.14,0,1,0,22,14.05,1,1,0,0,0,21.64,13Zm-9.5,6.69A8.14,8.14,0,0,1,7.08,5.22v.27A10.15,10.15,0,0,0,17.22,15.63a9.79,9.79,0,0,0,2.1-.22A8.11,8.11,0,0,1,12.14,19.73Z" />
            </svg>
          </label>
          {hasPermission('create') && (
            <Link to="/new">
              <button className="flex flex-row items-center ml-5 opacity-60 hover:opacity-80 transition-opacity">
                <i className="iconfont icon-add-circle text-black dark:text-white text-2xl"></i>
                <span className="text-black dark:text-white ml-1" style={{ marginTop: '-2px' }}>
                  New Note
                </span>
              </button>
            </Link>
          )}
          <div className="dropdown dropdown-end">
            <button className="flex flex-row items-center ml-5 opacity-60 hover:opacity-80 transition-opacity">
              <i className="iconfont icon-menu text-black dark:text-white text-2xl"></i>
              <span className="text-black dark:text-white ml-1" style={{ marginTop: '-2px' }}>
                Menu
              </span>
            </button>
            <ul
              tabIndex={0}
              className="dropdown-content z-[1] menu p-2 shadow bg-base-100 rounded-btn w-36 mt-2">
              <li>
                <Link to="/notes">
                  <div className="flex flex-row items-center opacity-60">
                    <i className="iconfont icon-Notes text-black dark:text-white text-2xl"></i>
                    <span className="text-black dark:text-white ml-2">All Notes</span>
                  </div>
                </Link>
              </li>
              {!hasPasscode() && (
                <li>
                  <Link to="/auth">
                    <div className="flex flex-row items-center opacity-60">
                      <i className="iconfont icon-user text-black dark:text-white text-2xl"></i>
                      <span className="text-black dark:text-white ml-2">Login</span>
                    </div>
                  </Link>
                </li>
              )}
              {hasPasscode() && (
                <li>
                  <Link to="" onClick={onLogout}>
                    <div className="flex flex-row items-center opacity-60">
                      <i className="iconfont icon-logout text-black dark:text-white text-2xl"></i>
                      <span className="text-black dark:text-white ml-2">Logout</span>
                    </div>
                  </Link>
                </li>
              )}
            </ul>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Navbar
