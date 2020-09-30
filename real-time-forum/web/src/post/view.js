export default class PostView {
    constructor() {
        this.feedWall = document.getElementById("feed-wall-wrapper")
        this.displayCreatePostButton()
    }

    displayCreatePostButton() {
        let btnWrapper = this.createElement("div", "create-post-button-wrapper")
        let btn = this.createElement("button", "btn-outline-dark", "btn")
        btn.id = "create-post-button"
        btn.textContent = "Create Post"
        btnWrapper.append(btn)
        this.btnWrapper = btnWrapper
        this.feedWall.append(btnWrapper)
    }

    displayLoginWarning() {
        const logginError = this.createElement("p", "error")
        logginError.textContent = "You have to login for this action!!!"
        this.btnWrapper.append(logginError)
        setTimeout(() => {
            let ps = this.btnWrapper.querySelector("p")
            ps.remove()
        }, 1000)
    }

    displayPostTemplate() {
        const feedCard = this.createElement("div", "feed-card")
        const titleWrapper = this.createElement('div', 'post-template-wrapper')
        const title = this.createElement("input", "post-template-input")
        title.setAttribute("placeholder", "Post title")

        const textWrapper = this.createElement('div', 'post-template-wrapper')
        const text = this.createElement("textarea", "post-template-textarea")
        text.setAttribute("placeholder", "Post text")

        const saverWrapper = this.createElement('div', 'post-template-wrapper')
        const saveButton = this.createElement("button", "post-template-save-button", "btn-outline-dark", "btn")
        saveButton.textContent = "Save"
        this.savePostButton = saveButton
        titleWrapper.append(title)
        textWrapper.append(text)
        saverWrapper.append(saveButton)
        feedCard.append(titleWrapper, textWrapper, saverWrapper)
        this.feedWall.insertBefore(feedCard, this.feedWall.children[1])
    }


    closePostTemplate() {
        this.feedWall.children[1].remove()
    }

    displayPost(post) {
        const feedCard = this.createElement("div", "feed-card")
        feedCard.id = `post-${post.id}`
        const title = this.createElement("h2")
        title.textContent = post.title
        const author = this.createElement("h5")
        author.textContent = post.author.nickname
        const image = this.createElement('div', 'fakeimg')
        image.textContent = "Fake Image"
        const text = this.createElement("p")
        text.textContent = post.text
        const commentsWrapper = this.createElement("div", "comments-wrapper")
        const commentsCount = this.createElement("div", "comments-count")
        commentsCount.textContent = `${post.comments} comments`
        commentsCount.setAttribute('post-id', post.id)

        //NewComment
        const newCommentWrapper = this.createElement('div', "new-comment-wrapper")
        const newCommentInputWrapper = this.createElement('div')
        const newCommentInput = this.createElement('input')
        newCommentInput.setAttribute('placeholder', "Write your comment...")
        newCommentInputWrapper.append(newCommentInput)
        const newCommentButton = this.createElement('button', 'new-comment-button', "btn-outline-dark", "btn")
        newCommentButton.setAttribute('post-id', post.id)
        newCommentButton.textContent = "Comment"
        newCommentWrapper.append(newCommentButton)
        newCommentWrapper.append(newCommentInputWrapper)

        const detCommentsWrapper = this.createElement("div", "det-comments-wrapper")

        commentsWrapper.append(commentsCount)
        feedCard.append(title, author, text, image, commentsWrapper, newCommentWrapper, detCommentsWrapper)
        this.feedWall.insertBefore(feedCard, this.feedWall.children[1])
    }


    displayPosts(posts) {
        posts.forEach(post => {
            this.displayPost(post)
        });
    }

    incrementCommentCount(nComments) {
        let commentCount = this.tempFeedCard.querySelector("div.comments-wrapper div")
        commentCount.textContent = `${nComments} comments`
    }

    displayComment(comment) {
        const commentsWrapper = this.tempFeedCard.lastChild
        const commentWrapper = this.createElement("div", "comment-wrapper")
        const commentAuthor = this.createElement("div", "comment-author")
        commentAuthor.textContent = comment.author.nickname
        const commentText = this.createElement("div", "comment-text")
        commentText.textContent = comment.text
        commentWrapper.append(commentAuthor, commentText)
        commentsWrapper.prepend(commentWrapper)
    }

    unfoldComments(comments) {
        comments.forEach((comment) => {
            this.displayComment(comment)
        })

    }

    foldComments() {
        const commentsWrapper = this.tempFeedCard.lastChild
        commentsWrapper.innerHTML = ""
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



    bindUnfoldComments(handler) {
        const comments = document.getElementsByClassName("comments-count")
        Array.from(comments).forEach((comment) => {
            comment.addEventListener('click', (event) => {
                this.tempFeedCard = event.target.parentElement.parentElement
                handler(parseInt(event.target.getAttribute("post-id")))
            })
        })
    }

    bindNewCommentButton(handler) {
        const comments = document.getElementsByClassName("new-comment-button")

        Array.from(comments).forEach((comment) => {
            comment.addEventListener('click', (event) => {
                this.tempFeedCard = event.target.parentElement.parentElement
                const commentBody = {
                    postId: parseInt(event.target.getAttribute("post-id")),
                    text: event.target.parentElement.querySelector("input").value,
                }
                handler(commentBody)
            })
        })
    }


    bindCreatePostButton(handler) {
        const btn = document.getElementById("create-post-button")
        btn.addEventListener('click', (event) => {
            this.tempFeedCard = event.target.parentElement.parentElement
            handler()
        })
    }


    bindSavePostButton(handler) {
        const btn = this.savePostButton
        btn.addEventListener('click', (event) => {
            // this.tempFeedCard = event.target.parentElement.parentElement
            const feedCard = event.target.parentElement.parentElement
            const title = feedCard.querySelector(".post-template-input").value
            const text = feedCard.querySelector(".post-template-textarea").value
            const body = {
                title: title,
                text: text,
            }
            handler(body)
        })
    }
}