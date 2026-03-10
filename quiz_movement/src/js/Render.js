export class Render{
    constructor(){
        this.question = document.querySelector('.quetions .question')
        this.ulAnswer = document.querySelector('.list-answer')
        this.templateAnswer = this.ulAnswer.querySelector('template')
        this.asideTotalPoints = document.querySelector('.total-points')
        this.ulRanking = document.querySelector('.list-ranking')
        this.templateRanking = this.ulRanking.querySelector('template')
    }

    renderQuestion(question){
        this.question.querySelector('h2').textContent = question.question
        this.question.querySelector('img').setAttribute('src', question.url)
        this.question.dataset.id = question.id
    }

    renderAnswers(arrAnswers = [], options){
        const frag = document.createDocumentFragment()
        this.ulAnswer.textContent = ""
        
        arrAnswers.forEach((answer, index) => {
            const li = this.templateAnswer.content.querySelector('li').cloneNode(true)
            li.querySelector('b').textContent = `${options[index]})`
            li.querySelector('p').textContent = answer
            frag.appendChild(li)
        })
        this.ulAnswer.appendChild(frag)
    }

    renderNextQuestion(question, arrAnswers, options, totalQuestions, questionAnswered, hits, errors){
        const aside = this.asideTotalPoints
        aside.querySelector('.total-questions span').textContent = `(${questionAnswered}/${totalQuestions})`
        aside.querySelector('.hits span').textContent = hits
        aside.querySelector('.errors span').textContent = errors
        this.renderQuestion(question)
        this.renderAnswers(arrAnswers, options)
    }

    renderTotalPoints(nameUser, totalQuestions, questionAnswered, hits, errors){
        const aside = this.asideTotalPoints
        aside.querySelector('.user span').textContent = nameUser
        aside.querySelector('.total-questions span').textContent = `(${questionAnswered}/${totalQuestions})`
        aside.querySelector('.hits span').textContent = hits
        aside.querySelector('.errors span').textContent = errors
    }

    renderRanking(ranking = []){
        const frag = document.createDocumentFragment()

        ranking.forEach((user, index) => {
            const li = this.templateRanking.content.querySelector('li').cloneNode(true)
            li.textContent = `${index + 1}° ${user.user} - pontos ${user.hits}`
            frag.appendChild(li)
        })
        this.ulRanking.appendChild(frag)
    }

    renderGameOver(user){
        const h2 = document.createElement('h2')
        h2.classList.add('gameover')
        h2.textContent = `${user.user}, você fez ${user.hits} pontos!`
        this.question.textContent = ''
        this.question.appendChild(h2)
    }
}