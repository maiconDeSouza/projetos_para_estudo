import { Controller } from './Controllers/Controllers.js'
const controller = new Controller()

const ulAnswers = document.querySelector('.main__answers-list')
const ulActions = document.querySelector('.main__actions-list')

document.addEventListener("DOMContentLoaded", e => {
    controller.startApp()
})

ulAnswers.addEventListener('click', e => {
    if(e.target.closest('li')){
        // const id = e.target.closest('li').dataset.id
        const answer = e.target.closest('li').querySelector('p').textContent

        controller.selectAnswer(answer)
    }
})

ulActions.addEventListener('click', e => {
    if(e.target.closest('.main__action-item-confirme button')){
        const id = document.querySelector('.main__answer-item.active').dataset.id
        const answer = document.querySelector('.main__answer-item.active p').textContent
        controller.checkQuestion(id, answer)
    }
})