export default class RoomController {
    constructor(userModel,room, roomView) {
        this.userModel = userModel
        this.model = room
        this.view = roomView
        this.view.displayRoom()
        this.view.bindSendMessageButton(this.handleSendMessage)
        this.handleMessages()
    }

    handleMessages = () => {
        this.model.monitorMessages(this.newMessageHandler)
    }

    newMessageHandler = (msg) => {
        if (msg.user.id == this.userModel.getUser().id){
            this.view.displayMessage(msg, true)
        }else{
            this.view.displayMessage(msg, false)
        }
        
        
    }

    handleSendMessage = (msg) => {
        this.model.sendMessage(msg)
    }
}