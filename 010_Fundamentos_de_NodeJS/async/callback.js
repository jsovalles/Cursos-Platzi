//La asincronia se puede generar mediante en setTimeout
console.log('Iniciando proceso...');
function soyAsincrono(elCallback){
    console.log("Asigno setTimeout para volver asincrona una función como esta misma: \n " + soyAsincrono);
    setTimeout(function(){
    console.log("Pasaron 3 segundos y me ejecuté");
    elCallback();
    }, 3000)
};

soyAsincrono(function(){console.log("Después de esto demuestro que Soy el primer Callback")});


function hola(nombre, miCallback) {
    setTimeout(function () {
        console.log('Hola, '+ nombre);
        miCallback(nombre);
    }, 1500);
};

function adios(nombre, otroCallback) {
    setTimeout(function() {console.log('Adios', nombre);otroCallback();}, 5000);
};


hola('Alejandro',function (nombre) {
    adios(nombre,function() {
        console.log('Terminando proceso...');
    });
});

hola('Alejandro estás probando  "hola" las funciones independientemente, las pasas vacías',function () {});
adios('Alejandro estás probando "adios" las funciones independientemente, las pasas vacías',function () {});