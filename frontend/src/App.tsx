import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import HomeScreen from '@/screens/HomeScreen'
import React from 'react'
import SearchScreen from '@/screens/SearchScreen'
import NoteScreen from '@/screens/NoteScreen'
import NewNoteScreen from '@/screens/NewNoteScreen'
import EditNoteScreen from '@/screens/EditNoteScreen'
import '@uiw/react-md-editor/markdown-editor.css'
import '@uiw/react-markdown-preview/markdown.css'
import NotesScreen from '@/screens/NotesScreen'
import TopViewContainer from '@/components/TopView.tsx'
import Root from '@/components/Root.tsx'

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
  }
])

const App: React.FC = () => {
  return (
    <TopViewContainer>
      <RouterProvider router={router} />
    </TopViewContainer>
  )
}

export default App
