document.addEventListener('htmx:afterRequest', function (evt) {
    const form = evt.target;
    if (form.tagName === "FORM") {
        form.reset();
    }
})



const allBtn = document.querySelectorAll('#task-filter button')[0]
const completedBtn = document.querySelectorAll('#task-filter button')[1]
const pendingBtn = document.querySelectorAll('#task-filter button')[2]

allBtn.addEventListener('click', e => {
    completedBtn.classList.remove('active')
    pendingBtn.classList.remove('active')
    allBtn.classList.add('active')
})

completedBtn.addEventListener('click', e => {
    pendingBtn.classList.remove('active')
    allBtn.classList.remove('active')
    completedBtn.classList.add('active')
})

pendingBtn.addEventListener('click', e => {
    allBtn.classList.remove('active')
    completedBtn.classList.remove('active')
    pendingBtn.classList.add('active')
})