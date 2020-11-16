import User from './src/user/model.js'
import UserView from './src/user/view.js'
import UserController from './src/user/controller.js'

import Post from './src/post/model.js'
import PostView from './src/post/view.js'
import PostController from './src/post/controller.js'


import Chat from './src/chat/model.js'
import ChatView from './src/chat/view.js'
import ChatController from './src/chat/controller.js'


import Room from './src/room/model.js'
import RoomView from './src/room/view.js'
import RoomController from './src/room/controller.js'

import { createElement, displayModal } from './src/utils/utils.js'

class MainView {
    constructor() {
        this.postsContainer = document.getElementById("posts-container")
        this.roomContainer = document.getElementById("room-container")

    }

    clearPostsContainer = () => {
        this.postsContainer.innerHTML = ""
    }

    clearRoomContainer = () => {
        this.roomContainer.innerHTML = ""
    }

    createPostFilters = (getPosts) => {
        const postFilter = document.getElementById("post-filter")
        const postHeader = createElement("p", "header")
        postHeader.textContent = "post categories"
        postFilter.append(postHeader)
        let categories = ["all", "game", "news", "movie"]
        for (let category of categories) {
            let button = createElement("button", "post-category")
            button.textContent = category
            button.setAttribute("value", category)

            button.addEventListener("click", (event) => {
                getPosts(event.target.value)
            })

            postFilter.append(button)
        }
    }

    displayCreatePostButton = () => {
        const postFilter = document.getElementById("post-filter")
        const button = createElement("button", "post-category")
        button.id = "create-post"
        button.textContent = "Create post"
        button.addEventListener("click", () => {
            this.showCreatePostTemplate()
        })
        this.createPostButton = button
        postFilter.append(
            button
        )
    }

    removeCreatePostButton = () => {
        if (this.createPostButton !== undefined) {
            this.createPostButton.remove()
        }
    }

    showCreatePostTemplate = () => {
        const modal = document.getElementById("create-post-modal")
        modal.style.display = "block"
        modal.querySelector(".close").addEventListener("click", () => {
            this.closeModal(modal)
        })
        this.modal = modal
    }

    closeModal(modal) {
        if (modal === undefined || modal === null) {
            modal = this.modal
        }
        modal.style.display = "none"

    }

    bindCreatePostSubmit = (handler) => {
        const modal = document.getElementById("create-post-modal")
        let submit = modal.querySelector("button")
        submit.addEventListener("click", () => {
            let info = {
                title: modal.querySelector('input[name="title"]').value,
                text: modal.querySelector('textarea[name="post-text"]').value,
                category: modal.querySelector('select[name="post-category"]').value
            }
            handler(info)
            this.closeModal(modal)
        })

    }
}

class MainModel {
    constructor() {
        this.authURL = "/auth"
        this.postURL = "/post"
    }
    async createPostInServer(body) {
        let jsonBody = {
            title: body.title,
            text: body.text,
            category: body.category,
        }
        const newPost = await fetch(`${this.postURL}`, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(jsonBody)
        }).then((response) => {
            if (!response.ok) {
                if (response.status == 401) {
                    return response.json()
                }
                if (response.status == 400) {
                    return response.json()
                }
                return Promise.reject(Error(response.statusText))
            }
            return response.json()
        }).then((json) => {
            if (json.error != null || json.error != undefined) {
                return Promise.reject(Error(json.error))
            }
            let newPost = json.post
            return newPost
        }).catch((e) => {
            displayModal(e)
        })

        return newPost
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
        this.view.createPostFilters(this.getPosts)
        this.view.bindCreatePostSubmit(this.handleCreatePost)
        this.auth()
        this.getPosts("all")
        this.chatController = new ChatController(new Chat(), new ChatView(), this.createRoom)
        this.roomController = undefined
    }

    handleCreatePost = (body) => {
        this.model.createPostInServer(body).then(() => {
            this.getPosts("all")
        })
    }

    update = () => {
        this.userController.createVisitorButtons()
        this.userController.createUserInfo()
        this.updateCreatePostButton()
        this.updateChat()
        if (this.roomController !== undefined && this.roomController !== null) {
            this.closeRoom()
        }
    }

    updateCreatePostButton = () => {
        let user = this.userController.model.getUser()
        if (user.isLoggedIn) {
            this.view.displayCreatePostButton()
        } else {
            this.view.removeCreatePostButton()
        }
    }

    updateChat = () => {
        let user = this.userController.model.getUser()
        if (user.isLoggedIn) {
            this.chatController.monitorGuests()
        } else {
            this.chatController.clearChat()
            this.chatController.closeChat()
        }
    }

    auth = () => {
        this.model.authorize().then((user) => {
            this.userController = new UserController(new User(user), new UserView(), this.update)
            this.update()
        })
    }


    getPosts = (category) => {
        this.view.clearPostsContainer()
        this.view.clearRoomContainer()
        this.model.getPosts(category).then((posts) => {
            let postControllers = []
            posts.forEach((p) => {
                const post = new PostController(new Post(p), new PostView())
                posts.push(post)
            })
            this.postControllers = postControllers
        })
    }

    closeRoom = () => {
        this.roomController.close()
        this.getPosts("all")
    }

    createRoom = (room) => {
        this.view.clearPostsContainer()
        this.view.clearRoomContainer()
        if (this.roomController !== undefined) {
            this.roomController.close()
        }
        this.roomController = new RoomController(new Room(room), new RoomView(), this.userController, this.closeRoom)
        this.roomController.displayRoom()
    }



}


const mainModel = new MainModel()
const mainView = new MainView()
const mainController = new MainController(mainModel, mainView)

