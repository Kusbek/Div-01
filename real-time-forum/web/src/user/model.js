export default class User {
    constructor() {
        this.isLoggedIn = false
        this.authURL = "/auth"
    }
    getState() {
        return {
            isLoggedIn: this.isLoggedIn,
        }
    }

    async authenticate() {
        const userState = await fetch(this.authURL).then((response) => {
                if (response.ok) {
                    this.isLoggedIn = true
                }
                return this.getState()
            }
        ).catch((e) => {
            console.log(e)
        })
        return userState
    }
}