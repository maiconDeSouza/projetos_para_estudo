let liPlay = document.querySelectorAll('.music ul li')

document.addEventListener('DOMContentLoaded', async (e) => {
    const cityName = 'São Paulo'
    const result = await searchCity(cityName)
    renderCity(result)
})

buttonSearch.addEventListener('click', e => {
    areaInputToggle()
    inputClean()
    
})

document.addEventListener('keyup',async (e) => {
    if(e.code === 'Enter') {
        const cityName = valueInput()
        areaInputToggle()
        inputClean()
        const result = await searchCity(cityName)
        renderCity(result)

        const threeMusic = await getPlaylist()
        renderMusic(threeMusic)
        liPlay = document.querySelectorAll('.music ul li')
        console.log(liPlay)
    }
})

liPlay.forEach(li => {
    
    li.addEventListener('click', e => {
        console.log('oi')
        const videoId = li.dataset.id
        playplay(videoId)
        
    })
})
