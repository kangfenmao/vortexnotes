export const runAsyncFunction = async (fn: () => void) => {
  fn()
}

export function displayName(name: string) {
  return name.replace('.md', '')
}
