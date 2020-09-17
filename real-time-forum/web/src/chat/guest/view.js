export default class GuestView {
    constructor(parentElement) {
        this.parent = parentElement
    }

    display = (guest) => {
        const guestWrapper = this.createElement('div', 'guest-wrapper')
        const nickname = this.createElement('div', 'guest-nickname')
        nickname.textContent = guest.nickname
        guestWrapper.append(nickname)
        this.parent.append(guestWrapper)
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
}