import { createElement } from '../utils/utils.js'

export default class CommentView {
    constructor(parentElement) {
        this.parentElement = parentElement
    }

    displayComment = ({author, text}) => {
        const commentWrapper = createElement("div", "comment-wrapper")
        // console.log(author)
        commentWrapper.append(
            this.displayAuthor(author),
            this.displayText(text)
        )
        this.parentElement.prepend(commentWrapper)
    }

    displayText = (text) => {
        const textBody = createElement("div", "comment-body")
        const textParagraph = createElement("p")
        textParagraph.textContent = text
        textBody.append(textParagraph)
        return textBody
    }

    displayAuthor = ({nickname}) => {
        const nicknameElement = createElement("div", "comment-nickname")
        nicknameElement.textContent = nickname
        return nicknameElement
    }

}