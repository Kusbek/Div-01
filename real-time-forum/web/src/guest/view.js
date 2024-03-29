import { createElement } from "../utils/utils.js"

export default class GuestView {
    // constructor(parentElement) {
    //     this.parent = parentElement
        
    // }

    // delete() {
    //     this.guestWrapper.remove()
    // }

    // updateState = ({action}) => {
    //     if (action == "online") {
    //         this.guestWrapper.classList.add("guest-online")
    //     } else {
    //         this.guestWrapper.classList.remove("guest-online");
    //     }
    // }

    // display = (guest) => {
    //     const guestWrapper = this.createElement('div', 'guest-wrapper')
    //     if (guest.action == "online") {
    //         guestWrapper.classList.add("guest-online")
    //     }
    //     const nickname = this.createElement('div', 'guest-nickname')
    //     nickname.textContent = guest.nickname
    //     guestWrapper.append(nickname)
    //     guestWrapper.setAttribute("guest-id", guest.id)
    //     nickname.setAttribute("guest-id", guest.id)
    //     this.guestWrapper = guestWrapper
    //     this.parent.append(guestWrapper)
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

    // bindHandleClick = (handler) => {
    //     this.guestWrapper.addEventListener('click', (event) => {
    //         handler()
    //     })
    // }

    // toTop = () => {
    //     this.parent.prepend(this.guestWrapper)
    // }

    constructor() {
         this.online = document.getElementById("online-users")   
         this.offline = document.getElementById("offline-users")   
    }

    display = ({status, user, newMessage}, handler) => {
        const guest = createElement("div", "chat-user")
        guest.textContent = user.nickname
        guest.addEventListener("click", (event)=> {
            handler()
        })
        if (status == "online") {
            this.online.append(guest)
        }else{
            this.offline.append(guest)
        }
        this.guest = guest
        if (newMessage) {
            this.makeGreen()
        }
    }

    makeGreen = () => {
        this.guest.classList.add("new-message")
    }

    makeGray = () => {
        this.guest.classList.remove("new-message") 
    }
}