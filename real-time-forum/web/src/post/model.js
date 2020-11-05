// export default class Post {
//     constructor() {
//         this.posts = []
//         this.postURL = "/post"
//         this.commentsURL = `/comment`
//         this.comments = {}
//     }

//     async getPostsFromServer(category) {
//         if (category == undefined || category == "") {
//             category = "all"
//         }
//         const posts = await fetch(this.postURL + `?category=${category}`, {
//             method: "GET",
//         }).then((response) => {

//             if (!response.ok) {
//                 return Promise.reject(Error(response.statusText))
//             }

//             return response.json()
//         }).then((json) => {
//             if (json.error != null || json.error != undefined) {
//                 return Promise.reject(Error(json.error))
//             }
//             this.setPosts(json.posts)
//             return this.getPosts()
//         }).catch((e) => {
//             return Promise.reject(Error(e))
//         })

//         return posts
//     }

//     getPosts() {
//         return this.posts
//     }

//     setPosts(posts) {
//         this.posts = []
//         for (let post of posts) {
//             this.posts.push(post)
//         }
//     }

//     async createPostInServer(body) {
//         let jsonBody = {
//             title: body.title,
//             text : body.text,
//             category: body.category,
//         }
//         const newPost = await fetch(`${this.postURL}`, {
//             method: "POST",
//             headers: {
//                 'Content-Type': 'application/json'
//               },
//             body: JSON.stringify(jsonBody)
//         }).then((response) => {
//             if (!response.ok) {
//                 return Promise.reject(Error(response.statusText))
//             }
//             return response.json()
//         }).then((json) => {
//             if (json.error != null || json.error != undefined) {
//                 return Promise.reject(Error(json.error))
//             }
//             let newPost = json.post
//             console.log(newPost)
//             return newPost
//         }).catch((e) => {
//             return Promise.reject(Error(e))
//         })

//         return newPost
//     }


//     async getCommentsFromServer(postId) {
//         const comments = await fetch(`${this.commentsURL}?post_id=${postId}`, {
//             method: "GET",
//         }).then((response) => {
//             if (!response.ok) {
//                 return Promise.reject(Error(response.statusText))
//             }
//             return response.json()
//         }).then((json) => {
//             if (json.error != null || json.error != undefined) {
//                 return Promise.reject(Error(json.error))
//             }
//             this.setComments(postId, json.comments)
//             return this.getComments(postId)
//         }).catch((e) => {
//             return Promise.reject(Error(e))
//         })

//         return comments
//     }

//     setComments(postId, comments) {
//         this.comments[postId] = comments
//     }

//     getComments(postId) {
//         return this.comments[postId]
//     }

//     async createCommentInServer({ postId, text }) {
//         let body = {
//             post_id: postId,
//             text: text,
//         }
//         const newComment = await fetch(`${this.commentsURL}`, {
//             method: "POST",
//             headers: {
//                 'Content-Type': 'application/json'
//                 // 'Content-Type': 'application/x-www-form-urlencoded',
//               },
//             body: JSON.stringify(body)
//             // body:`post_id=${postId}&text=${text}`
//         }).then((response) => {
//             if (!response.ok) {
//                 return Promise.reject(Error(response.statusText))
//             }
//             return response.json()
//         }).then((json) => {
//             if (json.error != null || json.error != undefined) {
//                 return Promise.reject(Error(json.error))
//             }
//             let newComment = json.comment 
//             return newComment
//         }).catch((e) => {
//             return Promise.reject(Error(e))
//         })

//         return newComment
//     }
// }

import Comment from '../comment/model.js'
import CommentView from '../comment/view.js'
import CommentController from '../comment/controller.js'
import { displayModal } from '../utils/utils.js'

export default class Post {
    constructor({ id, author, title, text, comments, category }) {
        this.id = id
        this.author = author
        this.title = title
        this.text = text
        this.nComments = comments
        this.category = category
        this.commentURL = `/comment`
    }

    get = () => {
        return {
            id: this.id,
            author: this.author,
            title: this.title,
            text: this.text,
            nComments: this.nComments,
            category: this.category,
        }
    }


    getPostComments = async (parentElement) => {
        const comments = await fetch(`${this.commentURL}?post_id=${this.id}`, {
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
            let comments = []
            json.comments.forEach(c => {
                const comment = new CommentController(new Comment(c), new CommentView(parentElement))
                comments.push(comment)
            })
            return comments
        }).catch((e) => {
            displayModal(e)
        })

        return comments

    }

    sendNewComment = async (comment) => {
        let body = {
            post_id: this.id,
            text: comment,
        }
        await fetch(`${this.commentURL}`, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
                // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: JSON.stringify(body)
            // body:`post_id=${postId}&text=${text}`
        }).then((response) => {
            if (!response.ok) {
                if (response.status === 401 || response.status === 500) {
                    return response.json()
                }
            }
            return response.json()
        }).then((json) => {
            if (json.error != null || json.error != undefined) {
                return Promise.reject(Error(json.error))
            }
        }).catch((e) => {
            displayModal(e)
        })
    }

}