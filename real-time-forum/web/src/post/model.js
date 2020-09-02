export default class Post {
    constructor() {
        this.posts = this.getPosts()
    }
    getPosts() {

        return [
            {
                id: 1,
                title: "TITLE HEADING 1",
                text: `Some text..

                Sunt in culpa qui officia deserunt mollit anim id est laborum consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.`,
                comments: 2,
                author: {
                    id: 1,
                    nickname: "kusbek"
                }
            },
            {
                id: 2,
                title: "TITLE HEADING 2",
                text: `Some text..

                Sunt in culpa qui officia deserunt mollit anim id est laborum consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.`,
                comments: 3,
                author: {
                    id: 2,
                    nickname: "postAuthorNickname"
                }
            }
        ]

    }
}