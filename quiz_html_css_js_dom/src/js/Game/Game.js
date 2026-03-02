export class Game {
    constructor(allQuestions = [], allRanking = [], points = 0, jumps = 0){
        this.allQuestions = allQuestions,
        this.optionsLetter = {
            0: 'a',
            1: 'b',
            2: 'c',
            3: 'd'
        }
        this.points = points,
        this.jumps = jumps
        this.totalMaxJumps = 3
        this.allRanking = allRanking
    }

    removeQuestions(id){
        const newArrayQuestions = this.allQuestions.filter(question => question.id !== id)
        return newArrayQuestions
    }

    sumPoints(){
        if(this.points >= 20 || this.allQuestions.length === 0){
            alert('fim do jogo')
        }
        this.points++
    }

    sumJumps(){
        if(this.jumps >= this.totalMaxJumps){
            return false
        }

        return this.jumps++
    }

    checkAnswer(id, answer){
        const question = this.allQuestions.find(question => question.id === id)
       
        if(question.answer === answer){
            this.sumPoints()
            this.removeQuestions(id)
            return true
        }

        return false
    }
}