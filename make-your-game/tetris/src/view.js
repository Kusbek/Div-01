export default class View {
    static colors = {
        1: 'cyan',
        2: 'blue',
        3: 'orange',
        4: 'yellow',
        5: 'green',
        6: 'purple',
        7: 'red',
    }
    constructor(element) {
        this.element = element
        this.width = 10
        this.height = 20
        this.gameZone = document.getElementById('gamezone')
        this.panel = document.getElementById("panel")
        this.menu = document.getElementById("menu")
        this.canvas = document.getElementById("canvas")
        for (let y = 0; y < this.height; y++) {
            for (let x = 0; x < this.width; x++) {
                let cell = document.createElement('div')
                cell.classList.add("cell")
                cell.setAttribute("y", y)
                cell.setAttribute("x", x)
                this.canvas.appendChild(cell)
            }
        }

        this.panelNextPiece = document.getElementById("next-piece")
        for (let y = 0; y < 4; y++) {
            for (let x = 0; x < 4; x++) {
                let cell = document.createElement('div')
                cell.classList.add("cell")
                cell.setAttribute("y", y)
                cell.setAttribute("x", x)
                this.panelNextPiece.appendChild(cell)
            }
        }        


    }



    renderMainScreen(state) {
        this.clearScreen()
        this.renderPlayfield(state)
        this.renderPanel(state)
    }


    renderPauseScreen() {
        let e = document.createElement("div")
        e.id = 'menu-content'
        e.textContent = "Press ENTER to Resume"
        this.menu.appendChild(e)
        this.menu.style.display = "block"
    }

    renderStartScreen() {
        let e = document.createElement("div")
        e.id = 'menu-content'
        e.textContent = 'Press ENTER to Start'
        this.menu.appendChild(e)
        this.menu.style.display = "block"
    }



    renderEndScreen({score}) {
        this.clearScreen()
        let e = document.createElement("div")
        e.id = 'menu-content'
        let gameover = document.createElement("p")
        gameover.appendChild(document.createTextNode("GAME OVER"))
        let sc = document.createElement("p")
        sc.appendChild(document.createTextNode(`Score ${score}`))
        let restart = document.createElement("p")
        restart.appendChild(document.createTextNode(`Press ENTER to Restart`))
        e.appendChild(gameover)
        e.appendChild(sc)
        e.appendChild(restart)
        this.menu.appendChild(e)
        this.menu.style.display = "block"
    }


    renderPlayfield({ playfield }) {
        for (let y = 0; y < playfield.length; y++) {
            const line = playfield[y]
            for (let x = 0; x < line.length; x++) {
                const block = line[x]
                if (block) {
                    this.renderBlock(
                        x,
                        y,
                        View.colors[block]
                    )
                }

            }
        }
    }

    renderPanel({ level, score, lines, nextPiece }) {
        let pScore = document.getElementById("score")
        pScore.textContent = `Score: ${score}`
        let pLevel = document.getElementById("level")
        pLevel.textContent = `Level: ${level}`
        let pLines = document.getElementById("lines")
        pLines.textContent = `Lines: ${lines}`
        let pNext = document.getElementById("next")
        pNext.textContent = `Next:`


        for (let y = 0; y < nextPiece.blocks.length; y++) {
            for (let x = 0; x < nextPiece.blocks[y].length; x++) {
                const block = nextPiece.blocks[y][x]

                if (block) {
                    let e = this.panelNextPiece.querySelector(`[x="${x}"][y="${y}"]`)
                    e.classList.add('active')
                    e.classList.add(View.colors[block])
                    // this.renderBlock(
                    //     this.panelX + (x * this.blockWidth * 0.5),
                    //     this.panelY +120 + (y * this.blockHeight * 0.5),
                    //     this.blockWidth*0.5,
                    //     this.blockHeight*0.5,
                    //     View.colors[block]
                    // )
                }
            }
        }
    }

    renderBlock(x, y, color) {
        let e = this.canvas.querySelector(`[x="${x}"][y="${y}"]`)
        e.classList.add('active')
        e.classList.add(color)
    }

    clearScreen() {
        this.menu.style.display = "none"
        let child = this.menu.lastElementChild

        while (child) {
            this.menu.removeChild(child)
            child = this.menu.lastElementChild
        }
        let cells = document.getElementsByClassName('cell')
        Array.from(cells).forEach((cell) => {
            cell.className = ''
            cell.classList.add('cell')
        })
    }

}