import { useState, useEffect } from 'react'

function useDebouncedValue(initialValue: any, debounceValue: any, delay: number) {
  const [value, setValue] = useState(initialValue)

  useEffect(() => {
    const timer = setTimeout(() => setValue(debounceValue), delay)
    return () => clearTimeout(timer)
  }, [debounceValue, delay])

  return value
}

export default useDebouncedValue
