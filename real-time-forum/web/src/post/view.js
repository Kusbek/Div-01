// export default class PostView {
//     constructor() {
//         this.feedWall = document.getElementById("feed-wall-wrapper")
//         this.selected = "all"
//     }

//     bindSelectCategory = (handler) => {
//         let catSelect = document.getElementById("post-category")
//         catSelect.addEventListener("change", (event)=>{
//             this.selected = catSelect.value
//             handler(catSelect.value)
//         })
//     }


//     displayCreatePostButton() {
//         let btnWrapper = this.createElement("div", "create-post-button-wrapper")
//         let btn = this.createElement("button", "btn-outline-dark", "btn")
//         btn.id = "create-post-button"
//         btn.textContent = "Create Post"
//         btnWrapper.append(btn)

//         const catSelect = this.createElement('select', "post-category-select")
//         catSelect.id = "post-category"
//         catSelect.append(
//             this.createSelectOption("All", "all"),
//             this.createSelectOption("Game", "game"),
//             this.createSelectOption("Movie", "movie"),
//             this.createSelectOption("News", "news")
//         )
//         catSelect.value = this.selected
//         btnWrapper.append(btn,catSelect)
//         this.btnWrapper = btnWrapper
//         this.feedWall.append(btnWrapper)
//     }

//     displayLoginWarning() {
//         const logginError = this.createElement("p", "error")
//         logginError.textContent = "You have to login for this action!!!"
//         this.btnWrapper.append(logginError)
//         setTimeout(() => {
//             let ps = this.btnWrapper.querySelector("p")
//             ps.remove()
//         }, 1000)
//     }

//     displayPostTemplate() {
//         const feedCard = this.createElement("div", "feed-card")
//         const titleWrapper = this.createElement('div', 'post-template-wrapper')
//         const title = this.createElement("input", "post-template-input")
//         title.setAttribute("placeholder", "Post title")

//         const textWrapper = this.createElement('div', 'post-template-wrapper')
//         const text = this.createElement("textarea", "post-template-textarea")
//         text.setAttribute("placeholder", "Post text")

//         const saverWrapper = this.createElement('div', 'post-template-wrapper')
//         const saveButton = this.createElement("button", "post-template-save-button", "btn-outline-dark", "btn")
//         saveButton.textContent = "Save"
//         this.savePostButton = saveButton
//         titleWrapper.append(title)
//         textWrapper.append(text)
//         saverWrapper.append(saveButton)
//         feedCard.append(this.createCategories(), titleWrapper, textWrapper, saverWrapper)
//         this.feedWall.insertBefore(feedCard, this.feedWall.children[1])
//     }

//     createCategories() {
//         const selectWrapper = this.createElement('div', 'post-template-wrapper')
//         const catSelect = this.createElement('select', "post-template-select")
//         catSelect.append(
//             this.createSelectOption("Game", "game"),
//             this.createSelectOption("Movie", "movie"),
//             this.createSelectOption("News", "news")
//         )

//         selectWrapper.append(catSelect)
//         return selectWrapper
//     }

//     createSelectOption(textContent, value) {
//         const option = this.createElement('option')
//         option.textContent = textContent
//         option.value = value
//         return option
//     }


//     closePostTemplate() {
//         this.feedWall.children[1].remove()
//     }

//     displayPost(post) {
//         const feedCard = this.createElement("div", "feed-card")
//         feedCard.id = `post-${post.id}`
//         const title = this.createElement("h2")
//         title.textContent = post.title
//         const author = this.createElement("h5")
//         author.textContent = post.author.nickname
//         const image = this.createElement('div', 'fakeimg')
//         image.textContent = "Fake Image"
//         const text = this.createElement("p")
//         text.textContent = post.text
//         const commentsWrapper = this.createElement("div", "comments-wrapper")
//         const commentsCount = this.createElement("div", "comments-count")
//         commentsCount.textContent = `${post.comments} comments`
//         commentsCount.setAttribute('post-id', post.id)

//         //NewComment
//         const newCommentWrapper = this.createElement('div', "new-comment-wrapper")
//         const newCommentInputWrapper = this.createElement('div')
//         const newCommentInput = this.createElement('input')
//         newCommentInput.setAttribute('placeholder', "Write your comment...")
//         newCommentInputWrapper.append(newCommentInput)
//         const newCommentButton = this.createElement('button', 'new-comment-button', "btn-outline-dark", "btn")
//         newCommentButton.setAttribute('post-id', post.id)
//         newCommentButton.textContent = "Comment"
//         newCommentWrapper.append(newCommentButton)
//         newCommentWrapper.append(newCommentInputWrapper)

//         const detCommentsWrapper = this.createElement("div", "det-comments-wrapper")

//         commentsWrapper.append(commentsCount)
//         feedCard.append(title, author, text, image, commentsWrapper, newCommentWrapper, detCommentsWrapper)
//         this.feedWall.insertBefore(feedCard, this.feedWall.children[1])
//     }


//     displayPosts(posts) {
//         this.displayCreatePostButton()
//         posts.forEach(post => {
//             this.displayPost(post)
//         });
//     }

//     incrementCommentCount(nComments) {
//         let commentCount = this.tempFeedCard.querySelector("div.comments-wrapper div")
//         commentCount.textContent = `${nComments} comments`
//     }

//     displayComment(comment) {
//         const commentsWrapper = this.tempFeedCard.lastChild
//         const commentWrapper = this.createElement("div", "comment-wrapper")
//         const commentAuthor = this.createElement("div", "comment-author")
//         commentAuthor.textContent = comment.author.nickname
//         const commentText = this.createElement("div", "comment-text")
//         commentText.textContent = comment.text
//         commentWrapper.append(commentAuthor, commentText)
//         commentsWrapper.prepend(commentWrapper)
//     }

//     unfoldComments(comments) {
//         comments.forEach((comment) => {
//             this.displayComment(comment)
//         })

//     }

//     foldComments() {
//         const commentsWrapper = this.tempFeedCard.lastChild
//         commentsWrapper.innerHTML = ""
//     }
//     createElement(tag, ...classNames) {
//         const element = document.createElement(tag)
//         if (classNames.length != 0) {
//             classNames.forEach((className) => {
//                 element.classList.add(className)
//             })
//         }
//         return element
//     }



//     bindUnfoldComments(handler) {
//         const comments = document.getElementsByClassName("comments-count")
//         Array.from(comments).forEach((comment) => {
//             comment.addEventListener('click', (event) => {
//                 this.tempFeedCard = event.target.parentElement.parentElement
//                 handler(parseInt(event.target.getAttribute("post-id")))
//             })
//         })
//     }

//     bindNewCommentButton(handler) {
//         const comments = document.getElementsByClassName("new-comment-button")

//         Array.from(comments).forEach((comment) => {
//             comment.addEventListener('click', (event) => {
//                 this.tempFeedCard = event.target.parentElement.parentElement
//                 const commentBody = {
//                     postId: parseInt(event.target.getAttribute("post-id")),
//                     text: event.target.parentElement.querySelector("input").value,
//                 }
//                 handler(commentBody)
//             })
//         })
//     }


//     bindCreatePostButton(handler) {
//         const btn = document.getElementById("create-post-button")
//         btn.addEventListener('click', (event) => {
//             this.tempFeedCard = event.target.parentElement.parentElement
//             handler()
//         })
//     }


//     bindSavePostButton(handler) {
//         const btn = this.savePostButton
//         btn.addEventListener('click', (event) => {
//             // this.tempFeedCard = event.target.parentElement.parentElement
//             const feedCard = event.target.parentElement.parentElement
//             const title = feedCard.querySelector(".post-template-input").value
//             const text = feedCard.querySelector(".post-template-textarea").value
//             const category = feedCard.querySelector(".post-template-select").value
//             const body = {
//                 title: title,
//                 text: text,
//                 category: category,
//             }
//             handler(body)
//         })
//     }
// }


import { createElement } from '../utils/utils.js'


//https://codepen.io/alvaromontoro/pen/ebPEWb
export default class PostView {
    constructor() {
        this.postsContainer = document.getElementById("posts-container")
    }

    displayPost = (post, newCommentHandler) => {
        const postWrapper = createElement("div", "post-wrapper")
        postWrapper.append(
            this.displayPostContainer(post,newCommentHandler),
            this.addCommentSection()
        )
        this.postsContainer.append(postWrapper)
    }

    displayPostContainer = ({ title, text, author, nComments }, newCommentHandler) => {
        const container = createElement("div", "post-container")
        this.postContainer = container
        const content = createElement("div", "post-content")
        content.append(
            this.displayTitle(title),
            this.displayText(text),
            this.displayNewComment(newCommentHandler),
            this.displayCommentsCount(nComments)
        )

        const additional = createElement("div", "additional")
        additional.append(
            this.displayAuthorCard(author),
            this.displayAuthorDetailed(author)
        )

        container.append(additional, content)
        return container
    }

    displayNewComment = (newCommentHandler) => {
        const wrapper = createElement("div", "new-comment-wrapper")
        const input = createElement("input", "new-comment-input")
        input.setAttribute("type", "text")
        input.setAttribute("placeholder", "leave a comment")
        const sendButton = createElement("button", "new-comment-button")
        sendButton.textContent = "Send"

        sendButton.addEventListener("click", (event) => {
            newCommentHandler(input.value,this.commentSection.style.display)
            input.value = ""
        })
        wrapper.append(input, sendButton)
        return wrapper
    }

    displayTitle = (title) => {
        const wrapper = createElement("h1", "post-title")
        const span = createElement("span", "post-title-span")
        span.textContent = title
        wrapper.append(span)

        return wrapper
    }

    displayAuthorCard = (author) => {
        const authorCard = createElement("div", "author-card")
        const nickname = createElement("div", "author-nickname", "center")
        nickname.textContent = `${author.nickname}`
        const age = createElement("div", "author-age", "center")
        age.textContent = `AGE: 26`
        authorCard.append(nickname, age)
        return authorCard
    }

    displayAuthorDetailed = (author) => {
        console.log(author)
        //detailed
        const detailedInfo = createElement("div", "detailed-info")

        const fullname = createElement("h1")
        fullname.textContent = `${author.first_name} ${author.last_name}`

        const email = createElement("div", "info")
        const emailSpan = createElement("span")
        emailSpan.textContent = `Email: ${author.email}`
        email.append(emailSpan)

        const sex = createElement("div", "info")
        const sexSpan = createElement("span")
        sexSpan.textContent = `Sex: ${author.gender}`
        sex.append(sexSpan)

        detailedInfo.append(fullname, email, sex)
        return detailedInfo
    }

    displayText = (text) => {
        const wrapper = createElement("p", "post-text")
        const span = createElement("span", "post-text-span")
        span.textContent = text
        wrapper.append(span)
        return wrapper
    }

    displayCommentsCount = (n) => {
        const wrapper = createElement("div", "comment-count")
        const span = createElement("span", "comment-count-span")
        span.textContent = `${n} comments`
        wrapper.append(span)
        this.commentCount = wrapper
        return wrapper
    }


    bindCommentCount = (handler) => {
        this.commentCount.addEventListener("click", (event) => {
            handler(this.commentSection.style.display)
        })
    }


    addCommentSection = () => {
        const commentSection = createElement("div", "comment-section")
        commentSection.style.display = "none"
        this.commentSection = commentSection
        return commentSection
    }


    displayCommentSection = () => {
        this.commentSection.innerHTML = ""
        this.commentSection.style.display = "block"
    }

    closeCommentSection = () => {
        this.commentSection.style.display = "none"
        this.commentSection.innerHTML = ""
    }

    clearCommentSection = () => {
        this.commentSection.innerHTML = ""
    }
}