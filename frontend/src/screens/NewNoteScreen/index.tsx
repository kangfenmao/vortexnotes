import React, { useEffect, useRef, useState } from 'react'
import { isValidFileName } from '@/utils'
import { useBlocker, useNavigate } from 'react-router-dom'
import MDEditor from '@uiw/react-md-editor'
import { isAxiosError } from 'axios'
import AlertPopup from '@/components/popups/AlertPopup.tsx'
import usePermissionCheckEffect from '@/hooks/usePermissionCheckEffect.ts'

const NewNoteScreen: React.FC = () => {
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')
  const [saving, setSaving] = useState(false)
  const navigate = useNavigate()

  const titleInputRef = useRef<HTMLInputElement>(null)
  const contentInputRef = useRef<HTMLTextAreaElement>(null)

  const isEdited = title !== '' || content !== ''
  const blocker = useBlocker(
    ({ currentLocation, nextLocation }) =>
      !saving && isEdited && currentLocation.pathname !== nextLocation.pathname
  )

  usePermissionCheckEffect('create')

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

      const res = await window.$http.post('notes/new', {
        name: title.trim(),
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
      <div className="container mx-auto px-5 pt-20 max-w-8xl">
        <div className="flex flex-row items-center mb-4">
          <input
            className="text-2xl sm:text-2xl w-full font-bold line-clamp-1 bg-transparent outline-none"
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
        />
        <footer className="h-5"></footer>
      </div>
    </main>
  )
}

export default NewNoteScreen
