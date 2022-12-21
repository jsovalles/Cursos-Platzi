const moment = require('moment');

let ahora = moment();

let ahoraFormat = ahora.format('HH:mm:ss');

console.log(ahora.toString());
console.log(ahoraFormat);