export default class ChatView {
    constructor(){
        this.chatWallWrapper = document.getElementById('chat-wall-wrapper')
        const chatWall = this.createElement('div', 'chat-wall')
        this.chatWall = chatWall
        this.chatWallWrapper.append(this.chatWall)
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