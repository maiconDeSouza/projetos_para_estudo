import { shuffleArray } from '../Utils/helpers.js'

export class Views {
    constructor() {
        this.sectionQuestion = document.querySelector('.main__section-question'),
        this.ulAnswersList = document.querySelector('.main__answers-list'),
        this.templateAnswerItem = document.querySelector('.template-answers-item')

        this.ulRankingList = document.querySelector('.main__ranking-list')
        this.templateRankimgItem = document.querySelector('.template-ranking-item')

        this.buttonsAction = document.querySelectorAll('.main__actions-list button')
    }
    renderNextQuestion(question, points, optionsLetter){
        const frag = document.createDocumentFragment()
        const header = this.sectionQuestion.querySelector('header')
        this.ulAnswersList.textContent = ""

        header.querySelector('h1').textContent = `Pontos: ${points}`
        header.querySelector('h2').textContent = question.question

        shuffleArray(question.options).forEach((op, index) => {
            const li = this.templateAnswerItem.content.querySelector('.main__answer-item').cloneNode(true)
            li.querySelector('b').textContent = optionsLetter[index]
            li.querySelector('p').textContent = op
            li.dataset.id = question.id
            frag.appendChild(li)
        })
        this.ulAnswersList.appendChild(frag)
    }

    renderRanking(ranking = []){
        const frag = document.createDocumentFragment()
        this.ulRankingList.textContent = ''
        
        ranking.forEach(item => {
            const li = this.templateRankimgItem.content.querySelector('.main__ranking-item').cloneNode(true)
            
            li.querySelector('.main__ranking-item-name').textContent = item.name
            li.querySelector('.main__ranking-item-points').textContent = `${item.points} Pontos`
            frag.appendChild(li)
        })
        this.ulRankingList.appendChild(frag)
    }

    selectAnswer(answer){
        const li = document.querySelectorAll('.main__answer-item')
    
        li.forEach(item => {
            if(item.querySelector('p').textContent === answer){
                item.classList.add('active')
            }
            item.classList.add('disabled')
        })

        this.buttonsActions()
    }

    buttonsActions(){
        this.buttonsAction.forEach(button => {
            button.disabled = !button.disabled
        })
    }
}