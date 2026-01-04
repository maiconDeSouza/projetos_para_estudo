async function getPlaylist() {
    const url = `https://www.googleapis.com/youtube/v3/playlistItems?part=snippet&maxResults=50&playlistId=${playlistID}&key=${apiKeyGoogle}`
    const response = await fetch(url)
    const data = await response.json()
    
    if (data.items) {
        const mixMusic = data.items.sort(() => Math.random() - 0.5)

        const threeMusic = mixMusic.slice(0, 3)

        return threeMusic
    }
}

function renderMusic(threeMusic) {
    const ul = document.querySelector('.music ul')
    ul.textContent = ''

    threeMusic.forEach(music => {
        const li = document.createElement('li')
        const img = document.createElement('img')
        const p = document.createElement('p')
        const url = music.snippet.thumbnails.high.url
        const videoId = music.snippet.resourceId.videoId

        img.setAttribute('src', url)
        img.setAttribute('data-id', videoId)
        p.textContent = music.snippet.title

        li.appendChild(img)
        li.appendChild(p)
        ul.appendChild(li)
    })
}

function playplay(videoId) {
    const playerVideo = document.querySelector('#player-youtube')
    playerVideo.setAttribute('src', `https://www.youtube.com/embed/${videoId}?autoplay=1`)
}