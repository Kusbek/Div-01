export default class Room {
    // constructor({ id }) {
    //     this.id = id
    //     this.next = 0
    //     this.msgUrl = "/messages"
    //     this.wsUrl = "ws://localhost:8082/message"
    // }

    // getMessagesFromServer = async () => {
    //     const messages = await fetch(`${this.msgUrl}?room_id=${this.id}&page_num=${this.next}`, {
    //         method: "GET",
    //     }).then((response) => {
    //         if (!response.ok) {
    //             return Promise.reject(Error(response.statusText))
    //         }
    //         return response.json()
    //     }).then((json) => {
    //         if (json.error != null || json.error != undefined) {
    //             return Promise.reject(Error(json.error))
    //         }
    //         this.next = json.next
    //         let messages = json.messages
    //         return messages
    //     }).catch((e) => {
    //         return Promise.reject(Error(e))
    //     })
    //     return messages
    // }

    // monitorMessages = (newMessageHandler) => {
    //     const socket = new WebSocket(this.wsUrl + `?room_id=${this.id}`)
    //     socket.onopen = () => {
    //         console.log("Opened Room")
    //     }

    //     socket.onmessage = (e) => {
    //         let msg = JSON.parse(e.data)

    //         newMessageHandler(msg.message, true)
    //         // console.log(this.guests)
    //     }

    //     socket.onclose = () => {
    //         console.log("Close Room")
    //     }

    //     this.socket = socket
    // }

    // sendMessage = (message) => {
    //     this.socket.send(JSON.stringify({ message: message }))
    // }

    // closeSocket = () => {
    //     this.socket.close()
    // }

    constructor({ id }) {
        this.id = id
        this.next = 0
        this.wsUrl = "ws://localhost:8080/message"
        this.msgUrl = "/messages"
    }

    getRoom = () => {
        return {
            id: this.id
        }
    }

    sendMessage = (message) => {
        this.socket.send(JSON.stringify({ message: message }))
    }

    monitorSocket = (newMessageHandler) => {
        const socket = new WebSocket(this.wsUrl + `?room_id=${this.id}`)
        this.socket = socket
        socket.onmessage = (e) => {
            let msg = JSON.parse(e.data)
            newMessageHandler(msg.message, true)
        }

        socket.onclose = () => {
            console.log(`close room ${this.id}`)
        }
    }

    closeSocket = () => {
        this.socket.close()
    }


    getMessages = async () => {
        const messages = await fetch(`${this.msgUrl}?room_id=${this.id}&page_num=${this.next}`, {
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
            this.next = json.next
            let messages = json.messages
            return messages
        }).catch((e) => {
            return Promise.reject(Error(e))
        })
        return messages
    }
}