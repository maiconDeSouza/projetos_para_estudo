import { nouns, adjectives, verbs } from './db/list.js'
import { getRandom, filterTheme } from './utils.js'

export class Haicais{
    constructor(){
        this.nouns = nouns 
        this.adjectives = adjectives
        this.verbs = verbs
    }

    generateHaicais(){
        const nouns = getRandom(this.nouns)

        const textVerse1 = this.verse1(nouns)
        const textVerse2 = this.verse2(nouns)
        const textVerse3 = this.verse3(nouns)

        return [textVerse1, textVerse2, textVerse3]
        
    }

    verse1(nouns){
        const result = []

        nouns.themes.forEach(theme => {
            const arr = filterTheme(this.adjectives, theme)

            result.push(...arr)
        })

        const adj = getRandom(result)

        return `${nouns.word} ${adj.word},`
    }

    verse2(nouns){
        const resultVerbs = []
        const resultAdj = []

        nouns.themes.forEach(theme => {
            const arrVerbs = filterTheme(this.verbs, theme)
            resultVerbs.push(...arrVerbs)

            const arrAdj = filterTheme(this.adjectives, theme)
            resultAdj.push(...arrAdj)
        })



        const verbs = getRandom(resultVerbs)
        const adj = getRandom(resultAdj)

        return `${verbs.word} ${adj.word},`
    }

    verse3(nouns){
        const resultNouns = []
        const resultVerbs = []

        nouns.themes.forEach(theme => {
            const arrNouns = filterTheme(this.nouns, theme)
            resultNouns.push(...arrNouns)

            const arrVerbs = filterTheme(this.verbs, theme)
            resultVerbs.push(...arrVerbs)
        })

        const nounsRandom = getRandom(resultNouns)
        const verbs = getRandom(resultVerbs)

        return `${nounsRandom.word} ${verbs.word}.`
    }

    generateObj(v1, v2, v3){
        return {
            v1,
            v2,
            v3
        }
    }
}