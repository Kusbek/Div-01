export default class RoomView {
    constructor() {
        this.wall = document.getElementById("feed-wall-wrapper")
    }


    displayRoom = () => {
        this.wall.innerHTML = ""
        const roomWrapper = this.createElement("div", "room-wrapper")
        const messagesWrapper = this.createElement("div", "messages-wrapper")

        const newMessageWrapper = this.createElement("div", "newmessage-wrapper")
        const newMessageInput = this.createElement('textarea', 'newmessage-input')
        this.newMessageInput = newMessageInput
        newMessageInput.setAttribute("type", "text")
        newMessageInput.setAttribute("placeholder", "write your message here")
        newMessageWrapper.append(newMessageInput)

        const newMessageSendButton = this.createElement("button", "btn-outline-dark", "btn")
        this.newMessageSendButton = newMessageSendButton
        newMessageSendButton.id = "send-message-button"
        newMessageSendButton.textContent = "Send"
        newMessageWrapper.append(newMessageSendButton)
        roomWrapper.append(messagesWrapper, newMessageWrapper)
        this.messagesWrapper = messagesWrapper
        this.wall.append(roomWrapper)
    }

    displayMessage = (msg, self) => {
        let messageWrapper
        if (self) {
            messageWrapper = this.createElement("div", "message-wrapper", "darker")
        }else{
            messageWrapper = this.createElement("div", "message-wrapper")
        }

        const timestamp = this.createElement("div", "message-timestamp")
        timestamp.textContent = msg.timestamp
        const nickname = this.createElement("div", "message-nickname")

        nickname.textContent = msg.user.nickname
        const text = this.createElement("div", "message-text")
        text.textContent = msg.text
        messageWrapper.append(timestamp, nickname, text)
        this.messagesWrapper.append(messageWrapper)
    }

    createElement(tag, ...classNames) {
        const element = document.createElement(tag)
        if (classNames.length != 0) {
            classNames.forEach((className) => {
                element.classList.add(className)
            })
        }
        return element
    }

    bindSendMessageButton(handler) {
        this.newMessageSendButton.addEventListener('click', (event) => {
            handler(this.newMessageInput.value)
        })
    }
}