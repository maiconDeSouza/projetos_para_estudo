import { PLAYLIST_ID, API_KEY_GOOGLE } from './env.js'


export async function getPlaylist(){
    const url = `https://www.googleapis.com/youtube/v3/playlistItems?part=snippet&maxResults=50&playlistId=${PLAYLIST_ID}&key=${API_KEY_GOOGLE}`
    try {
        const response = await fetch(url)

        if(response.status !== 200) throw Error 

        const data = await response.json()

        if(!data.items) throw Error

        const musics = data.items.sort(() => Math.random() - 0.5)
        const threeMusics = musics.slice(0, 3)

        return threeMusics
        
    } catch (error) {
        alert('Erro ao buscar playlist')
    }
}

export async function renderPlaylist(buttonStartMusics, threeMusics) {
    buttonStartMusics.forEach((button, index) => {
        const videoId = threeMusics[index].snippet.resourceId.videoId
        const url = threeMusics[index].snippet.thumbnails.high.url
        const title = threeMusics[index].snippet.title

        const img = button.querySelector('img')
        const p = button.querySelector('p')

        button.setAttribute('data-id', videoId)
        img.setAttribute('src', url)
        img.setAttribute('alt', title)
        p.textContent = title
    })
}

export async function playVideo(videoId, iframe) {
    const origin = window.location.origin;
    const url = `https://www.youtube-nocookie.com/embed/${videoId}?autoplay=1&origin=${origin}&widget_referrer=${origin}&enablejsapi=1&rel=0&modestbranding=1`;

    iframe.setAttribute('src', url)
}