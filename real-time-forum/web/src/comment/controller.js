export default class CommentController {
    constructor(commentModel, commentView) {
        this.model = commentModel
        this.view = commentView
    }

    displayComment = () => {
        this.view.displayComment(this.model.get())
    }
}