import React, { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import Navbar from '@/components/Navbar.tsx'
import { NoteType } from '@/types'
import useRequest from '@/hooks/useRequest.ts'
import useDebouncedValue from '@/hooks/useDebouncedValue.ts'
import LoadingView from '@/components/LoadingView.tsx'
import { displayName } from '@/utils'

let cachedNotes: NoteType[] = []

const NotesScreen: React.FC = () => {
  const [notes, setNotes] = useState<NoteType[]>(cachedNotes)
  const page = 1
  const limit = 1000
  const { data, isLoading } = useRequest<NoteType[]>({
    method: 'GET',
    url: `notes?page=${page}&limit=${limit}`
  })

  const loading = useDebouncedValue(false, isLoading, 1000)

  useEffect(() => {
    data && setNotes(data)
  }, [data])

  useEffect(() => {
    return () => {
      if (notes.length) {
        cachedNotes = notes
      }
    }
  }, [notes])

  return (
    <main className="w-full">
      <Navbar />
      <div className="container mx-auto px-5 mt-20 max-w-lg sm:max-w-6xl">
        {loading && <LoadingView />}
        {notes.map(note => (
          <div className="py-1">
            <Link to={`/notes/${note.id}`}>{displayName(note.name)}</Link>
          </div>
        ))}
        <footer className="h-10"></footer>
      </div>
    </main>
  )
}

export default NotesScreen
