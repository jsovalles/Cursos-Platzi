const sharp = require('sharp');

//ruta de la imagen a modificar
sharp('original.png')
  .resize(500) //scalamos la imagen a 500 de alto
  .grayscale() //otorgamos una escala de grises
  .toFile('original-small.jpg'); //generamos el archivo modificado