import { NavigateFunction } from 'react-router/dist/lib/hooks'

export const runAsyncFunction = async (fn: () => void) => {
  fn()
}

export function displayName(name: string) {
  return name.replace('.md', '')
}

export function isValidFileName(fileName: string) {
  const forbiddenChars = ['/', '\\', ':', '*', '?', '"', '<', '>', '|']

  for (const char of forbiddenChars) {
    if (fileName.includes(char)) {
      return false
    }
  }

  return !fileName.startsWith('.')
}

export function onSearch(keywords: string, navigate: NavigateFunction) {
  const searchWords = keywords.trim()

  if (searchWords === '*') {
    return navigate('/notes')
  }

  if (searchWords.length < 1) {
    return
  }

  navigate(`/search?keywords=${searchWords}`)
}
