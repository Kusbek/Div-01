export default class PostController {
    constructor(postModel, postView) {
        this.postModel = postModel
        this.postView = postView
        this.postView.displayPosts(this.postModel.getPosts())
    }
}