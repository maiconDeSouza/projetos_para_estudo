import { questions } from './questions.js'

export function setName(name){
    if(!name || name.length < 3 || name.length > 10){
        alert(`O nome precisa ter mais de 3 caracteres e menos que 11`)
        location.reload()
        return
    }
    return name
}

export function initQuestions(){
    const ul = document.querySelector('.area-card-questions__list')
    const templateAreaCardQuestion = document.querySelector('.area-card-questions__list template')
    const frag = document.createDocumentFragment()
    const idQuestionsRandom = new Set()

    while(questions.length !== idQuestionsRandom.size){
        const index = Math.floor(Math.random() * questions.length)

        idQuestionsRandom.add(questions[index].id)
    }

    Array.from(idQuestionsRandom).forEach((item, index) => {
        const li = templateAreaCardQuestion.content.querySelector('.area-card-questions__item').cloneNode(true)
        const span = li.querySelector('article span')
        li.dataset.id = item
        span.textContent = `Pergunta ${index + 1}`
        frag.appendChild(li)
    })

    ul.appendChild(frag)
}

export function selectionQuest(id){
    const li = document.querySelectorAll('.area-card-questions__item')
    if(!li)return

    const liFilter = Array.from(li).filter(item => !item.classList.contains('remove'))

    liFilter.forEach(item => {
        item.classList.remove('active')
        item.classList.add('disabled')
        if(item.dataset.id === id){
            item.classList.remove('disabled')
            item.classList.add('active')
        }

    })
}

export function renderQuestion(id){
    const main = document.querySelector('main')
    const p = document.querySelector('.main__text-question')
    const li = document.querySelectorAll('.main__item_answer')
    main.dataset.id = id
    const question = questions.find(item => item.id === id)

    p.textContent = `${question.question}?`

    li.forEach((item, index) => {
        const span = item.querySelector('article > span')
        span.textContent = question.options[index]
    })
    
}

export function setResponse(letter){
    const li = document.querySelectorAll('.main__item_answer')
    const buttonsActions = document.querySelectorAll('main > footer > button')

    li.forEach(item => {
        item.closest('.main__item_answer').classList.remove('active')
        if(item.querySelector('article > strong').textContent === letter){
            item.closest('.main__item_answer').classList.add('active')
        }
    })

    buttonsActions.forEach(item => {
        item.classList.toggle('disabled')
    })
}