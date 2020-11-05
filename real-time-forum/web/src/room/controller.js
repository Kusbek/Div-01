export default class RoomController {
    // constructor(guestController, newPost, userModel,room, roomView) {
    //     // this.guestController = guestController
    //     // this.newPost = newPost
    //     // this.userModel = userModel
    //     // this.model = room
    //     // this.view = roomView
    //     // this.view.displayRoom()
    //     // this.view.bindGetOldMessagesScrool(this.handleScrollToCeil)
    //     // this.view.bindCloseButton(this.handleClose)
    //     // this.view.bindSendMessageButton(this.handleSendMessage)
    //     // this.monitorNewMessages()
    //     // this.getMessages()
    // }

    // monitorNewMessages = () => {
    //     this.model.monitorMessages(this.newMessageHandler)
    // }

    // getMessages = () => {
    //     this.model.getMessagesFromServer().then((msgs)=>{
    //         msgs.forEach((msg)=>{

    //             this.newMessageHandler(msg,false)
    //         })
    //     }
    //     ).catch((e) => {
    //         alert(e)
    //     })
    // }

    // handleScrollToCeil = () => {
    //     this.getMessages()
    // }

    // newMessageHandler = (msg, recent) => {
    //     let self = false
    //     if (msg.user.id == this.userModel.getUser().id){
    //         self = true
    //     }
    //     this.view.displayMessage(msg, self, recent)
    // }

    // handleClose = () => {
    //     this.view.close()
    //     this.closeSocket()
    //     this.newPost(this.userModel)
    // }

    // handleSendMessage = (msg) => {
    //     this.guestController.toTop()
    //     this.model.sendMessage(msg)
    // }

    // closeSocket = () => {
    //     this.model.closeSocket()
    // }

    constructor(room, roomView, userController, closeRoom) {
        this.model = room
        this.view = roomView
        this.user = userController
        this.monitorRoomSocket()
        this.displayHistory()
        this.closeRoom = closeRoom
    }

    monitorRoomSocket = () => {
        this.model.monitorSocket(this.handleNewMessage)
    }


    displayHistory = () => {
        this.model.getMessages().then((msgs)=> {
            msgs.forEach((msg)=> {
                this.handleNewMessage(msg, false)
            })
        })
    }

    displayRoom = () => {
        this.view.display(this.model.getRoom(), this.handleSendMessageButtonClick, this.displayHistory,this.closeRoom)
    }

    handleNewMessage = (data, recent) => {
        if (data.user.id === this.user.getUser().id) {
            this.view.displayMsg(data, true, recent)
        } else {
            this.view.displayMsg(data, false, recent)
        }
    }

    handleSendMessageButtonClick = (msg) => {
        this.model.sendMessage(msg)
    }

    close = () => {
        this.model.closeSocket()
    }
}