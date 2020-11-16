

export default class Chat {
    constructor() {
        this.wsUrl = "ws://"+location.host+"/chat"
    }

    // monitorGuestsInServer(newGuestHandler, updateState) {
    //     const socket = new WebSocket(this.wsUrl + `?`)
    //     socket.onopen = () => {
    //         console.log("Opened Socket")
    //     }

    //     socket.onmessage = (e) => {
    //         let msg = JSON.parse(e.data)

    //         let action = msg.action

    //         if (msg.user.id in this.guests) {
    //             updateState(this.guests[msg.user.id], action)
    //         } else {
    //             let guest = newGuestHandler(msg)
    //             this.guests[msg.user.id] = guest
    //         }
    //     }

    //     socket.onclose = () => {
    //         console.log("Close Socket")
    //     }

    //     this.socket = socket
    // }

    // closeWS = () => {
    //     this.socket.close()
    // }

    // getGuests() {
    //     return this.guests
    // }

    // closeAllRooms = () => {
    //     for (const [id, guest] of Object.entries(this.guests)) {
    //         guest.controller.closeRoom()
    //     }
    // }


    connectToChatSocket = (newGuestHandler) => {
        const socket = new WebSocket(this.wsUrl)
        this.socket = socket
        socket.onmessage = (msg) => {
            let data = JSON.parse(msg.data)
            newGuestHandler(data)
        } 
    }

    closeChatSocket = () => {
        if (this.socket !== undefined) {
            this.socket.close()
        }
    }
}