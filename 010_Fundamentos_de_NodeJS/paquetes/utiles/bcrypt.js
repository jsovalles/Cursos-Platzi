const bcrypt = require('bcrypt');

const password = '1234Segura!';

bcrypt.hash(password, 5)
    .then(hash => {
        console.log(hash)
        return hash
    })
    .then(hash => {
        return bcrypt.compare(password, hash)
    })
    .then(res => {
        console.log(res)
    })
    .catch(err => {
        console.log(err)
    })