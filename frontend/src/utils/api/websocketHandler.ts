const socket = new WebSocket("ws://localhost:8080/ws")

export function connect(addToMsgHistory: any) {
  console.log("Attempting Connection...")

  socket.onopen = () => {
    console.log("Successfully Connected")
  }

  socket.onmessage = msg => {
    console.log(msg)
    addToMsgHistory(msg)
  }

  socket.onclose = event => {
    console.log("Server Closed Connection: ", event)
  }

  socket.onerror = err => {
    console.log("Socket Error: ", err)
  }
}

export function sendMsg(msg: string) {
  console.log("Sending Message: ", msg)
  socket.send(msg)
}