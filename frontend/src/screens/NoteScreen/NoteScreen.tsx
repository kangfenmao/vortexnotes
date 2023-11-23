import React, { useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { displayName, runAsyncFunction } from '@/utils'
import Markdown from 'react-markdown'
import remarkGfm from 'remark-gfm'
import './index.css'
import Navbar from '@/components/Navbar.tsx'
import { isAxiosError } from 'axios'
import { NoteType } from '@/types'

const NoteScreen: React.FC = () => {
  const [note, setNote] = useState<NoteType>()
  const params = useParams()
  const id = params.id
  const navigate = useNavigate()

  useEffect(() => {
    runAsyncFunction(async () => {
      const res = await window.$http.get(`notes/${id}`)
      setNote(res.data)
    })
  }, [id])

  const onDelete = async () => {
    if (!confirm(`Delete note ${note?.name}?`)) {
      return
    }

    try {
      await window.$http.delete(`notes/${id}`)
      navigate('/')
    } catch (error) {
      if (isAxiosError(error)) {
        if (error.response) {
          return alert('Delete note error: ' + error.response.data?.message)
        }
      }
      return alert('Delete note error')
    }
  }

  const onEdit = () => {}

  return (
    <main className="w-full">
      <Navbar />
      <div className="container mx-auto px-5 mt-24 max-w-lg sm:max-w-6xl">
        {note && (
          <>
            <div className="flex flex-row items-center mb-5">
              <h1 className="flex-1 text-2xl sm:text-3xl font-bold line-clamp-1">
                {displayName(note.name)}
              </h1>
              <button
                tabIndex={4}
                className="p-2 px-4 hover:bg-zinc-900 transition-all rounded-sm"
                onClick={onEdit}>
                Edit
              </button>
              <button
                className="p-2 px-4 text-red-400 hover:bg-red-500 hover:text-white transition-all rounded-sm"
                onClick={onDelete}
                tabIndex={3}>
                Delete
              </button>
            </div>
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
