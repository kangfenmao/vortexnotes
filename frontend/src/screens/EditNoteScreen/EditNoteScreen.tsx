import React, { useEffect, useRef, useState } from 'react'
import Navbar from '@/components/Navbar.tsx'
import { displayName, isValidFileName } from '@/utils'
import { isAxiosError } from 'axios'
import { useBlocker, useNavigate, useParams } from 'react-router-dom'
import useRequest from '@/hooks/useRequest.ts'
import { NoteType } from '@/types'
import MDEditor from '@uiw/react-md-editor'
import AlertPopup from '@/components/popups/AlertPopup.tsx'

const EditNoteScreen: React.FC = () => {
  const params = useParams()
  const id = params.id
  const sessionNote = JSON.parse(sessionStorage.getItem(`EDIT_NOTE:${id}`)!)
  const [title, setTitle] = useState(sessionNote?.name || '')
  const [content, setContent] = useState(sessionNote?.content || '')
  const [edited, setEdited] = useState(false)
  const [saving, setSaving] = useState(false)
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

  const blocker = useBlocker(
    ({ currentLocation, nextLocation }) =>
      !saving && edited && currentLocation.pathname !== nextLocation.pathname
  )

  useEffect(() => {
    if (blocker.state === 'blocked') {
      if (confirm('You have unsaved edits. Are you sure you want to leave?')) {
        blocker.proceed()
        navigate(location.pathname)
        return
      }
      blocker.reset()
    }
  }, [blocker, navigate])

  const onEdit = () => !edited && setEdited(true)

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
      return AlertPopup.show({
        title: 'Message',
        content: 'The title cannot contain special characters.'
      })
    }

    try {
      setSaving(true)

      const res = await window.$http.patch(`notes/${id}`, {
        name: displayName(title.trim()),
        content: content.trim()
      })

      if (res.data?.id) {
        return navigate(`/notes/${res.data.id}`, { replace: true })
      }

      history.back()
    } catch (error) {
      setSaving(false)
      if (isAxiosError(error)) {
        if (error.response) {
          return AlertPopup.show({
            title: 'Save failed',
            content: error.response.data.message
          })
        }
      }
    }
  }

  return (
    <main className="w-full">
      <Navbar />
      <div className="container mx-auto px-5 pt-20 max-w-lg sm:max-w-6xl flex flex-col flex-1">
        <div className="flex flex-row items-center mb-4">
          <input
            className="text-2xl sm:text-2xl w-full font-bold line-clamp-1 bg-transparent outline-none"
            placeholder="Title"
            name="title"
            value={displayName(title)}
            ref={titleInputRef}
            tabIndex={1}
            onChange={e => setTitle(e.target.value)}
            onKeyDown={onEdit}
          />
          <button
            tabIndex={4}
            className="p-1 px-2 transition-all rounded-md flex flex-row items-center opacity-70 hover:bg-gray-200 dark:hover:bg-zinc-900 hover:opacity-90"
            onClick={onCancel}>
            <i className="iconfont icon-return text-2xl mr-1"></i>
            <span className="hidden sm:inline">Cancel</span>
          </button>
          <button
            className="p-1 px-2 transition-all rounded-md flex flex-row items-center opacity-70 hover:opacity-100 hover:bg-green-800 hover:text-white"
            onClick={onSave}
            tabIndex={3}>
            <i className="iconfont icon-editsaved text-2xl mr-1"></i>
            <span className="hidden sm:inline">Save</span>
          </button>
        </div>
        <MDEditor
          value={content}
          onChange={v => setContent(v!)}
          tabIndex={2}
          placeholder="Note..."
          autoFocus
          ref={contentInputRef}
          height={window.innerHeight - 180}
          preview={window.innerWidth < 500 ? 'edit' : 'live'}
          onKeyDown={onEdit}
        />
        <footer className="h-5"></footer>
      </div>
    </main>
  )
}

export default EditNoteScreen
