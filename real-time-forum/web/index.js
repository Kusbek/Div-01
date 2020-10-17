// import User from './src/user/model.js'
// import UserView from './src/user/view.js'
// import UserController from './src/user/controller.js'

import Post from './src/post/model.js'
import PostView from './src/post/view.js'
import PostController from './src/post/controller.js'



// class Main {
//     constructor() {
//         const userModel = new User()
//         this.newPost(userModel)
//         const userView = new UserView()
//         const userController = new UserController(this.newPost, userModel, userView)

//     }



//     newPost = (user) => {
//         const post = new Post()
//         const postView = new PostView()
//         const postController = new PostController(user, post, postView)
//         this.post = this.newDependency(post,postView,postController)
//     }



//     newDependency = (model, view, controller) => {
//         return {
//             model: model,
//             view: view,
//             controller: controller,
//         }
//     }
// }



// const main = new Main()
// window.user = user
// window.controller = controller

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

            let posts = []
            json.posts.forEach((p) => {
                const post = new PostController(new Post(p), new PostView())
                posts.push(post)
            })
            return posts
        }).catch((e) => {
            displayModal(e)
        })
        return posts
    }

}

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


class MainController {
    constructor(model, view) {
        this.model = model
        this.view = view
        // this.auth()

        this.posts = []
        this.getPosts()
    }

    auth = () => {
        this.model.authorize(this.userHandler)
    }

    userHandler = (user) => {
        console.log(user)
    }

    getPosts = () => {
        this.view.createPostsContainer()
        this.model.getPosts("all").then((postControllers) => {
            this.postControllers = postControllers
        })
    }

}
const displayModal = (msg) => {
    alert(msg)
}


const createElement = (tag, ...classNames) => {
    const element = document.createElement(tag)
    if (classNames.length != 0) {
        classNames.forEach((className) => {
            element.classList.add(className)
        })
    }
    return element
}


const mainModel = new MainModel()
const mainView = new MainView()
const mainController = new MainController(mainModel, mainView)

