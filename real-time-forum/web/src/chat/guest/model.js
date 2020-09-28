export default class Guest {
    constructor({id, nickname}) {
        this.id = id
        this.nickname = nickname
        this.getRoomUrl = "/room"
    }

    get = () => {
        return {
            id: this.id,
            nickname: this.nickname
        }
    }

    async getRoomNumberFromServer(newRoomHandler) {
        const room = await fetch(`${this.getRoomUrl}?guest_id=${this.id}`, {
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
            console.log(json.room)
            let room = newRoomHandler(json.room)
            return room
        }).catch((e) => {
            return Promise.reject(Error(e))
        })

        return room
    }
}