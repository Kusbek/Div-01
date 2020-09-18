import Guest from './guest/model.js'
import GuestView from './guest/view.js'
import GuestController from './guest/controller.js'

export default class ChatController {
    constructor(model, view) {
        this.model = model
        this.view = view
        this.initGuestsSocket()
    }


    initGuestsSocket = () => {
        this.model.monitorGuestsInServer(this.handleNewGuest, this.handleDeleteGuest)
        
    }

    handleNewGuest = (user) => {
        const guestModel = new Guest(user)
        const guestView = new GuestView(this.view.chatWall)
        const guestController = new GuestController(guestModel, guestView)
        return {
            model: guestModel,
            view: guestView,
            controller: guestController,
        }
    }

    handleDeleteGuest = (guest) => {
        guest.controller.delete()
    }
}