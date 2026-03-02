function nine(mult){
    const nine = 9
    let result = nine * mult
    let returnResponse = `${nine} x ${mult} = ${result}`

    while(result.toString().length >= 2){
        returnResponse += '|'
        result.toString().split('').forEach((n, index) => {
            if(result.toString().length === index + 1){
                returnResponse += ` ${n} `
            } else {
                returnResponse += ` ${n} + `
            }
            
        })

        result = result.toString().split('').map(n => Number(n)).reduce((acc, current) => acc += current, 0)

        returnResponse += `= ${result}`
    }

    return returnResponse
}

console.log(nine(1024))