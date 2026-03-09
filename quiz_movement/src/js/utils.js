export function shuffleArray(array) {
  if (!array || !Array.isArray(array)) return []

  const shuffled = [...array]
  
  for (let i = shuffled.length - 1; i > 0; i--) {
    const randomIndex = Math.floor(Math.random() * (i + 1))
    
    
    ;[shuffled[i], shuffled[randomIndex]] = [shuffled[randomIndex], shuffled[i]]
  }
  
  return shuffled
}

export function topFive(arr = []){
  const newArr = [...arr]

  return newArr.sort((a, b) => b.hits - a.hits).slice(0, 5)
}

export const time = (temp) => new Promise(resolve => setTimeout(resolve, temp))