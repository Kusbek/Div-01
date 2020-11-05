
export default class GuestController {
    // constructor(closeAllRooms, newPost, userModel, model, view) {
    //     this.userModel = userModel
    //     this.model = model
    //     this.view = view
    //     this.newPost = newPost
    //     this.closeAllRooms = closeAllRooms
    //     this.displayGuest()
    // }

    // updateState = (action) => {
    //     this.model.action = action
    //     this.view.updateState(this.model.get())
    // }

    // displayGuest = () => {
    //     this.view.display(this.model.get())
    //     this.view.bindHandleClick(this.handleClick)
    // }


    // handleClick = () => {
    //     this.closeAllRooms()
    //     this.model.getRoomNumberFromServer(this.handleNewRoom).then((room) => {
    //         this.room = room
    //     }).catch((error) => {
    //         console.log(error)
    //     })
    // }

    // handleNewRoom = (r) => {
    //     const room = new Room(r)
    //     const roomView = new RoomView()
    //     const roomController = new RoomController(this,this.newPost, this.userModel, room, roomView)

    //     return {
    //         model: room,
    //         view: roomView,
    //         controller: roomController,
    //     }
    // }

    // delete = () => {
    //     this.view.delete()
    // }

    // closeRoom = () => {
    //     if (this.room != undefined && this.room != null) {
    //         this.room.controller.closeSocket()
    //     }
    // }

    // toTop = () => {
    //     this.view.toTop()
    // }

    constructor(guest, guestView, createRoom) {
        this.model = guest
        this.view = guestView
        this.createRoom = createRoom
    }

    display = () => {
        this.view.display(this.model.getGuest(), this.guestClickHandler)
    }

    guestClickHandler = () => {
        this.view.makeGray()
        this.model.setNewMessage(false)
        let g = this.model.getGuest()
        this.model.getRoom().then(room => {
            this.model.setRoom(room)
            this.createRoom(this.model.room)
        })
    }

    getId = () => {
        return this.model.getGuest().user.id
    }

    update = ({ user, last_message, status, new_message }) => {
        // console.log(user, lastMessage, status)
        if (last_message !== undefined && last_message !== null) {
            this.model.setLastMessage(last_message)
        }

        if (status !== undefined && status !== null && status !== "") {
            this.model.setStatus(status)
        }

        if (new_message) {
            this.model.setNewMessage(new_message)
        }
    }
}