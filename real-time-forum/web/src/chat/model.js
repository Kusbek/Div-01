

export default class Chat {
    constructor() {
        this.guests = {}
        this.wsUrl = "ws://localhost:8082/chat"
    }

    monitorGuestsInServer(newGuestHandler, updateState) {
        const socket = new WebSocket(this.wsUrl + `?`)
        socket.onopen = () => {
            console.log("Opened Socket")
        }

        socket.onmessage = (e) => {
            let msg = JSON.parse(e.data)

            let action = msg.action

            if (msg.user.id in this.guests) {
                updateState(this.guests[msg.user.id], action)
            } else {
                let guest = newGuestHandler(msg)
                this.guests[msg.user.id] = guest
            }
            
            // if (action == "online") {
            //     let guest = newGuestHandler(msg.user)
            //     this.guests[msg.user.id] = guest
            // } else {
            //     let g = this.guests[msg.user.id]
            //     deleteGuestHandler(g)
            //     delete this.guests[msg.user.id]
            // }
            // console.log(this.guests)
        }

        socket.onclose = () => {
            console.log("Close Socket")
        }

        this.socket = socket
    }

    closeWS = () => {
        this.socket.close()
    }

    getGuests() {
        return this.guests
    }

    closeAllRooms = () => {
        for (const [id, guest] of Object.entries(this.guests)) {
            guest.controller.closeRoom()
        }
    }
}