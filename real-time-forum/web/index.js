import User from './src/user/model.js'
import View from './src/user/view.js'
import Controller from './src/user/controller.js'

const user = new User()
const view = new View()
const controller = new Controller(user,view)

window.user = user
window.view = view
window.controller = controller