import { questions } from '../db/questions.js'

export class Game{
    constructor(){
        this.questions = [...questions]
        this.totalQuestions = this.questions.length
        this.questionAnswered = 0
        this.hits = 0
        this.errors = 0
        this.user = ''
        this.options = {
            0: 'a',
            1: 'b',
            2: 'c',
            3: 'd'
        }
    }

    getOptions(){
        return this.options
    }

    addHits(){
        return ++this.hits
    }

    getHits(){
        return this.hits
    }

    addErrors(){
        return ++this.errors
    }

    getErrors(){
        return this.errors
    }

    getQuestionAnswered(){
        return this.hits + this.errors
    }

    getTotalQuestions(){
        return this.totalQuestions
    }

    askQuestions(id){
        const index = this.questions.findIndex(question => question.id === id)
        this.questions.splice(index, 1)
    }

    getQuestion(id){
        const question = questions.find(question => question.id === id)
        return question
    }

    getQuestions(){
        return this.questions
    }

    setUser(userName){
        this.user = userName
    }

    getUserPoints(){
        return {
            user: this.user,
            hits: this.hits,
            errors: this.errors
        }
    }

    check(id, answer){
        const question = this.questions.find(question => question.id === id)
        this.askQuestions(id)
        

        if(question.answer === answer){
            this.addHits()
            return true
        }
        this.addErrors()
        return false
    }

    gameOver(){
        if(this.questions.length === 0){
            return true
        }
        return false
    }
}