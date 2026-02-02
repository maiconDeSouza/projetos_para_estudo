export function setLocalStore(arrCashFlow = []){
    localStorage.setItem('finance-flow', JSON.stringify(arrCashFlow))
}

export function getLocalStore(){
    const arrCashFlow = JSON.parse(localStorage.getItem('finance-flow')) || []
    return arrCashFlow
}