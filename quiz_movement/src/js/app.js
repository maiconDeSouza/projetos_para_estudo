import { Game } from './Game.js'
import { Render } from './Render.js'
import { shuffleArray } from './utils.js'

const game = new Game()
const render = new Render()

document.addEventListener("DOMContentLoaded", e => {
    const arr = shuffleArray(game.getQuestions())
  
    const question = arr[0]
    render.renderQuestion(question)
})