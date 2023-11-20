import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import HomeScreen from '@/screens/HomeScreen'
import React from 'react'
import SearchScreen from '@/screens/SearchScreen'
import NoteScreen from '@/screens/NoteScreen'

const router = createBrowserRouter([
  {
    path: '/',
    element: <HomeScreen />
  },
  {
    path: '/search',
    element: <SearchScreen />
  },
  {
    path: '/notes/:id',
    element: <NoteScreen />
  }
])

const App: React.FC = () => {
  return <RouterProvider router={router} />
}

export default App
