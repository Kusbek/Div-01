// export default class User {
//     constructor() {
//         this.isLoggedIn = false
//         this.authURL = "/auth"
//         this.signinURL = "/signin"
//         this.signoutURL = "/signout"
//         this.signupURL = "/signup"
//         this.resetUser()
//     }

//     getUser() {
//         return {
//             id: this.id,
//             isLoggedIn: this.isLoggedIn,
//             nickname: this.nickname,
//             email: this.email,
//             firstName: this.firstName,
//             lastName: this.lastName,
//             age: this.age,
//         }
//     }

//     setUser(user) {
//         this.id = user.id
//         this.isLoggedIn = true
//         this.nickname = user.nickname
//         this.email = user.email
//         this.firstName = user.first_name
//         this.lastName = user.last_name
//         this.age = user.age
//     }

//     resetUser() {
//         this.isLoggedIn = false
//         this.nickname = ""
//         this.email = ""
//         this.firstName = ""
//         this.lastName = ""
//         this.age = 0
//         this.id = undefined
//     }


//     async authenticate() {
//         const user = await fetch(this.authURL).then((response) => {
//             if (!response.ok) {
//                 return Promise.reject(Error(response.statusText))
//             }
//             return response.json()
//         }).then((json) => {
//             if (json.error != null || json.error != undefined) {
//                 return Promise.reject(Error(json.error))
//             }
//             this.setUser(json.user)
//             return this.getUser()
//         }).catch((e) => {
//             return Promise.reject(Error(e))
//         })

//         return user
//     }

//     async signin(body) {
//         const user = fetch(this.signinURL, {
//             body: JSON.stringify(body),
//             method: 'POST',
//         }).then((response) => {
//         if (!response.ok) {
//             if (response.status == 403) {
//                 return Promise.reject(Error("No such user"))
//             }
//             return Promise.reject(Error(response.statusText))
//         }
//         return response.json()
//     }).then((json) => {
//         if (json.error != null || json.error != undefined) {
//             return Promise.reject(Error(json.error))
//         }
//         this.setUser(json.user)
//         return this.getUser()
//     })
//         return user
//     }

//     async signup(body) {
//         const user = fetch(this.signupURL, {
//             body: JSON.stringify(body),
//             method: 'POST',
//         }).then((response) => {
//             if (!response.ok) {
//                 if (response.status == 409) {
//                     return Promise.reject(Error("User Already Exists"))
//                 }
//                 return Promise.reject(Error(response.statusText))
//             }
//             return response.json()
//         }).then((json) => {
//             if (json.error != null || json.error != undefined) {
//                 return Promise.reject(Error(json.error))
//             }
//             this.setUser(json.user)
//             return this.getUser()
//         })
//         return user
//     }


//     async signout() {
//         const user = fetch(this.signoutURL, {
//             method: 'POST',
//         }).then((response) => {
//             if (!response.ok) {

//                 return Promise.reject(Error(response.statusText))
//             }

//             this.resetUser()
//             return this.getUser()
//         }).catch((error) => {
//             return Promise.reject(error)
//         })

//         return user
//     }
// }

import { displayModal } from '../utils/utils.js'
export default class User {
    constructor(user) {
        this.setUser(user)
        this.signinURL = "/signin"
        this.signoutURL = "/signout"
        this.signupURL = "/signup"
    }

    resetUser() {
        this.id = undefined
        this.isLoggedIn = false
        this.nickname = undefined
        this.email = undefined
        this.firstName = undefined
        this.lastName = undefined
        this.age = undefined
        this.gender = undefined
    }

    setUser({ id, nickname, first_name, last_name, email, gender, age }) {
        this.id = id
        if (id === undefined || id === null) {
            this.isLoggedIn = false
        } else {
            this.isLoggedIn = true
        }
        this.nickname = nickname
        this.email = email
        this.firstName = first_name
        this.lastName = last_name
        this.age = age
        this.gender = gender
    }

    getUser() {
        return {
            id: this.id,
            isLoggedIn: this.isLoggedIn,
            nickname: this.nickname,
            email: this.email,
            firstName: this.firstName,
            lastName: this.lastName,
            age: this.age,
            gender: this.gender
        }
    }

    async login(credentials) {

        const user = fetch(this.signinURL, {
            body: JSON.stringify(credentials),
            method: 'POST',
        }).then((response) => {
            if (!response.ok) {
                if (response.status === 403) {
                    return response.json()
                }
                return Promise.reject(Error(response.statusText))
            }
            return response.json()
        }).then((json) => {
            if (json.error != null || json.error != undefined) {
                return Promise.reject(Error(json.error))
            }
            return json.user
        }).catch((e) => {
            displayModal(e)
            return {}
        })
        return user
    }

    async logout() {
        fetch(this.signoutURL, {
            method: 'POST',
        }).then((response) => {
            if (!response.ok) {
                return Promise.reject(Error(response.statusText))
            }
        }).catch((e) => {
            displayModal(e)
        })
    }

    async signup(userInfo) {
        const user = fetch(this.signupURL, {
            body: JSON.stringify(userInfo),
            method: 'POST',
        }).then((response) => {
            if (!response.ok) {
                if (response.status == 409) {
                    return response.json()
                }
                return Promise.reject(Error(response.statusText))
            }
            return response.json()
        }).then((json) => {
            if (json.error != null || json.error != undefined) {
                return Promise.reject(Error(json.error))
            }
            
            return json.user
        }).catch((e) => {
            displayModal(e)
            return {}
        })
        return user
    }

}