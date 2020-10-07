import Chat from '../chat/model.js'
import ChatView from '../chat/view.js'
import ChatController from '../chat/controller.js'

export default class UserController {
    constructor(newPost, user, view) {
        this.user = user
        this.view = view
        this.newPost = newPost
        // this.view.updateAuthButton(this.user.getState())
        this.authenticate()
        this.view.signInUpButton.addEventListener("click", this.handleSignInUpClick.bind(this))
        this.view.signInButton.addEventListener("click", this.handleSignInClick.bind(this))
        this.view.signUpButton.addEventListener("click", this.handleSignUpClick.bind(this))
    }

    authenticate() {
        this.user.authenticate().then((u) => {
            this.view.updateFullnameToNavBar(u)
            this.view.updateSignInUpButton(u)
            this.newChat()
        }).catch((error) => {
            this.view.updateSignInUpButton(this.user.getUser())
            
            console.log(error)
        })
    }

    handleSignInUpClick() {
        const user = this.user.getUser()
        if (user.isLoggedIn) {
            this.user.signout().then((user) => {
                this.view.updateFullnameToNavBar(user)
                this.view.updateSignInUpButton(user)
                this.deleteChat()
            }).catch((error) => {
                console.log(error)
            })
        } else {
            this.view.toggleSignInUpModal()
        }
    }

    handleSignInClick() {
        const body = {
            creds: document.getElementById("sign-in-credentials").value,
            password: document.getElementById("sign-in-password").value
        }

        this.user.signin(body).then((user) => {
            this.view.updateFullnameToNavBar(user)
            this.view.updateSignInUpButton(user)
            this.view.toggleSignInUpModal()
            this.newChat()
        
        }).catch((error) => {
            alert(error)
        })
    }



    handleSignUpClick() {
        const body = {
            nickname: document.getElementById("sign-up-username").value,
            email: document.getElementById("sign-up-email").value,
            first_name: document.getElementById("sign-up-first-name").value,
            last_name: document.getElementById("sign-up-last-name").value,
            gender: document.getElementById("sign-up-gender").value,
            age: parseInt(document.getElementById("sign-up-age").value),
            password: document.getElementById("sign-up-password").value
        }
        this.user.signup(body).then((user) => {
            this.view.updateFullnameToNavBar(user)
            this.view.updateSignInUpButton(user)
            this.view.toggleSignInUpModal()
            this.newChat()
        }).catch((error) => {
            console.log(error)
        })
    }

    newChat = () => {
        const model = new Chat()
        const view = new ChatView()
        const controller = new ChatController(this.newPost,this.user, model, view)
        this.chat = this.newDependency(model, view, controller)
    }


    deleteChat = () => {
        console.log(this.chat)
        this.chat.controller.delete()
        this.chat = undefined
    }

    newDependency = (model, view, controller) => {
        return {
            model: model,
            view: view,
            controller: controller,
        }
    }
 
}