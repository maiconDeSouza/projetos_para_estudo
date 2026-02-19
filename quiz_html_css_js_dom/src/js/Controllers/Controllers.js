import { questions } from '../Data/questions.js'
import { Views } from '../Views/Views.js'
import { Game } from '../Game/Game.js'
import { LocalStoreRanking } from '../LocalStore/LocalStore.js'
import { shuffleArray } from '../Utils/helpers.js'

const views = new Views()
const game = new Game(questions)
const localStoreRanking = new LocalStoreRanking()

export function renderQuestions(){
    const newArrayQuestions = shuffleArray(game.allQuestions)
    views.renderQuestions(newArrayQuestions)
}

export function renderRanking(){
    const newArrayRankings = localStoreRanking.getLocalStoreRanking()
    views.renderRanking(newArrayRankings)
}