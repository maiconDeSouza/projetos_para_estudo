const ul = document.querySelector('.list-tasks')

ul.addEventListener('click', e => {
    const btnDel = e.target.closest('.del')
    if (btnDel) {
        const li = btnDel.closest('li')
        const pk = li.dataset.id

        fetch(`/delete/${pk}`)
        li.remove()
        return
    }

    const btnEdit = e.target.closest('.edit')
    if (btnEdit) {
        const li = btnEdit.closest('li')
        const pk = li.dataset.id

        console.log('EDIT TASK:', pk)
        return
    }   
})
