import User from './src/user/model.js'
import UserView from './src/user/view.js'
import UserController from './src/user/controller.js'

import Post from './src/post/model.js'
import PostView from './src/post/view.js'
import PostController from './src/post/controller.js'


import Chat from './src/chat/model.js'
import ChatView from './src/chat/view.js'
import ChatController from './src/chat/controller.js'
import { createElement, displayModal } from './src/utils/utils.js'


class MainView {
    constructor() {
        this.wall = document.getElementById("wall")

    }

    createPostsContainer = () => {
        const postsContainer = createElement("div")
        postsContainer.id = "posts-container"
        this.wall.append(postsContainer)
    }
}
class MainModel {
    constructor() {
        this.authURL = "/auth"
        this.postURL = "/post"
    }
    async getPosts(category) {
        if (category === undefined || category === "") {
            category = "all"
        }
        const posts = await fetch(this.postURL + `?category=${category}`, {
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
            return json.posts
        }).catch((e) => {
            displayModal(e)
        })
        return posts
    }

    async authorize() {
        const user = await fetch(this.authURL).then((response) => {
            if (!response.ok) {
                if (response.status == 401) {
                    return response.json()
                }
                return Promise.reject(Error(response.statusText))
            }
            return response.json()
        }).then((json) => {
            if (json.error != null || json.error != undefined) {
                return {}
            }
            return json.user
        }).catch((e) => {
            displayModal(e)
            return {}
        })

        return user
    }

}

class MainController {
    constructor(model, view) {
        this.model = model
        this.view = view
        // this.auth()

        this.posts = []
        this.auth()
        this.getPosts()
        this.chatController = new ChatController(new Chat(), new ChatView())
    }

    update = () => {
        this.userController.createVisitorButtons()
        this.userController.createUserInfo()
        this.updateChat()

    }

    updateChat = () => {
        let user = this.userController.model.getUser()
        if (user.isLoggedIn) {
            this.chatController.monitorGuests()
        } else {
            this.chatController.closeChat()
        }
    }

    auth = () => {
        this.model.authorize().then((user) => {
            this.userController = new UserController(new User(user), new UserView(), this.update)
            this.update()
        })
    }


    getPosts = () => {
        this.view.createPostsContainer()
        this.model.getPosts("all").then((posts) => {
            let postControllers = []
            posts.forEach((p) => {
                const post = new PostController(new Post(p), new PostView())
                posts.push(post)
            })
            this.postControllers = postControllers
        })
    }

}


const mainModel = new MainModel()
const mainView = new MainView()
const mainController = new MainController(mainModel, mainView)

