
import { createElement, formatDate } from '../utils/utils.js'

export default class RoomView {
    // constructor() {
    //     this.wall = document.getElementById("feed-wall-wrapper")
    // }
    // bindCloseButton = (handler) => {
    //     this.closeButton.addEventListener('click', (event) => {
    //         console.log("CLOSE ROOM")
    //         handler()
    //     })
    // }

    // close = () => {
    //     this.wall.innerHTML = ""
    // }


    // displayRoom = () => {
    //     this.wall.innerHTML = ""
    //     const roomWrapper = this.createElement("div", "room-wrapper")
    //     const closeWrapper = this.createElement("div", "room-close-wrapper")
    //     const closeButton = this.createElement("div", "room-close-button","btn-outline-dark", "btn")
    //     this.closeButton = closeButton
    //     closeButton.textContent = "Close"
    //     closeWrapper.append(closeButton)
    //     const messagesWrapper = this.createElement("div", "messages-wrapper")
    //     this.messagesWrapper = messagesWrapper
    //     const newMessageWrapper = this.createElement("div", "newmessage-wrapper")
    //     const newMessageInput = this.createElement('textarea', 'newmessage-input')
    //     this.newMessageInput = newMessageInput
    //     newMessageInput.setAttribute("type", "text")
    //     newMessageInput.setAttribute("placeholder", "write your message here")
    //     newMessageWrapper.append(newMessageInput)

    //     const newMessageSendButton = this.createElement("button", "btn-outline-dark", "btn")
    //     this.newMessageSendButton = newMessageSendButton
    //     newMessageSendButton.id = "send-message-button"
    //     newMessageSendButton.textContent = "Send"
    //     newMessageWrapper.append(newMessageSendButton)
    //     roomWrapper.append(closeWrapper,messagesWrapper, newMessageWrapper)
    //     this.messagesWrapper = messagesWrapper
    //     this.wall.append(roomWrapper)
    // }

    // displayMessage = (msg, self, recent) => {
    //     let messageWrapper
    //     if (self) {
    //         messageWrapper = this.createElement("div", "message-wrapper", "darker")
    //     }else{
    //         messageWrapper = this.createElement("div", "message-wrapper")
    //     }

    //     const timestamp = this.createElement("div", "message-timestamp")
    //     timestamp.textContent = msg.timestamp
    //     const nickname = this.createElement("div", "message-nickname")

    //     nickname.textContent = msg.user.nickname
    //     const text = this.createElement("div", "message-text")
    //     text.textContent = msg.text
    //     messageWrapper.append(timestamp, nickname, text)
    //     if (recent) {
    //         this.messagesWrapper.append(messageWrapper)
    //     } else {
    //         this.messagesWrapper.prepend(messageWrapper)
    //     }

    // }

    // createElement(tag, ...classNames) {
    //     const element = document.createElement(tag)
    //     if (classNames.length != 0) {
    //         classNames.forEach((className) => {
    //             element.classList.add(className)
    //         })
    //     }
    //     return element
    // }

    // bindSendMessageButton(handler) {
    //     this.newMessageSendButton.addEventListener('click', (event) => {
    //         handler(this.newMessageInput.value)
    //     })
    // }

    // bindGetOldMessagesScrool(handler) {
    //     this.messagesWrapper.addEventListener('scroll', (event) => {

    //         if (this.messagesWrapper.scrollTop == 0) {
    //             handler()
    //         }
    //     })
    // }

    constructor() {
        this.roomContainer = document.getElementById("room-container")

    }


    display = ({ id }, sendMessageHandler, getOldMessages, closeRoom) => {
        const room = createElement("div", "room")
        room.id = `room-${id}`
        const close = createElement("div", "imgcontainer")
        const closeSpan = createElement("span", "close")
        closeSpan.innerHTML = `&times`
        closeSpan.addEventListener("click", (event) => {
            closeRoom()
        })
        close.append(closeSpan)

        const msgHistory = createElement("div", "msg-history")
        this.msgHistory = msgHistory
        msgHistory.addEventListener('scroll', (event) => {
            console.log(event)
            if (msgHistory.scrollTop == 0 ) {
                getOldMessages()
            }
        })
        const msgType = createElement("div", "type-msg")
        const inputMsgWrite = createElement("div", "input-msg-write")
        inputMsgWrite.append(
            this.createMsgInput(),
            this.createSendMsgButton(sendMessageHandler)
        )
        msgType.append(inputMsgWrite)
        room.append(close, msgHistory, msgType)
        this.roomContainer.append(
            room
        )
    }

    createMsgInput = () => {
        const input = createElement("input", "write-msg")
        input.setAttribute("type", "text")
        input.setAttribute("placeholder", "Type a message")
        this.input = input
        return input
    }

    createSendMsgButton = (handler) => {
        const button = createElement("button", "msg-send-btn")
        button.setAttribute("type", "button")
        button.addEventListener("click", () => {
            if (this.input.value !== "") {
                handler(this.input.value)
                this.input.value = ""
            }
        })
        const image = createElement("i", "fa", "fa-paper-plane-o")
        image.setAttribute("aria-hidden", "true")
        button.append(image)
        return button
    }

    displayMsg = ({ text, timestamp, user }, self, recent) => {
        let msgWrapperClass, msgClass
        if (self) {
            msgWrapperClass = "outgoing_msg"
            msgClass = "sent_msg"
        } else {
            msgWrapperClass = "incoming_msg"
            msgClass = "received_msg"
        }
        const msgWrapper = createElement("div", msgWrapperClass)
        const msg = createElement("div", msgClass)
        if (self) {
            msg.append(
                this.createMsgTxt(text),
                this.createAdditionalInfo(timestamp, user)
            )
        } else {
            const incTxtWrapper = createElement("div", "received_withd_msg")

            incTxtWrapper.append(
                this.createMsgTxt(text),
                this.createAdditionalInfo(timestamp, user)
            )
            msg.append(
                incTxtWrapper
            )
        }


        msgWrapper.append(
            msg
        )

        if (recent) {
            this.msgHistory.append(
                msgWrapper
            )
        } else {
            this.msgHistory.prepend(
                msgWrapper
            )
        }

    }

    createMsgTxt = (text) => {
        const txt = createElement("p")
        txt.textContent = text
        return txt
    }

    createAdditionalInfo = (timestamp, user) => {
        const time = createElement("span", "time_date")
        time.textContent = `${formatDate(timestamp)} | ${user.nickname}`
        return time
    }
}