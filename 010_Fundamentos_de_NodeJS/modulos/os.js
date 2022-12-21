const os = require('os');

const SIZE = 1024;

console.log(os.hostname());//  Voy a saber el hostname de la máquina
console.table(os.networkInterfaces());// Puedo acceder a mi interfaz de red activas en mi máquina, puedo saber  IPVX
console.log(os.tmpdir())//-->Me muestra los directorios temporales, temproales una imagen que voy a procesar
console.log(os.homedir()) // Me permite saber cual es el directorio raíz
console.log(os.arch()); //----> Me devuelve la arquitecura de mi OS
console.log(os.platform());//---> Me dice en qué plataforma estoy
console.table(os.cpus());//--->podemos acceder a la información de las cpus de mi pc.
//console.table(os.constants);//--->  Me muestran todos los errores de sistema.


//Acceder a espacios de memoria es muy útil para saber si tengo a memoria suficiente para realizar esta operación.
console.log(os.freemem());// ---> Me dice en bytes la memoria libre que tenemos
// console.log(kb(os.freemem()));
// console.log(mb(os.freemem()));
// console.log(gb(os.freemem()));
console.log(gb(os.totalmem())); // ---> Me muestra la memoria disponible del pc.

function kb(bytes) { return bytes / SIZE }
function mb(bytes) { return kb(bytes) / SIZE }
function gb(bytes) { return mb(bytes) / SIZE }