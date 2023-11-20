import React from 'react'
import { Link } from 'react-router-dom'

const NoteScreen: React.FC = () => {
  return (
    <main className="container mx-auto max-w-3xl px-4">
      <div className="flex flex-row h-20 items-center">
        <Link to="/">
          <h1
            className="text-4xl font-bold text-white"
            style={{ fontFamily: 'Major Mono Display' }}>
            <span className="text-violet-700">V</span>
            <span className="text-red-500">o</span>rtex
          </h1>
        </Link>
      </div>
    </main>
  )
}

export default NoteScreen
