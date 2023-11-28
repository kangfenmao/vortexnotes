import { TopView } from '@/components/TopView.tsx'
import React, { useEffect } from 'react'

interface Props {
  params: AlertShowParams
}

let topViewKey = 0

const AlertContainer: React.FC<Props> = ({ params }) => {
  const onClosed = async () => {
    setTimeout(() => TopView.hide(topViewKey), 300)
  }

  useEffect(() => {
    const modal: any = document.getElementById('alert_popup_modal')!
    modal.showModal()
  }, [])

  return (
    <dialog id="alert_popup_modal" className="modal">
      <div className="modal-box">
        <h3 className="font-bold text-lg">{params.title}</h3>
        <p className="py-4">{params.content}</p>
        <div className="modal-action">
          <form method="dialog">
            <button className="btn" onClick={onClosed}>
              {params.buttonOkText || 'OK'}
            </button>
          </form>
        </div>
      </div>
    </dialog>
  )
}

interface AlertShowParams {
  title: string
  content: string | React.ReactNode
  buttonOkText?: string
}

const AlertPopup = {
  show: (params: AlertShowParams) => {
    topViewKey = TopView.show(<AlertContainer params={params} />)
  }
}

export default AlertPopup
