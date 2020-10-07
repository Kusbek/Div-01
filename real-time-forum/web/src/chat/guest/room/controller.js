export default class RoomController {
    constructor(newPost, userModel,room, roomView) {
        this.newPost = newPost
        this.userModel = userModel
        this.model = room
        this.view = roomView
        this.view.displayRoom()
        this.view.bindGetOldMessagesScrool(this.handleScrollToCeil)
        this.view.bindCloseButton(this.handleClose)
        this.view.bindSendMessageButton(this.handleSendMessage)
        this.monitorNewMessages()
        this.getMessages()
    }

    monitorNewMessages = () => {
        this.model.monitorMessages(this.newMessageHandler)
    }

    getMessages = () => {
        this.model.getMessagesFromServer().then((msgs)=>{
            msgs.forEach((msg)=>{

                this.newMessageHandler(msg,false)
            })
        }
        ).catch((e) => {
            alert(e)
        })
    }

    handleScrollToCeil = () => {
        this.getMessages()
    }

    newMessageHandler = (msg, recent) => {
        let self = false
        if (msg.user.id == this.userModel.getUser().id){
            self = true
        }
        this.view.displayMessage(msg, self, recent)
    }

    handleClose = () => {
        this.view.close()
        this.closeSocket()
        this.newPost(this.userModel)
    }

    handleSendMessage = (msg) => {
        this.model.sendMessage(msg)
    }

    closeSocket = () => {
        this.model.closeSocket()
    }
}