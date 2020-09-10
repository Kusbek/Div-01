import User from './src/user/model.js'
import UserView from './src/user/view.js'
import UserController from './src/user/controller.js'

import Post from './src/post/model.js'
import PostView from './src/post/view.js'
import PostController from './src/post/controller.js'


const user = new User()
const userView = new UserView()
const userController = new UserController(user, userView)

const post = new Post()
const postView = new PostView()
const postController = new PostController(user, post, postView)

window.post = post
// window.view = view
// window.controller = controller