const buttonSearch = document.querySelector('.field-search button')
const areaInput = document.querySelector('.field-search')
const city = document.querySelector('.main-content .weather')

function areaInputToggle() {
    areaInput.classList.toggle('open')
}

function inputClean() {
    areaInput.querySelector('input').value = ''
}

function valueInput() {
    const value = areaInput.querySelector('input').value
    return value
}

async function searchCity(cityName) {
    const url = `https://api.openweathermap.org/data/2.5/weather?q=${cityName}&appid=${apiKeyWeather}&units=metric&lang=pt_br`
    try {
        if(cityName === '') {
            throw Error
        }
        const data = await fetch(url)

        if(data.status !== 200){
            throw Error
        }

        const result = await data.json()
        return result
    } catch (error) {
        alert('A pesquisa por cidade deu errado!!!')
    }
}

function renderCity(result){
    console.log(result)
    const nameCity = city.querySelector('h2')
    const img = city.querySelector('img')
    const temp = city.querySelector('h3')
    const tempMax = city.querySelector('div p:nth-of-type(1)')
    const tempMin = city.querySelector('div p:nth-of-type(2)')

    img.setAttribute('src', `./src/images/${result.weather[0].icon}.png`)
    nameCity.textContent = result.name
    temp.textContent = `${Math.round(result.main.temp)}°C`
    tempMax.textContent = `Max: ${Math.round(result.main.temp_max)}°C`
    tempMin.textContent = `Min: ${Math.round(result.main.temp_min)}°C`
}

