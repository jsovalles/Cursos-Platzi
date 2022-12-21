process.on('beforeExit', () => {
    console.log('Ey !! El proceso esta a punto de acabar')
})

process.on('exit', () => {
    console.log('El proceso acabo')
})


process.on('uncaughtException', (err) => {
    console.log(`Ocurio un error: ${err.message}`)
})

functionQueDaError();