import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import HomeScreen from '@/screens/HomeScreen'
import React, { useEffect, useState } from 'react'
import SearchScreen from '@/screens/SearchScreen'
import NoteScreen from '@/screens/NoteScreen'
import NewNoteScreen from '@/screens/NewNoteScreen'
import EditNoteScreen from '@/screens/EditNoteScreen'
import '@uiw/react-md-editor/markdown-editor.css'
import '@uiw/react-markdown-preview/markdown.css'
import NotesScreen from '@/screens/NotesScreen'
import TopViewContainer from '@/components/TopView.tsx'
import Root from '@/components/Root.tsx'
import AuthScreen from '@/screens/AuthScreen'
import LoadingView from '@/components/LoadingView.tsx'
import { fetchConfig } from '@/utils/api.ts'

const router = createBrowserRouter([
  {
    path: '/',
    element: <Root />,
    children: [
      {
        index: true,
        element: <HomeScreen />
      },
      {
        path: '/search',
        element: <SearchScreen />
      },
      {
        path: '/notes',
        element: <NotesScreen />
      },
      {
        path: '/notes/:id',
        element: <NoteScreen />
      },
      {
        path: '/notes/:id/edit',
        element: <EditNoteScreen />
      },
      {
        path: '/new',
        element: <NewNoteScreen />
      }
    ]
  },
  {
    path: '/auth',
    element: <AuthScreen />
  }
])

const App: React.FC = () => {
  const [ready, setReady] = useState(false)

  useEffect(() => {
    if (localStorage.vortexnotes_auth_scope) {
      setReady(true)
      fetchConfig()
    } else {
      fetchConfig().then(() => setReady(true))
    }
  }, [])

  if (!ready) {
    return (
      <main className="flex flex-1 justify-center items-center">
        <LoadingView />
      </main>
    )
  }

  return (
    <TopViewContainer>
      <RouterProvider router={router} />
    </TopViewContainer>
  )
}

export default App
