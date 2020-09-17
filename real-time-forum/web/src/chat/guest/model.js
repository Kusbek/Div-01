export default class Guest {
    constructor({id, nickname}) {
        this.id = id
        this.nickname = nickname
    }

    get = () => {
        return {
            id: this.id,
            nickname: this.nickname
        }
    }
}