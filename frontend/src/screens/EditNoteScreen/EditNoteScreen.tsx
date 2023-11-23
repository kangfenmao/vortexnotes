import React, { useEffect, useRef, useState } from 'react'
import Navbar from '@/components/Navbar.tsx'
import { displayName, isValidFileName } from '@/utils'
import { isAxiosError } from 'axios'
import { useNavigate, useParams } from 'react-router-dom'
import useRequest from '@/hooks/useRequest.ts'
import { NoteType } from '@/types'

const EditNoteScreen: React.FC = () => {
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')
  const params = useParams()
  const id = params.id

  const { data } = useRequest<NoteType>({ method: 'GET', url: `notes/${id}` })

  useEffect(() => {
    if (data) {
      setTitle(displayName(data.name))
      setContent(data.content)
    }
  }, [data])

  const titleInputRef = useRef<HTMLInputElement>(null)
  const contentInputRef = useRef<HTMLTextAreaElement>(null)

  const navigate = useNavigate()

  const onCancel = () => {
    history.back()
  }

  const onSave = async () => {
    if (title.trim().length === 0) {
      return titleInputRef.current?.focus()
    }

    if (content.trim().length === 0) {
      return contentInputRef.current?.focus()
    }

    if (!isValidFileName(title)) {
      return alert('The title cannot contain special characters.')
    }

    try {
      const res = await window.$http.patch(`notes/${id}`, {
        name: displayName(title.trim()),
        content: content.trim()
      })

      if (res.data?.id) {
        return navigate(`/notes/${res.data.id}`)
      }

      history.back()
    } catch (error) {
      if (isAxiosError(error)) {
        if (error.response) {
          alert('Save failed: ' + error.response.data.message)
        }
      }
    }
  }

  return (
    <main className="w-full">
      <Navbar />
      <div className="container mx-auto px-5 mt-24 max-w-lg sm:max-w-6xl">
        <div className="flex flex-row items-center mb-5">
          <input
            className="text-2xl sm:text-3xl w-full font-bold line-clamp-1 bg-transparent outline-none"
            placeholder="Title"
            name="title"
            value={title}
            ref={titleInputRef}
            tabIndex={1}
            onChange={e => setTitle(e.target.value)}
          />
          <button
            tabIndex={4}
            className="p-2 px-4 hover:bg-zinc-900 transition-all rounded-sm"
            onClick={onCancel}>
            Cancel
          </button>
          <button
            className="p-2 px-4 hover:bg-zinc-900 transition-all rounded-sm"
            onClick={onSave}
            tabIndex={3}>
            Save
          </button>
        </div>
        <textarea
          className="w-full p-5"
          rows={25}
          placeholder="Note..."
          name="content"
          autoFocus
          value={content}
          ref={contentInputRef}
          tabIndex={2}
          onChange={e => setContent(e.target.value)}></textarea>
        <footer className="h-10"></footer>
      </div>
    </main>
  )
}

export default EditNoteScreen
