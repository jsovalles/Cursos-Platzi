/* Limpiar la consola */
console.clear()
/* Normal */
console.log("Lo que sea");
/* Informando */
console.info("Lo que sea con info");
/* Marcando un error */
console.error("Lo que sea con error");
/* Adviritendo */
console.warn("Lo que sea advirtiendo")
/* Mostrando en forma de tabla */
let tabla = [
    {
        a: 1,
        b: "Z"
    }, {
        a: 2,
        b: "X"
    }
]
console.table(tabla);

/* Agrupar logs */
console.log("Conversacion:");

console.group("Conver")
console.log("Hola");

console.group("Conver");
console.log("Blablabla");
console.groupEnd("Conver")

console.log("Adios");
console.groupEnd("conver")

console.log("Otra cosa");
/* Ejemplo de agrupar logs con funciones */
function function1() {
    console.group("Funcion 1");
    console.log("Soy la funcion 1");
    function2()
    console.log("Sigo siendo la funcion 1");
    console.groupEnd("Funcion 1")
}
function function2() {
    console.group("Funcion 2")
    console.log("Soy la funcion 2");
    console.groupEnd("Funcion 2")

}
console.log("Empezamos");
function1()

/* Contar las veces que hago un log */

console.count("Veces")
console.count("Veces")
console.count("Veces")
console.count("Veces")
console.countReset("Veces");
console.count("Veces")
console.count("Veces")
console.count("Veces")

/* Temporizador */
console.time("Timer")
console.log("--------")
console.timeEnd("Timer");