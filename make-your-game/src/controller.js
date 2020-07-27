export default class Controller {
    constructor(game, view) {
        this.game = game
        this.view = view
        this.intervalId = null
        this.isPLaying = false

        document.addEventListener('keydown', this.handleKeyDown.bind(this))
        document.addEventListener('keyup', this.handleKeyUp.bind(this))
        this.view.renderStartScreen()
    }
    update() {
        this.game.movePieceDown()
        this.updateView()
    }

    play() {
        this.isPLaying = true
        console.log("I'm Here!!!!", this.isPLaying)
        this.startTimer()
        this.updateView()
    }

    pause() {
        this.isPLaying = false
        this.stopTimer()
        this.updateView()
    }

    reset() {
        this.game.reset()
        this.play()
    }

    updateView() {
        const state = this.game.getState()
        if (state.isGameOver) {
            this.view.renderEndScreen(state)
        }
        else if (!this.isPLaying) {
            this.view.renderPauseScreen()

        } else {
            this.view.renderMainScreen(state)
        }
        
    }

    startTimer() {
        const speed = 1000 - this.game.getState().level * 100
        if (!this.intervalId) {
            this.intervalId = setInterval(() => {
                this.update()
            }, speed > 0 ? speed : 100)
        }
    }

    stopTimer() {
        if (this.intervalId) {
            clearInterval(this.intervalId)
            this.intervalId = null
        }
    }

    handleKeyUp(event) {
        switch (event.keyCode) {
            case 40: //Down
                this.startTimer()
                break
        }
    }
    handleKeyDown(event) {
        const state = this.game.getState()
        switch (event.keyCode) {
            case 82:
                if(!this.isPLaying){
                    this.reset()
                }
                break
            case 13: //ENTER
                if (state.isGameOver) {
                    this.reset()
                } else if (this.isPLaying) {
                    this.pause()
                } else {
                    this.play()
                }
                break
            case 37: //LEFT
                this.game.movePieceLeft();
                this.updateView()
                break
            case 38: //Up
                this.game.rotatePiece();
                this.updateView()
                break
            case 39: //Right
                this.game.movePieceRight();
                this.updateView()
                break
            case 40: //Down
                this.stopTimer()
                this.game.movePieceDown();
                this.updateView()
                break
        }
        console.log(this.isPLaying)
    }
}