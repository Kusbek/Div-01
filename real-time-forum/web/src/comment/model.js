export default class Comment {
    constructor({author, id, post_id, text}) {
        this.id = id
        this.postId = post_id
        this.author = author
        this.text = text
    }


    get = () => {
        return {
            id: this.id,
            postId: this.postId,
            text: this.text,
            author: this.author
        }
    }

}