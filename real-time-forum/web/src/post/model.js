export default class Post {
    constructor() {
        this.posts = []
        this.postURL = "/post"
        this.commentsURL = `/comment`
        this.comments = {}
    }

    async getPostsFromServer() {
        const posts = await fetch(this.postURL, {
            method: "GET",
        }).then((response) => {
            if (!response.ok) {
                return Promise.reject(Error(response.statusText))
            }

            return response.json()
        }).then((json) => {
            if (json.error != null || json.error != undefined) {
                return Promise.reject(Error(json.error))
            }
            this.setPosts(json.posts)
            return this.getPosts()
        }).catch((e) => {
            return Promise.reject(Error(e))
        })

        return posts
    }

    getPosts() {
        return this.posts
    }

    setPosts(posts) {
        for (post of posts) {
            this.posts.push(post)
        }
    }

    async createPostInServer(body) {
        let jsonBody = {
            title: body.title,
            text : body.text,
        }
        const newPost = await fetch(`${this.postURL}`, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
              },
            body: JSON.stringify(jsonBody)
        }).then((response) => {
            if (!response.ok) {
                return Promise.reject(Error(response.statusText))
            }
            return response.json()
        }).then((json) => {
            if (json.error != null || json.error != undefined) {
                return Promise.reject(Error(json.error))
            }
            let newPost = json.post
            console.log(newPost)
            return newPost
        }).catch((e) => {
            return Promise.reject(Error(e))
        })
       
        return newPost
    }


    async getCommentsFromServer(postId) {
        const comments = await fetch(`${this.commentsURL}?post_id=${postId}`, {
            method: "GET",
        }).then((response) => {
            if (!response.ok) {
                return Promise.reject(Error(response.statusText))
            }
            return response.json()
        }).then((json) => {
            if (json.error != null || json.error != undefined) {
                return Promise.reject(Error(json.error))
            }
            this.setComments(postId, json.comments)
            return this.getComments(postId)
        }).catch((e) => {
            return Promise.reject(Error(e))
        })

        return comments
    }

    setComments(postId, comments) {
        this.comments[postId] = comments
    }

    getComments(postId) {
        return this.comments[postId]
    }

    async createCommentInServer({ postId, text }) {
        let body = {
            post_id: postId,
            text: text,
        }
        const newComment = await fetch(`${this.commentsURL}`, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
                // 'Content-Type': 'application/x-www-form-urlencoded',
              },
            body: JSON.stringify(body)
            // body:`post_id=${postId}&text=${text}`
        }).then((response) => {
            if (!response.ok) {
                return Promise.reject(Error(response.statusText))
            }
            return response.json()
        }).then((json) => {
            if (json.error != null || json.error != undefined) {
                return Promise.reject(Error(json.error))
            }
            let newComment = json.comment 
            return newComment
        }).catch((e) => {
            return Promise.reject(Error(e))
        })
       
        return newComment
    }
}