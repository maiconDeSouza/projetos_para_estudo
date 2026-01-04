import { API_IKEY_WEATHER } from './env.js'

const weather = document.querySelector('#main .main__weather')

export async function searchCity(cityName) {
    const url = `https://api.openweathermap.org/data/2.5/weather?q=${cityName}&appid=${API_IKEY_WEATHER}&units=metric&lang=pt_br`

    try {
        if(cityName === '') throw Error

        const response = await fetch(url)

        if(response.status !== 200) throw Error

        const data = await response.json()
      
        return data
    } catch (error) {
        alert('A pesquisa por cidade deu errado!!!')
    }
}

export function renderCity(city) {
    const img = weather.querySelector('.main__weather_header img')
    const nameCity = weather.querySelector('.main__weather_content h2')
    const temp = weather.querySelector('.main__weather_content h3')
    const tempMax = weather.querySelector('.main__weather_footer p:nth-of-type(1)')
    const tempMin = weather.querySelector('.main__weather_footer p:nth-of-type(2)')

    img.setAttribute('src', `./src/images/${city.weather[0].icon}.png`)
    img.setAttribute('alt', `${city.weather[0].description}`)
    nameCity.textContent = city.name
    temp.textContent = `${Math.round(city.main.temp)}°C`
    tempMax.textContent = `Max: ${Math.round(city.main.temp_max)}°C`
    tempMin.textContent = `Min: ${Math.round(city.main.temp_min)}°C`

}

export function inputClean() {
    document.querySelector('#search').value = ''
}