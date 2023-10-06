"use client"
import ChatHistory from "@/components/ChatHistory"
import { connect, sendMsg } from "@/utils/api/websocketHandler"
import { useEffect, useState } from "react"

export default function Home() {
  const [msgHistory, setMsgHistory] = useState<string[]>([])
  const [msg, setMsg] = useState("")

  useEffect(() => {
    connect((msg: any) => {
      setMsgHistory(prev => {
        return [...prev, msg]
      })
    })
  }, [])

  const handleClick = () => {
    if (msg == "") {
      return
    }
    sendMsg(msg)
    setMsg("")
  }

  return (
    <main className="bg-slate-200">
      <div className="flex flex-col w-screen h-fit justify-center items-center">
        <div className="p-2.5 flex flex-col justify-center items-center gap-4">
          <input type="text p-2.5 bg-black" value={msg} onChange={e => setMsg(e.target.value)}></input>
          <button onClick={handleClick} className="bg-white w-fit h-fit rounded-lg p-2.5 text-xl">
            Send
          </button>
        </div>
        <ChatHistory chatHistory={msgHistory} />
      </div>
    </main>
  )
}
