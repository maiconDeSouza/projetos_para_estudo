export function generateFrag(){
    return document.createDocumentFragment()
}

function generateDate(){
    const date = new Date()
    const formatDate = new Intl.DateTimeFormat('pt-br', {
        dateStyle: 'short',
        timeStyle: 'short'
    })
    return formatDate.format(date)
}

export function generateTransaction(name, value){
    return {
        id: crypto.randomUUID(),
        name,
        value,
        created: generateDate()
    }
}

export function formatCoin(value){
    const formatCoin =  new Intl.NumberFormat('pt-BR', {
        style: 'currency',
        currency: 'BRL',
    })

    return formatCoin.format(value)
}

export function generateList(arrCashFlow, ul, frag){
    ul.textContent =''
    const template = document.querySelector('template')
    arrCashFlow.forEach(item => {
        const li = template.content.querySelector('.main__list_item').cloneNode(true)
        li.setAttribute('title', item.created)
        li.dataset.id = item.id
        li.querySelector('.main__list_item_name').textContent = item.name
        li.querySelector('.main__list_item_value').textContent = formatCoin(item.value)
        li.querySelector('.main__list_item_value ').classList.add(`${item.value >= 0 ? 'entries' : 'out'}`)
        frag.appendChild(li)
    })
    console.log(ul)
    console.log(frag)
    ul.appendChild(frag)
    console.log(ul.appendChild(frag))
}

export function calc(arrCashFlow){
    const entriesOutput = document.querySelector('.header__finance-flow_entries output b')
    const outOutput = document.querySelector('.header__finance-flow_out output b')
    const totalOutput = document.querySelector('.header__finance-flow_total output b')

    const entriesTotal = arrCashFlow.filter(item => item['value'] >= 0).reduce((acc, current) => acc += current['value'], 0)
    const outTotal = arrCashFlow.filter(item => item['value'] < 0).reduce((acc, current) => acc += current['value'], 0)
    const total = entriesTotal + outTotal
    
    entriesOutput.textContent = formatCoin(entriesTotal)
    outOutput.textContent = formatCoin(outTotal)
    totalOutput.textContent = formatCoin(total)
}