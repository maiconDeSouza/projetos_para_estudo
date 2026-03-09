export class RankingLocalStore{
    constructor(){
        this.ranking = JSON.parse(localStorage.getItem('quizMovement')) || []
    }

    getRanking(){
        return this.ranking
    }

    setRanking(user){
        this.ranking.push(user)
        localStorage.setItem('quizMovement', JSON.stringify(this.ranking))
    }
}