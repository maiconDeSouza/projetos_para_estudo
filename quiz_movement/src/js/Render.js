export class Render{
    constructor(){
        this.question = document.querySelector('.quetions .question')
    }

    renderQuestion(question){
        this.question.querySelector('h2').textContent = question.question
        this.question.querySelector('img').setAttribute('src', question.url)
    }
}