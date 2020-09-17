

export default class Chat {
    constructor() {
        this.guests = []
    }

    getGuestsFromServer(newGuestHandler) {
        const guests = [
            {
                id: 2,
                nickname: "guest2"
            },
            {
                id: 3,
                nickname: "guest3"
            },    
            {
                id: 4,
                nickname: "guest4"
            },          
        ]

        for (let guest of guests) {
            this.guests.push(newGuestHandler(guest))
        }

        return this.getGuests()
    }

    getGuests(){
        return this.guests
    }

    
}