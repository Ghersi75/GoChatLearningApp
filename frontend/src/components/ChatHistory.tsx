import React, { Dispatch, SetStateAction } from 'react'

export default function ChatHistory({
  chatHistory
}: {
  chatHistory: string[]
}) {
  return (
    <div className="flex justify-center items-center flex-col p-4 gap-4 w-screen ">
      <h2 className="text-lg font-bold"> Chat History </h2>
      {
        chatHistory.map((msg: any, index: number) => {
          const data = JSON.parse(msg.data)
          return(
            <p key={index} className="outline solid w-fit p-2.5">{data.body}</p>
          )
        })
      }
    </div>
  )
}
