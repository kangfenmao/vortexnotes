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
