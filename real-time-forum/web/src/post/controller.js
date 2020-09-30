export default class PostController {
    constructor(userModel, postModel, postView) {
        this.userModel = userModel
        this.postModel = postModel
        this.postView = postView
        this.postTemplateIsDisplayed = false
        this.displayPosts()

        this.postView.bindCreatePostButton(this.handleCreatePostClick)

    }

    displayPosts() {
        this.postModel.getPostsFromServer().then((posts) => {
            this.postView.displayPosts(posts)
            this.postView.bindUnfoldComments(this.handleCommentsClick)
            this.postView.bindNewCommentButton(this.handleNewCommentClick)
        }).catch((error) => {
            console.log(error)
        })
    }

    handleCreatePostClick = () => {
        if (this.userModel.getUser().isLoggedIn) {
            if (!this.postTemplateIsDisplayed) {
                this.postView.displayPostTemplate()
                this.postView.bindSavePostButton(this.handleSavePost)
                this.postTemplateIsDisplayed = true
            } else {
                this.postView.closePostTemplate()
                this.postTemplateIsDisplayed = false
            }
        } else {
            this.postView.displayLoginWarning()
        }

    }

    handleSavePost = (body) => {
        this.postModel.createPostInServer(body).then((createdPost)=>{
            this.postView.closePostTemplate()
            this.postModel.posts.unshift(createdPost)
            this.postView.displayPost(createdPost)
        }).catch((error)=>{
            this.postView.displayLoginWarning()
        })
    }

    getComments = (postId) => {
        const comments = this.postModel.getCommentsFromServer(postId).then((coms) => {
            return coms
        }).catch((error) => {
            console.log(error)
        })

        return comments
    }

    handleCommentsClick = async (postId) => {
        const comments = await this.getComments(postId)
        if (comments != undefined && comments != null) {
            if (this.postView.tempFeedCard.querySelector(".det-comments-wrapper").innerHTML == "") {
                this.postView.unfoldComments(comments)
            } else {
                this.postView.foldComments()
            }
        }
    }

    handleNewCommentClick = async (body) => {
        if (this.userModel.getUser().isLoggedIn) {
            if (body.text != undefined && body.text != "") {
                this.postModel.createCommentInServer(body).then(() => {
                    const comments = this.getComments(body.postId).then((comms) => {
                        return comms
                    })
                    return comments
                }).then((comments) => {
                    this.postView.incrementCommentCount(comments.length)
                    if (this.postView.tempFeedCard.querySelector(".det-comments-wrapper").innerHTML == "") {
                        this.postView.unfoldComments(comments)
                    } else {
                        this.postView.tempFeedCard.querySelector(".det-comments-wrapper").innerHTML = ""
                        this.postView.unfoldComments(comments)
                    }
                }).catch((error) => {
                    this.postView.displayLoginWarning()
                })


            }
        } else {
            this.postView.displayLoginWarning()
        }
    }
}