const results = ['3 x 0', '2 x 1', '1 x 2', '0 x 3']

function randomResult(){
    const index = Math.floor(Math.random() * results.length)
    return results[index]
}

console.log(randomResult())