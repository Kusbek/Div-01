import { createElement } from "../utils/utils.js"

export default class ChatView {
    // constructor(){
    //     this.chatWallWrapper = document.getElementById('chat-wall-wrapper')
    //     const chatWall = this.createElement('div', 'chat-wall')
    //     this.chatWall = chatWall
    //     this.chatWallWrapper.append(this.chatWall)
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

    // delete = ()=>{
    //     this.chatWall.remove()
    // }


    constructor () {
        this.online = document.getElementById("online-users")   
        this.offline = document.getElementById("offline-users") 
    }

    clearChat = () => {
        this.offline.innerHTML = "" 
        this.online.innerHTML = ""
    }

    addHeaders = () => {
        this.online.append(this.createHeader("online"))
        this.offline.append(this.createHeader("offline"))
    }

    createHeader = (status) => {
        const el = createElement("p", "header")
        el.textContent = status
        return el
    }
}