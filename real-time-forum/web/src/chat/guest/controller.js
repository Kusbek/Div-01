import Room from './room/model.js'
import RoomView from './room/view.js'
import RoomController from './room/controller.js'
export default class GuestController {
    constructor(closeAllRooms, newPost, userModel, model, view) {
        this.userModel = userModel
        this.model = model
        this.view = view
        this.newPost = newPost
        this.closeAllRooms = closeAllRooms
        this.displayGuest()
    }

    displayGuest = () => {
        this.view.display(this.model.get())
        this.view.bindHandleClick(this.handleClick)
    }


    handleClick = () => {
        this.closeAllRooms()
        this.model.getRoomNumberFromServer(this.handleNewRoom).then((room) => {
            this.room = room
        }).catch((error) => {
            console.log(error)
        })
    }

    handleNewRoom = (r) => {
        const room = new Room(r)
        const roomView = new RoomView()
        const roomController = new RoomController(this.newPost, this.userModel, room, roomView)

        return {
            model: room,
            view: roomView,
            controller: roomController,
        }
    }

    delete = () => {
        this.view.delete()
    }

    closeRoom = () => {
        if (this.room != undefined && this.room != null) {
            this.room.controller.closeSocket()
        }
    }
}