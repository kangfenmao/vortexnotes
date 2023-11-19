export const runAsyncFunction = async (fn: () => void) => {
  await fn()
}
