import { getLocalStore, setLocalStore } from './localstore.js'
import { generateFrag, generateTransaction, formatCoin, generateList, calc } from './utils.js'

const arrCashFlow = getLocalStore()

console.log(arrCashFlow)

const buttonEntries = document.querySelector('.header__action_button-entries')
const buttonOut = document.querySelector('.header__action_button-out')
const closeModal = document.querySelector('.dialog__close')
const formDialog = document.querySelector('.dialog__form')
const dialog = document.querySelector('dialog')
const ul = document.querySelector('.main__list')


document.addEventListener('DOMContentLoaded', e => {
    const frag = generateFrag()

    generateList(arrCashFlow, ul, frag)
    calc(arrCashFlow)
})

buttonEntries.addEventListener('click', e => {
    formDialog.dataset.transaction = 'entries'
    formDialog.dataset.mode = 'create'
    dialog.showModal()
})

buttonOut.addEventListener('click', e => {
    formDialog.dataset.transaction = 'out'
    formDialog.dataset.mode = 'create'

    dialog.showModal()
})

closeModal.addEventListener('click', e => {
    dialog.close()
})

formDialog.addEventListener('submit', e => {
    e.preventDefault()

    const mode = formDialog.dataset.mode
    const name = formDialog.querySelector('.dialog__form_title').value.trim()
    const value = formDialog.querySelector('.dialog__form_value').value.trim()
    console.log(value)

    if (!name || !value) return

    let clearValue = value.replace(/\.|R\$\s?/g, '')
    clearValue = Number(clearValue.replace(',', '.'))

    if (mode === 'create') {
        const operation = formDialog.dataset.transaction
        if (operation === 'out') clearValue = -clearValue

        const newTransaction = generateTransaction(name, clearValue)
        arrCashFlow.unshift(newTransaction)
    }

    if (mode === 'edit') {
        const id = formDialog.dataset.editID
        const index = arrCashFlow.findIndex(item => item.id === id)

        const operation = formDialog.dataset.transaction
        if (operation === 'out') clearValue = -clearValue

        arrCashFlow[index].name = name
        arrCashFlow[index].value = clearValue
    }

    const frag = generateFrag()
    generateList(arrCashFlow, ul, frag)
    calc(arrCashFlow)
    setLocalStore(arrCashFlow)

    formDialog.dataset.transaction = ''
    formDialog.reset()
    dialog.close()
})


formDialog.querySelector('.dialog__form_value').addEventListener('input', e => {
    let value = e.target.value.replace(/\D/g, '')

    value = (value / 100).toFixed(2)

    if(value === '0.00'){
        e.target.value = ""
        return
    }

    e.target.value = new Intl.NumberFormat('pt-BR', {
      style: 'currency',
      currency: 'BRL',
    }).format(value)
})

ul.addEventListener('click', e => {
    if(e.target.classList.contains('main__list_item_action_button-remove')){
        const frag = generateFrag()
        const id = e.target.parentElement.parentElement.parentElement.dataset.id
        const index = arrCashFlow.findIndex(item => item.id === id)
        arrCashFlow.splice(index, 1)
        generateList(arrCashFlow, ul, frag)
        setLocalStore(arrCashFlow)
    }

    if(e.target.classList.contains('main__list_item_action_button-edit')){
        const id = e.target.parentElement.parentElement.parentElement.dataset.id
        const index = arrCashFlow.findIndex(item => item.id === id)
        formDialog.querySelector('.dialog__form_title').value = arrCashFlow[index].name
        formDialog.querySelector('.dialog__form_value').value = formatCoin(arrCashFlow[index].value)
        formDialog.dataset.mode = 'edit'
        formDialog.dataset.editID = id
        if(arrCashFlow[index].value < 0) formDialog.dataset.transaction = 'out'
        dialog.showModal()
    }
})