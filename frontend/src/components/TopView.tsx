import React, { useCallback, useEffect, useState } from 'react'
import { findIndex, pullAt } from 'lodash-es'
import EventEmitter from 'eventemitter3'

const DeviceEventEmitter = new EventEmitter()

let id = 0

interface Props {
  children?: React.ReactNode
}

type ElementItem = {
  key: number
  element: React.FC | React.ReactNode
}

const TopViewContainer: React.FC<Props> = ({ children }) => {
  const [elements, setElements] = useState<ElementItem[]>([])

  const pop = useCallback(() => {
    const views = [...elements]
    views.pop()
    setElements(views)
  }, [elements])

  useEffect(() => {
    const listeners = [
      DeviceEventEmitter.addListener('ADD_OVERLAY', ({ element, key }) => {
        setElements(elements.concat([{ element, key }]))
      }),
      DeviceEventEmitter.addListener('REMOVE_OVERLAY', ({ key }: { key: number }) => {
        const views = [...elements]
        pullAt(views, findIndex(views, { key }))
        setElements(views)
      }),
      DeviceEventEmitter.addListener('POP_OVERLAY', pop)
    ]

    return () => listeners.forEach(listener => listener.removeAllListeners())
  }, [pop, elements])

  return (
    <div className="flex flex-1 absolute w-full h-full">
      {children}
      {elements.length > 0 && (
        <div className="flex flex-1 absolute w-full h-full">
          <div className="absolute w-full h-full" onClick={pop} />
          {elements.map(({ element: Element, key }) => (
            <div key={`TOPVIEW_${key}`}>
              {typeof Element === 'function' ? <Element /> : Element}
            </div>
          ))}
        </div>
      )}
    </div>
  )
}

export const TopView = {
  show: (element: React.FC | React.ReactNode) => {
    id = id + 1
    DeviceEventEmitter.emit('ADD_OVERLAY', { element, key: id })
    return id
  },
  hide: (key: number) => {
    DeviceEventEmitter.emit('REMOVE_OVERLAY', { key })
  },
  pop: () => {
    DeviceEventEmitter.emit('POP_OVERLAY')
  }
}

export default TopViewContainer
