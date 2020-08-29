export default class Controller {
    constructor(user, view) {
        this.user = user
        this.view = view
        // this.view.updateAuthButton(this.user.getState())
        this.authenticate()
        this.view.authButton.addEventListener("click", this.authenticate.bind(this))
    }

    authenticate() {
        this.user.authenticate().then((userState)=>{
            this.view.updateAuthButton(userState)
        })
        
    }
}