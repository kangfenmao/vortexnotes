import React, { useRef, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { isAxiosError } from 'axios'
import { getAxiosInstance } from '@/config/http.ts'
import ToastPopup from '@/components/popups/ToastPopup.tsx'

const AuthScreen: React.FC = () => {
  const [passcode, setPasscode] = useState('')
  const [loading, setLoading] = useState(false)
  const navigate = useNavigate()
  const inputRef = useRef<HTMLInputElement>(null)
  const [error, setError] = useState(false)

  const onAuth = async () => {
    if (loading) {
      return
    }

    if (passcode.trim().length === 0) {
      return inputRef.current?.focus()
    }

    setLoading(true)

    try {
      const res = await window.$http.post('auth', { passcode })
      localStorage.vortexnotes_passcode = passcode
      localStorage.vortexnotes_auth_scope = res.data.auth_scope
      window.$http = getAxiosInstance()
      navigate('/')
      ToastPopup.show({ message: 'Authentication successful' })
    } catch (error) {
      setLoading(false)
      setError(true)

      if (isAxiosError(error)) {
        if (error.response) {
          setPasscode('')
          return ToastPopup.show({ message: error.response.data?.message })
        }
      }

      return ToastPopup.show({ message: 'Auth Failure' })
    }
  }

  const inputErrorClass = error && 'border-red-500 focus:border-red-500 placeholder-red-500'

  return (
    <main className="flex flex-1 flex-col justify-center items-center">
      <div className="relative flex flex-col items-center justify-center h-screen overflow-hidden w-96">
        <div className="w-full p-6 pt-4 bg-white bg-opacity-10 border-gray-200 dark:border-gray-600 rounded-md shadow-md border-top lg:max-w-lg border">
          <div className="flex flex-col items-center">
            <img
              src="/public/icon-200x200.png"
              className={`w-24 cursor-pointer ${loading && 'animate-spin'}`}
              alt=""
            />
            <span
              style={{ fontFamily: 'Major Mono Display' }}
              className="text-sm mb-5 mt-1 select-none">
              Vortexnotes
            </span>
          </div>
          <section className="relative overflow-hidden">
            <input
              type="password"
              placeholder={error ? 'Passcode Invalid' : 'Enter Passcode'}
              className={`w-full input input-bordered pr-20 ${inputErrorClass}`}
              value={passcode}
              ref={inputRef}
              onKeyDown={e => e.key === 'Enter' && onAuth()}
              onChange={e => {
                setPasscode(e.target.value)
                error && setError(false)
              }}
            />
            <button
              type="button"
              value="Input"
              onClick={onAuth}
              disabled={error}
              className="absolute right-1 top-1 bottom-1 w-10 rounded-md transition"
              style={{ borderTopLeftRadius: 0, borderBottomLeftRadius: 0 }}>
              <i className={`iconfont icon-arrow-right ${error && 'text-red-500'}`}></i>
            </button>
          </section>
        </div>
      </div>
    </main>
  )
}

export default AuthScreen
