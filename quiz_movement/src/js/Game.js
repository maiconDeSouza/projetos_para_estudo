import { questions } from '../db/questions.js'

export class Game{
    constructor(){
        this.questions = questions
        this.questionAnswered = 0
        this.hits = 0
        this.erros = 0
    }

    askQuestions(id){
        const index = this.questions.findIndex(question => question.id === id)
        this.questions.splice(index, 1)
    }

    getQuestions(){
        return this.questions
    }
}