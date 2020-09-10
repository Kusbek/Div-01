export default class Post {
    constructor() {
        this.posts = []
        this.postURL = "/post"
        this.comments = {
            1: [
                {
                    author: "nickfury",
                    text: "Avengers Assemble",
                },
                {
                    author: "kusbek",
                    text: "Debich",
                }
            ],

            2: [
                {
                    author: "gavnojui",
                    text: "TEXT TEXT TEXT",
                },
                {
                    author: "kusbek",
                    text: "Debich",
                },
                {
                    author: "kusbek",
                    text: "Debich",
                }
            ]
        }
    }
    async getPostsFromServer() {
        const posts = await fetch (this.postURL, {
            method: "GET",
        }).then((response)=>{
            if (!response.ok) {
                return Promise.reject(Error(response.statusText))
            }

            return response.json()
        }).then((json)=>{
            if (json.error != null || json.error != undefined) {
                return Promise.reject(Error(json.error))
            }
            console.log(json.posts)
            this.setPosts(json.posts)
            return this.getPosts()
        }).catch((e) => {
            return Promise.reject(Error(e))
        })

        return posts
    }

    getPosts(){
        return this.posts
    }

    setPosts(posts) {
        for (post of posts) {
            this.posts.push(post)
        }
        
    }

    createPost(body) {
        return {
         
                id: 4,
                title: body.title,
                text: body.text,
                comments: 0,
                author: {
                    id: 1,
                    nickname: "kusbek"
                }

        }
    }


    getComments(postId) {
        return this.comments[postId]
    }
    createComment({postId, text}) {
        const newComment = {
            author: "kusbek",
            id: text.length,
            text: text,
        }
        if (this.comments[postId] == undefined) {
            this.comments[postId] = [
                newComment,
            ]
        } else {
            this.comments[postId].push(newComment)
        }
        return newComment
    }
}



        // return [
        //     {
        //         id: 1,
        //         title: "TITLE HEADING 1",
        //         text: `Some text..

        //         Sunt in culpa qui officia deserunt mollit anim id est laborum consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.`,
        //         comments: 2,
        //         author: {
        //             id: 1,
        //             nickname: "kusbek"
        //         }
        //     },
        //     {
        //         id: 2,
        //         title: "TITLE HEADING 2",
        //         text: `Some text..

        //         Sunt in culpa qui officia deserunt mollit anim id est laborum consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.`,
        //         comments: 3,
        //         author: {
        //             id: 2,
        //             nickname: "postAuthorNickname"
        //         }
        //     },
        //     {
        //         id: 3,
        //         title: "TITLE HEADING 3",
        //         text: `Some text..

        //         Sunt in culpa qui officia deserunt mollit anim id est laborum consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.`,
        //         comments: 0,
        //         author: {
        //             id: 1,
        //             nickname: "kusbek"
        //         }
        //     }
        // ]