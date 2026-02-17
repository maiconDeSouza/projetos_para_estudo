import { initQuestions, selectionQuest, setName, renderQuestion, setResponse, score, checkResposte } from './utils.js'

const ul = document.querySelector('.area-card-questions__list')
const main = document.querySelector('main')
const points = main.querySelector('header h2 span')
const responseAction = document.querySelector('.main__options')
const buttonActions = document.querySelector('main > footer')

const h1 = main.querySelector('header h1')
const name = setName(prompt('Digite seu nome:'))
h1.textContent = `Olá, ${name}!`


document.addEventListener('DOMContentLoaded', e => {    
    initQuestions()
    score.getLocoalstore()
    score.renderRanking()
})

ul.addEventListener('click', e => {
    if(!e.target.closest('.area-card-questions__item')) return

    const target = e.target.closest('.area-card-questions__item')
  
    if(target.classList.contains('disabled'))return

    const id = target.dataset.id

    selectionQuest(id)
    renderQuestion(id)
})

responseAction.addEventListener('click', e => {
    const letter = e.target.textContent

    setResponse(letter)
})

buttonActions.addEventListener('click', e => {
    if(e.target.closest('.main__button-cancel')){
        setResponse('')
    }

    if(e.target.closest('.main__button-confirm')){
        const li = document.querySelector('.main__item_answer.active')
        const id = e.target.closest('main').dataset.id
        setResponse('')
        const response = li.querySelector('span').textContent

        checkResposte(id, response, name)
    }

    if(e.target.closest('.main__button-jump')){
        const id = e.target.closest('main').dataset.id

        score.jumpQuestion(id)
    }
})