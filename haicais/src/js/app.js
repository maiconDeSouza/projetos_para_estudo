import { Haicais } from './Haicais.js'
import { Render } from './Render.js'
import { LocalStoreHaicais } from './LocalStore.js'

const haicais = new Haicais()
const render = new Render()
const local = new LocalStoreHaicais()


const buttonGenerate = document.querySelector('.button-generate')
const buttonFavorite = document.querySelector('.button-favorite')
const ul = document.querySelector('footer ul')

document.addEventListener("DOMContentLoaded", e => {
    const listHaicais = local.get()
    render.renderHaiciasFavorite(listHaicais)
})

buttonGenerate.addEventListener('click', e => {
    const arrHaicais = haicais.generateHaicais()
    render.renderHaicias(arrHaicais)
    buttonFavorite.classList.remove('close')
})

buttonFavorite.addEventListener('click', e => {
    const v1 = document.querySelector('.verse1').textContent
    const v2 = document.querySelector('.verse2').textContent
    const v3 = document.querySelector('.verse3').textContent

    local.save(v1, v2, v3)
    const listHaicais = local.get()
    render.renderHaiciasFavorite(listHaicais)
    buttonFavorite.classList.add('close')
})

ul.addEventListener('click', e => {
    if(e.target.closest('.button-trash')){
        e.target.closest('li').remove()

        const lis = document.querySelectorAll('li')
        const listHaicais = []

        lis.forEach(li => {
            const v1 = li.querySelector('.verse-favorite1').textContent
            const v2 = li.querySelector('.verse-favorite2').textContent
            const v3 = li.querySelector('.verse-favorite3').textContent
            listHaicais.push(haicais.generateObj(v1, v2, v3))
        })
        local.set(listHaicais)
    }
})