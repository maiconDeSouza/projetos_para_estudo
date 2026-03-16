import { Game } from './Game.js'
import { Render } from './Render.js'
import { RankingLocalStore } from './LocalStore.js'
import { shuffleArray, time, topFive } from './utils.js'

const game = new Game()
const render = new Render()
const ranking = new RankingLocalStore()

const ulAnswer = document.querySelector('.list-answer')
const hitAnswer = document.getElementById('hitAnswer')
const errorAnswer = document.getElementById('errorAnswer')
const buttonStop = document.querySelector('.button-stop')


document.addEventListener("DOMContentLoaded", e => {
    const nameUser = prompt('Digite seu nome:')
    const arr = shuffleArray(game.getQuestions())

    if(!nameUser){
        location.reload()
        return
    }

    game.setUser(nameUser)
  
    const question = arr[0]
    const alternatives = shuffleArray(question.alternatives)
    const options = game.getOptions()
    const totalQuestions = game.getTotalQuestions()
    const questionAnswered = game.getQuestionAnswered()
    const hits = game.getUserPoints().hits
    const errors = game.getUserPoints().errors
    const rankingStatus = topFive(ranking.getRanking())

    render.renderTotalPoints(nameUser, totalQuestions, questionAnswered, hits, errors)
    render.renderQuestion(question)
    render.renderAnswers(alternatives, options)
    render.renderRanking(rankingStatus)
})

ulAnswer.addEventListener('click', async e => {
    const li = e.target.closest('li')

    if(!li) return

    li.classList.add('active')

    const confirmAnswer = confirm('Confirma a resposta?')
    
    if(!confirmAnswer){
        li.classList.remove('active')
        return
    }

    await time(3000)

    const answer = li.querySelector('p').textContent
    const id = document.querySelector('.question').dataset.id

    const hit = game.check(id, answer)
    

    if(!hit){
        li.classList.remove('active')
        li.classList.add('fault')
        errorAnswer.currentTime = 0
        errorAnswer.play()
    
        const questionResponse = game.getQuestion(id)
        const feedback = '❌ Errou! Tente novamente.'
        const answer = `Resposta Certa: ${questionResponse.answer}`
        render.renderPopover(feedback, questionResponse.question, answer)
        
        if(game.gameOver()){
            const user = game.getUserPoints()
            ranking.setRanking(user)
            // alert(`Você fez ${game.getHits()} pontos`)
            // render.renderRanking(ranking.getRanking())
            render.renderGameOver(user)
        }
        const arr = shuffleArray(game.getQuestions())
        const question = arr[0]
        const alternatives = shuffleArray(question.alternatives)
        const options = game.getOptions()
        const totalQuestions = game.getTotalQuestions()
        const questionAnswered = game.getQuestionAnswered()
        const hits = game.getUserPoints().hits
        const errors = game.getUserPoints().errors

        await time(3000)
        render.renderNextQuestion(question, alternatives, options, totalQuestions, questionAnswered, hits, errors)
        return
    }

    hitAnswer.currentTime = 0
    hitAnswer.play()

    const questionResponse = game.getQuestion(id)
    const feedback = '✅ Parabéns! Você acertou. 🎉👏🏻'
    render.renderPopover(feedback, questionResponse.question, questionResponse.answer)

    if(game.gameOver()){
        const user = game.getUserPoints()
        ranking.setRanking(user)
        // alert(`Você fez ${game.getHits()} pontos`)
        // render.renderRanking(ranking.getRanking())
        render.renderGameOver(user)
        return
    }

    const arr = shuffleArray(game.getQuestions())
    const question = arr[0]
    const alternatives = shuffleArray(question.alternatives)
    const options = game.getOptions()
    const totalQuestions = game.getTotalQuestions()
    const questionAnswered = game.getQuestionAnswered()
    const hits = game.getUserPoints().hits
    const errors = game.getUserPoints().errors
    await time(3000)
    render.renderNextQuestion(question, alternatives, options, totalQuestions, questionAnswered, hits, errors)
})

buttonStop.addEventListener('click', e => {
    const user = game.getUserPoints()
    ranking.setRanking(user)
    render.renderGameOver(user)
})