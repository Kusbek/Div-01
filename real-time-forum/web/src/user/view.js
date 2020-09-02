export default class UserView {
    constructor() {
        this.body = document.querySelector("body");
        this.signInUpButton = document.getElementById("sign-in-up-button")
        this.signInUpModal = '#sign-in-up-modal'
        this.signInButton = document.getElementById("sign-in-button")
        this.signUpButton = document.getElementById("sign-up-button")
        this.navbarFullname = document.getElementById("navbar-fullname")
    }

    updateSignInUpButton({isLoggedIn}) {
        if (isLoggedIn) {
            this.signInUpButton.textContent = "Sign out"
        } else {
            this.signInUpButton.textContent = "Sign up/Sign in"
        }
    }

    toggleSignInUpModal(){
        $(this.signInUpModal).modal('toggle')
    }

    updateFullnameToNavBar({nickname}) {
        if (nickname == undefined || nickname == null) {
            this.navbarFullname.textContent = null
        }
        this.navbarFullname.textContent = `${nickname}`
    }

}