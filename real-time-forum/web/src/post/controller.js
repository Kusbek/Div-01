export default class PostController {
    constructor(userModel, postModel, postView) {
        this.userModel = userModel
        this.postModel = postModel
        this.postView = postView
        this.postTemplateIsDisplayed = false
        this.displayPosts()
        this.postView.bindUnfoldComments(this.handleCommentsClick)
        this.postView.bindCreatePostButton(this.handleCreatePostClick)
        this.postView.bindNewCommentButton(this.handleNewCommentClick)
    }

    displayPosts() {
        this.postModel.getPostsFromServer().then((posts) => {
            this.postView.displayPosts(posts)
        }).catch((error) => {
            console.log(error)
        })
    }

    handleCreatePostClick = () => {
        // if (this.userModel.getUser().isLoggedIn){
        if (!this.postTemplateIsDisplayed) {
            this.postView.displayPostTemplate()
            this.postView.bindSavePostButton(this.handleSavePost)
            this.postTemplateIsDisplayed = true
        } else {
            this.postView.closePostTemplate()
            this.postTemplateIsDisplayed = false
        }
        // }else{
        //     this.postView.displayLoginWarningToCreatePost()
        // }

    }

    handleSavePost = (body) => {
        const createdPost = this.postModel.createPost(body)
        this.postView.closePostTemplate()
        this.postModel.posts.unshift(createdPost)
        this.postView.displayPost(createdPost)
    }


    handleCommentsClick = (postId) => {
        const comments = this.postModel.getComments(postId)
        if (this.postView.tempFeedCard.querySelector(".det-comments-wrapper").innerHTML == "") {
            this.postView.unfoldComments(comments)
        } else {
            this.postView.foldComments()
        }
    }

    handleNewCommentClick = (body) => {
        if (body.text != undefined && body.text != "") {
            const createdComment = this.postModel.createComment(body)
            this.postView.incrementCommentCount()

            const comments = this.postModel.getComments(body.postId)
            if (this.postView.tempFeedCard.querySelector(".det-comments-wrapper").innerHTML == "") {
                this.postView.unfoldComments(comments)
            } else {
                this.postView.tempFeedCard.querySelector(".det-comments-wrapper").innerHTML = ""
                this.postView.unfoldComments(comments)
            }
        }
    }
}