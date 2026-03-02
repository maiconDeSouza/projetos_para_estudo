export class LocalStoreRanking {
    constructor(){}

    getLocalStoreRanking(){
        const rankingArray = JSON.parse(localStorage.getItem('ranking-quiz')) ?? []
        return rankingArray
    }

    setLocalStoreRanking(rankingArray = []){
        localStorage.setItem('ranking-quiz', JSON.parse(rankingArray))
    }
}