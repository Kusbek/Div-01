import Guest from '../guest/model.js'
import GuestView from '../guest/view.js'
import GuestController from '../guest/controller.js'

export default class ChatController {
    // constructor(newPost, userModel, model, view) {
    //     this.userModel = userModel
    //     this.model = model
    //     this.view = view
    //     this.initGuestsSocket()
    //     this.newPost = newPost
    // }


    // delete() {
    //     this.view.delete()
    //     this.model.closeWS()
    // }


    // initGuestsSocket = () => {
    //     this.model.monitorGuestsInServer(this.handleNewGuest, this.updateState)
    // }

    // closeAllRooms = () => {
    //     this.model.closeAllRooms()
    // }

    // handleNewGuest = (msg) => {
    //     const guestModel = new Guest(msg)
    //     const guestView = new GuestView(this.view.chatWall)
    //     const guestController = new GuestController(this.closeAllRooms, this.newPost, this.userModel, guestModel, guestView)
    //     return {
    //         model: guestModel,
    //         view: guestView,
    //         controller: guestController,
    //     }
    // }

    // handleDeleteGuest = (guest) => {
    //     guest.controller.delete()
    // }

    // updateState = (guest, action) => {
    //     guest.controller.updateState(action)
    // }


    constructor(chatModel, chatView, createRoom) {
        this.model = chatModel
        this.view = chatView
        this.guests = []
        this.createRoom = createRoom
    }

    monitorGuests() {
        this.model.connectToChatSocket(this.handleNewGuest)
    }

    handleNewGuest = (data) => {
        let existingGuest = this.filterGuestsByID(data.user.id)
        if (existingGuest === undefined) {
            let guest = new GuestController(new Guest(data), new GuestView(), this.createRoom)
            this.guests.push(guest)
        }else {
            existingGuest.update(data)
        }
        this.updateChat()
    }

    filterGuestsByID = (id) => {
        let guest = this.guests.filter(g => g.getId() == id)
        if (guest.length === 0) {
            return undefined
        }
        return guest[0]
    }

    clearChat = () => {
        this.view.clearChat()
        this.guests = []
    }

    updateChat = () => {
        this.view.clearChat()
        this.view.addHeaders()
        this.updateGuests("offline")
        this.updateGuests("online")
    }

    updateGuests = (status) => {
        let guests = this.guests.filter(guest => guest.model.status == status)
        let [mg, ng] = this.partitionByLastMessage(guests)
        // console.log(mg, ng)
        this.sortByLastMessage(mg)
        this.sortInAlphabetical(ng)
        guests = mg.concat(ng)
        guests.forEach(g => g.display())
    }

    sortByLastMessage = (guests) => {
        guests.sort((g1, g2) => {
            return new Date(g2.model.lastMessage) - new Date(g1.model.lastMessage)
        })
    }

    sortInAlphabetical = (guests) => {
        guests.sort((g1, g2) => {
            if (g1.model.user.nickname < g2.model.user.nickname) { return -1; }
            if (g1.model.user.nickname > g2.model.user.nickname) { return 1; }
            return 0;
        })
    }

    partitionByLastMessage = (guests) => {
        let messageGuests = []
        let noMessageGuests = []
        guests.forEach((guest) => {
            if (guest.model.haveMessages()) {
                messageGuests.push(guest)
            } else {
                noMessageGuests.push(guest)
            }
        })

        return [messageGuests, noMessageGuests]
    }

    closeChat = () => {
        this.model.closeChatSocket()
    }
}