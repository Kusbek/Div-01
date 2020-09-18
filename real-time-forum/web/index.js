import User from './src/user/model.js'
import UserView from './src/user/view.js'
import UserController from './src/user/controller.js'

import Post from './src/post/model.js'
import PostView from './src/post/view.js'
import PostController from './src/post/controller.js'

import Chat from './src/chat/model.js'
import ChatView from './src/chat/view.js'
import ChatController from './src/chat/controller.js'

class Main {
    constructor() {
        const userModel = new User()
        const userView = new UserView()
        const userController = new UserController(userModel, userView)


       
        this.newPost(userModel)
        this.newChat()
    }

    newPost = (user) => {
        const post = new Post()
        const postView = new PostView()
        const postController = new PostController(user, post, postView)
        //  window.post = post
        this.newDependency(post,postView,postController)
    }

    newChat = () => {
        const model = new Chat()
        const view = new ChatView()
        const controller = new ChatController(model, view)
        this.chat = this.newDependency(model, view, controller)
    }

    newDependency = (model, view, controller) => {
        return {
            model: model,
            view: view,
            controller: controller,
        }
    }
}



const main = new Main()


// window.user = user
// window.controller = controller



