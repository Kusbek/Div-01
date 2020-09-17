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
        this.model.getGuestsFromServer(this.handleNewGuest)
    }

    handleNewGuest = (guest) => {
        const guestModel = new Guest(guest)
        const guestView = new GuestView(this.view.chatWall)
        const guestController = new GuestController(guestModel, guestView)

        return {
            model: guestModel,
            view: guestView,
            controller: guestController,
        }
    }
}