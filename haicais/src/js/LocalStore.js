export class LocalStoreHaicais {
    constructor(){
        this.name = 'haicais'
        this.upDateListHaicais = []
    }

    get(){
        const listHaicais = JSON.parse(localStorage.getItem(this.name)) || []

        return listHaicais
    }

    set(listHaicais){
        localStorage.setItem(this.name, JSON.stringify(listHaicais))
    }

    save(v1, v2, v3){
        const listHaicais = this.get()
        listHaicais.push({
            v1,
            v2,
            v3
        })

        this.set(listHaicais)
    }
}