import React, { useRef, useState } from 'react'
import Navbar from '@/components/Navbar.tsx'
import { isValidFileName } from '@/utils'
import { isAxiosError } from 'axios'
import { useNavigate } from 'react-router-dom'
import MDEditor from '@uiw/react-md-editor'

const NewNoteScreen: React.FC = () => {
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')

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
      const res = await window.$http.post('notes/new', {
        name: title.trim(),
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
      <div className="container mx-auto px-5 mt-20 max-w-lg sm:max-w-6xl">
        <div className="flex flex-row items-center mb-4">
          <input
            className="text-2xl sm:text-3xl w-full font-bold line-clamp-1 bg-transparent outline-none"
            placeholder="Title"
            autoFocus
            name="title"
            value={title}
            ref={titleInputRef}
            tabIndex={1}
            onChange={e => setTitle(e.target.value)}
          />
          <button
            tabIndex={4}
            className="p-1 px-3 hover:bg-zinc-900 transition-all rounded-md flex flex-row items-center opacity-70 hover:opacity-100"
            onClick={onCancel}>
            <i className="iconfont icon-return text-2xl mr-1"></i>
            Cancel
          </button>
          <button
            className="p-1 px-3 hover:bg-green-800 transition-all rounded-md flex flex-row items-center opacity-70 hover:opacity-100"
            onClick={onSave}
            tabIndex={3}>
            <i className="iconfont icon-editsaved text-2xl mr-1"></i>
            Save
          </button>
        </div>
        <MDEditor
          value={content}
          onChange={v => setContent(v!)}
          tabIndex={2}
          placeholder="Note..."
          autoFocus
          ref={contentInputRef}
          height={700}
        />
        <footer className="h-5"></footer>
      </div>
    </main>
  )
}

export default NewNoteScreen
