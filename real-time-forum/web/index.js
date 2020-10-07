import User from './src/user/model.js'
import UserView from './src/user/view.js'
import UserController from './src/user/controller.js'

import Post from './src/post/model.js'
import PostView from './src/post/view.js'
import PostController from './src/post/controller.js'



class Main {
    constructor() {
        const userModel = new User()
        this.newPost(userModel)
        const userView = new UserView()
        const userController = new UserController(this.newPost, userModel, userView)
        
    }



    newPost = (user) => {
        const post = new Post()
        const postView = new PostView()
        const postController = new PostController(user, post, postView)
        this.post = this.newDependency(post,postView,postController)
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



