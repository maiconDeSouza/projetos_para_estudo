const card = document.querySelector('.details .card-details')
const button = document.querySelector('.details .card-details .question button')

button.addEventListener('click', e => {
    card.classList.toggle('collapse')
})