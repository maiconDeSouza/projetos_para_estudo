export class Views {
    constructor() {
        this.ulQuestionsList = document.querySelector('.area-questions__list')
        this.areaQuestionsTemplates = this.ulQuestionsList.querySelector('.area-questions__template')
        this.ulRankingList = document.querySelector('.ranking__list')
        this.areaRankingTemplates = this.ulRankingList.querySelector('.ranking__template')
    }

    renderQuestions(allQuestions = []){
        const frag = document.createDocumentFragment()
        
        allQuestions.forEach((question, index) => {
            const li = this.areaQuestionsTemplates.content.querySelector('.area-questions__item').cloneNode(true)
            li.querySelector('button').textContent = `Pergunta ${index + 1}`
            li.dataset.id = question.id
            frag.appendChild(li)
        })
        this.ulQuestionsList.appendChild(frag)
    }

    renderRanking(allRanking = []){
         const frag = document.createDocumentFragment()
         
         allRanking.forEach(user => {
            const li = this.areaRankingTemplates.content.querySelector('.ranking__item').cloneNode(true)
            li.querySelector('.ranking__name').textContent = user.name
            li.querySelector('.ranking__points').textContent = `${user.points} Pontos`
            frag.appendChild(li)
         })
         this.ulRankingList.appendChild(frag)
    }
}