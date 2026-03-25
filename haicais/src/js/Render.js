export class Render{
    constructor(){
        this.template = document.querySelector('template')
        this.outputHaicais = document.querySelector('.haicais')
        this.ul = document.querySelector('footer ul')
    }

    renderHaicias(arrHaicais){
        this.outputHaicais.querySelector('.verse1').textContent = arrHaicais[0]
        this.outputHaicais.querySelector('.verse2').textContent = arrHaicais[1]
        this.outputHaicais.querySelector('.verse3').textContent = arrHaicais[2]
    }

    renderHaiciasFavorite(arrHaicaisFavorite = []){
        const frag = document.createDocumentFragment()
        this.ul.textContent = ''
        
        arrHaicaisFavorite.forEach(haicai => {
            const li = this.template.content.querySelector('li').cloneNode(true)
            li.querySelector('.verse-favorite1').textContent = haicai.v1
            li.querySelector('.verse-favorite2').textContent = haicai.v2
            li.querySelector('.verse-favorite3').textContent = haicai.v3
            frag.appendChild(li)
        })

        this.ul.appendChild(frag)
    }
}