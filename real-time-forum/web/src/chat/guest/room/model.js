export default class Room {
    constructor({id}) {
        this.id = id
        this.wsUrl = "ws://localhost:8082/message"
    }

     getMessagesFromServer = async () => {
        const messages = await fetch(`${this.msgUrl}?room_id=${this.id}`, {
            method: "GET",
        }).then((response) => {
            if (!response.ok) {
                return Promise.reject(Error(response.statusText))
            }
            return response.json()
        }).then((json) => {
            if (json.error != null || json.error != undefined) {
                return Promise.reject(Error(json.error))
            }
            let messages = json.messages
            return messages
        }).catch((e) => {
            return Promise.reject(Error(e))
        })
        return messages
    }

    monitorMessages = (newMessageHandler)=> {
        const socket = new WebSocket(this.wsUrl + `?room_id=${this.id}`)
        socket.onopen = () => {
            console.log("Opened Room")
        }

        socket.onmessage = (e) => {
            let msg = JSON.parse(e.data)
            
            newMessageHandler(msg)
            // console.log(this.guests)
        }

        socket.onclose = () => {
            console.log("Close Room")
        }

        this.socket = socket
    }

    sendMessage = (message) => {
        this.socket.send(JSON.stringify({message:message}))
    }
}