import { useEffect, useState } from 'react'

export default function useTheme() {
  const [theme, setTheme] = useState(localStorage.theme || 'dark')

  useEffect(() => {
    localStorage.theme = theme
    document.documentElement.setAttribute('data-color-mode', theme)
    document.documentElement.setAttribute('data-theme', theme)
  }, [theme])

  return [theme, setTheme]
}
