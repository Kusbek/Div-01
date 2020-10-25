import Guest from './guest/model.js'
import GuestView from './guest/view.js'
import GuestController from './guest/controller.js'

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


    constructor(chatModel, chatView) {
        this.model = chatModel
        this.view = chatView
    }

    monitorGuests() {
        this.model.connectToChatSocket(this.handleNewGuest)
    }

    handleNewGuest = (data) => {
        console.log(data)
    }

    closeChat = () => {
        this.model.closeChatSocket()
    }
}