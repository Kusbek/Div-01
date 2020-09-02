export default class User {
    constructor() {
        this.isLoggedIn = false
        this.authURL = "/auth"
        this.signinURL = "/signin"
        this.signoutURL = "/signout"
        this.signupURL = "/signup"
        this.resetUser()
    }

    getUser() {
        return {
            isLoggedIn: this.isLoggedIn,
            nickname: this.nickname,
            email: this.email,
            firstName: this.firstName,
            lastName: this.lastName,
            age: this.age,
        }
    }

    setUser(user) {
        this.isLoggedIn = true
        this.nickname = user.nickname
        this.email = user.email
        this.firstName = user.first_name
        this.lastName = user.last_name
        this.age = user.age
    }

    resetUser() {
        this.isLoggedIn = false
        this.nickname = ""
        this.email = ""
        this.firstName = ""
        this.lastName = ""
        this.age = 0
    }


    async authenticate() {
        const user = await fetch(this.authURL).then((response) => {
            if (!response.ok) {
                return Promise.reject(Error(response.statusText))
            }
            return response.json()
        }).then((json) => {
            if (json.error != null || json.error != undefined) {
                return Promise.reject(Error(json.error))
            }
            this.setUser(json.user)
            return this.getUser()
        }).catch((e) => {
            return Promise.reject(Error(e))
        })

        return user
    }

    async signin(body) {
        const user = fetch(this.signinURL, {
            body: JSON.stringify(body),
            method: 'POST',
        }
        ).then((response) => {
            if (!response.ok) {
                return Promise.reject(Error(response.statusText))
            }

            return response.json()
        }).then((json) => {
            if (json.error != null || json.error != undefined) {
                return Promise.reject(Error(json.error))
            }
            this.setUser(json.user)
            return this.getUser()
        }).catch((error) => {
            return Promise.reject(error)
        })
        return user
    }

    async signup(body) {
        console.log(body)
        const user = fetch(this.signupURL, {
            body: JSON.stringify(body),
            method: 'POST',
        }
        ).then((response) => {
            if (!response.ok) {
                return Promise.reject(Error(response.statusText))
            }

            return response.json()
        }).then((json) => {
            if (json.error != null || json.error != undefined) {
                return Promise.reject(Error(json.error))
            }
            this.setUser(json.user)
            return this.getUser()
        }).catch((error) => {
            return Promise.reject(error)
        })
        return user
    }


    async signout() {
        const user = fetch(this.signoutURL, {
            method: 'POST',
        }).then((response) => {
            if (!response.ok) {
                return Promise.reject(Error(response.statusText))
            }

            this.resetUser()
            return this.getUser()
        }).catch((error) => {
            return Promise.reject(error)
        })

        return user
    }
}