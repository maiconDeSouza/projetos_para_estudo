import { searchCity, renderCity, inputClean } from './weather.js'
import { getPlaylist, renderPlaylist, playVideo } from './music.js'

const buttonSearch = document.querySelector('#header label button')
const inputSearch = document.querySelector('#search')
const buttonsStartMusic = document.querySelectorAll('#main .main__music .main__music_content ul li button')
const iframe = document.querySelector('#footer .footer__player iframe')

document.addEventListener('DOMContentLoaded', async (e) => {
    const city = await searchCity('São Paulo')
    renderCity(city)
    const threeMusics = await getPlaylist()
    renderPlaylist(buttonsStartMusic, threeMusics)
})

buttonSearch.addEventListener('click',  async (e) => {
    const cityName = inputSearch.value
    const city = await searchCity(cityName)
    inputClean()
    renderCity(city)
})

document.addEventListener('keyup', async (e) => {
    if(e.code === 'Enter') {
        const cityName = inputSearch.value
        const city = await searchCity(cityName)
        inputClean()
        renderCity(city)
        const threeMusics = await getPlaylist()
        renderPlaylist(buttonsStartMusic, threeMusics)
    }
})

buttonsStartMusic.forEach(button => {
    button.addEventListener('click', e => {
        const videoId = button.dataset.id

        playVideo(videoId, iframe)
    })
})