export class Game {
    constructor(allQuestions = [], maxAnswers = 10, maxJumps = 3, jumpsUser = 0, pointsTotal = 0){
        this.allQuestions = allQuestions
        this.maxAnswers = maxAnswers,
        this.maxJumps = maxJumps,
        this.jumpsUser = jumpsUser,
        this.pointsTotal = pointsTotal
    }
}