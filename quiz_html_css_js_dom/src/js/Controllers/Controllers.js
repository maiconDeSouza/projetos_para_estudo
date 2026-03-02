import { questions } from '../Data/questions.js'
import { Views } from '../Views/Views.js'
import { Game } from '../Game/Game.js'
import { LocalStoreRanking } from '../LocalStore/LocalStore.js'
import { shuffleArray, getRandomItem } from '../Utils/helpers.js'

const view = new Views()
const localStoreRanking = new LocalStoreRanking()
const game = new Game(questions, localStoreRanking)


export class Controller {
    constructor() {}

    startApp(){
        const question = getRandomItem(game.allQuestions)
        const ranking = localStoreRanking.getLocalStoreRanking()

        view.renderNextQuestion(question, game.points, game.optionsLetter)
        view.renderRanking(ranking)
    }

    selectAnswer(answer){
        view.selectAnswer(answer)
    }

    checkQuestion(id, answer){
        const result = game.checkAnswer(id, answer)

        if(result){
            alert('Você acertou!')
            this.startApp()
            view.buttonsActions()
            return
        }

        const name = prompt('Fim de jogo digite seu nome:')
        const points = game.points

        game.allRanking.push({name, points})

        localStoreRanking.setLocalStoreRanking(game.allRanking)
        location.reload()

    }
}