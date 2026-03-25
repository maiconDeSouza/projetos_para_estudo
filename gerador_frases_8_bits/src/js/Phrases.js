import { phrases } from './phrasesList.js'

export class Phrases {
    constructor(){
        this.phraseList = [...phrases]
        this.phraseTotal = 0
    }

    randomPhrases(){
        const len = this.phraseList.length

        const index = Math.floor(Math.random() * len)

        const phrase = this.phraseList[index]

        this.phraseList.splice(index, 1)

        return phrase
    }

    phraseSun(){
        this.phraseTotal++
    }

    phraseReturn(){
        return this.phraseTotal
    }

    setList(){
        this.phraseList = [...phrases]
    }
}