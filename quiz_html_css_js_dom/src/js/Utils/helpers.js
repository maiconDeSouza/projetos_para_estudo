export const shuffleArray = (array) => {
  const shuffled = [...array]

  for (let i = shuffled.length - 1; i > 0; i--) {
    const randomIndex = Math.floor(Math.random() * (i + 1)); // <--- Ponto e vírgula obrigatório aqui
    
    [shuffled[i], shuffled[randomIndex]] = [shuffled[randomIndex], shuffled[i]]
  }

  return shuffled
}
