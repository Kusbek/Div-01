import Game from './src/game.js'
import View from './src/view.js'
import Controller from './src/controller.js'
const root = document.querySelector('#root')

const game = new Game()
const view = new View(root)
const controller = new Controller(game, view)

window.game = game
window.view = view

// view.renderMainScreen(game.getState())
// view.renderPauseScreen()
// view.renderStartScreen()
// view.renderEndScreen({score: 1200})

