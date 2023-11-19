import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import HomeScreen from '@/screens/HomeScreen'
import React from 'react'
import SearchScreen from '@/screens/SearchScreen'

const router = createBrowserRouter([
  {
    path: '/',
    element: <HomeScreen />
  },
  {
    path: '/search',
    element: <SearchScreen />
  }
])

const App: React.FC = () => {
  return <RouterProvider router={router} />
}

export default App
