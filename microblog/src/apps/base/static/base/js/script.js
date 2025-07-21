const buttonSearch = document.querySelector('#search')
const dialog = document.querySelector('dialog')
const buttonClose = document.querySelector('#close')

buttonSearch.addEventListener('click', e => {
    dialog.showModal()
})

buttonClose.addEventListener('click', e => {
    dialog.close()
})