import { Phrases } from './Phrases.js'

const display = document.querySelector('.display p')
const range = document.querySelector('#speed')
const output = document.querySelector('#speedValue')
const count = document.querySelector('.counter span')
const button = document.querySelector('.controls button')

const phraseList = new Phrases()
let intervalId


const updatePhrase = () => {
    const phrase = phraseList.randomPhrases()
    display.textContent = ''
    
    phrase.split('').forEach((letter, index) => {
        setTimeout(() => {
            display.textContent += letter
        }, index * 200) 
    })

    phraseList.phraseSun()
    count.textContent = phraseList.phraseReturn()
}

const startTimer = (minutes) => {
    if (intervalId) clearInterval(intervalId)
    
    const ms = 60000 * minutes;
   
    intervalId = setInterval(updatePhrase, ms)
}

document.addEventListener("DOMContentLoaded", e => {
    updatePhrase() 
    startTimer(Number(range.value))
})

range.addEventListener('input', e => {
    output.value = e.target.value
    startTimer(Number(e.target.value))
})

button.addEventListener('click', e => {
    updatePhrase() 
    startTimer(Number(range.value))
})