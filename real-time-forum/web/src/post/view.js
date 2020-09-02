export default class PostView {
    constructor(){
        this.feedWall= document.getElementById("feed-wall-wrapper")
    }

    displayPosts(posts){
        posts.forEach(post => {
            const feedCard = this.createElement("div", "feed-card")
            const title = this.createElement("h2")
            title.textContent = post.title
            const author = this.createElement("h5")
            author.textContent = post.author.nickname
            const image = this.createElement('div', 'fakeimg')
            image.textContent = "Fake Image"
            const text = this.createElement("p")
            text.textContent = post.text
            const commentsWrapper = this.createElement("div", "comments-wrapper")
            const comments = this.createElement("div", "comments")
            comments.textContent = `${post.comments} comments`
            commentsWrapper.append(comments)
            feedCard.append(title)
            feedCard.append(author)
            feedCard.append(text)
            feedCard.append(image)
            feedCard.append(commentsWrapper)
            this.feedWall.append(feedCard)
        });
    }

    createElement(tag, className) {
        const element = document.createElement(tag)
        if (className) {element.classList.add(className)}
        return element
    }
}