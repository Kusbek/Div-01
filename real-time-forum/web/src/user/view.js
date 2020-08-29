export default class View {
    constructor() {
        this.body = document.querySelector("body");
        const authButton = document.getElementById("auth-button")
        this.authButton = authButton
    }

    updateAuthButton({isLoggedIn}) {
        if (isLoggedIn) {
            this.authButton.textContent = "Logout"
        } else {
            this.authButton.textContent = "Login"
        }
    }
}