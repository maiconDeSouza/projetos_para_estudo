const form = document.querySelector('.form-tasks')
const ul = document.querySelector('.list-tasks')

class ArrayTasks{
    constructor(){
        this.arrTasks = []
    }

    addTask(newTask){
        this.arrTasks.push(newTask)
        localStorage.setItem('arrayTasks', JSON.stringify(this.arrTasks))
    }

    deltask(taskName){
        this.arrTasks = this.arrTasks.filter(item => item.task !== taskName)
        localStorage.setItem('arrayTasks', JSON.stringify(this.arrTasks))
    }

    checkTask(taskName){
        this.arrTasks = this.arrTasks.map(item =>  {
            if(item.task === taskName){
                item.done = !item.done
                return item
            }
            return item
        })
        localStorage.setItem('arrayTasks', JSON.stringify(this.arrTasks))
    }
}

const arrayTasks = new ArrayTasks()

function createTask(task){
    return {
        task,
        done: false
    }
}

function createElement(element){
    return document.createElement(element)
}

function createItem(taskOBJ){
    const {task, done} = taskOBJ
    const li = createElement('li')

    if(done) li.classList.toggle('check')
    
    const buttonCheck = createElement('button')
    const buttonTrash = createElement('button')
    const iCheck = createElement('i')
    const iTrash = createElement('i')
    const span = createElement('span')

    buttonCheck.classList.add('button-check')
    iCheck.classList.add('bi', 'bi-check')
    buttonTrash.classList.add('button-trash')
    iTrash.classList.add('bi', 'bi-trash')
    span.textContent = task
    buttonCheck.appendChild(iCheck)
    buttonTrash.appendChild(iTrash)
    li.appendChild(buttonCheck)
    li.appendChild(span)
    li.appendChild(buttonTrash)

    return li
}

document.addEventListener('DOMContentLoaded', e => {
    const arrTasks = JSON.parse(localStorage.getItem('arrayTasks')) || []
    const fragment = document.createDocumentFragment()

    arrTasks.forEach(item => {
        const li = createItem(item)
        fragment.appendChild(li)
    })

    ul.appendChild(fragment)

    arrayTasks.arrTasks = arrTasks
})


form.addEventListener('click', e => {
    e.preventDefault()

    const taskName = form.querySelector('input').value.trim()

    form.querySelector('input').value = ''

    if(!taskName) return

    if(arrayTasks.arrTasks.some(item => item.task === taskName)){
        alert('Tarefa Já existente')
        return
    }

    const newTask = createTask(taskName)

    const li = createItem(newTask)

    ul.appendChild(li)
    arrayTasks.addTask(newTask)
})

ul.addEventListener('click', e => {
    if(e.target.classList.contains('button-trash')){
        e.target.parentElement.remove()
        const deleteTask = e.target.parentElement.querySelector('span').innerText
        arrayTasks.deltask(deleteTask)
    }

    if(e.target.classList.contains('button-check')){
        e.target.parentElement.classList.toggle('check')
        const checkTask = e.target.parentElement.querySelector('span').innerText
        arrayTasks.checkTask(checkTask)
    }
})

