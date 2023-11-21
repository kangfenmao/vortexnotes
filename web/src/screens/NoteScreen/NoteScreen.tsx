import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { displayName, runAsyncFunction } from '@/utils'
import Markdown from 'react-markdown'
import remarkGfm from 'remark-gfm'
import './index.css'
import Navbar from '@/components/Navbar.tsx'

const NoteScreen: React.FC = () => {
  const [note, setNote] = useState<any>()
  const params = useParams()
  const id = params.id

  useEffect(() => {
    runAsyncFunction(async () => {
      const res = await window.$http.get(`notes/${id}`)
      setNote(res.data)
    })
  }, [])

  return (
    <main className="w-full">
      <Navbar />
      <div className="container mx-auto px-5 mt-24 max-w-lg sm:max-w-6xl">
        {note && (
          <>
            <h1 className="text-2xl sm:text-3xl mb-5 font-bold line-clamp-1">
              {displayName(note.name)}
            </h1>
            <Markdown className="markdown-body" remarkPlugins={[remarkGfm]}>
              {note.content}
            </Markdown>
          </>
        )}
        <footer className="h-10"></footer>
      </div>
    </main>
  )
}

export default NoteScreen
