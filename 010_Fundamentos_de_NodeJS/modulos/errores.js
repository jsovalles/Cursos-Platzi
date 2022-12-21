const badfunction_1 = () => 5 + z;
try {
    badfunction_1()
} catch (error) {
    console.log('bad function ha fallado')
    console.error(error.message)
}
console.log('continuamos1...')

// Manejar errores asincronos
function badfunction() {
    setImmediate(() => {
        try {
            return 5 + z
        } catch (error) {
            console.log('bad function ha fallado')
            console.error(error.message)
            console.log('continuamos...')
        }
    });
}
badfunction();