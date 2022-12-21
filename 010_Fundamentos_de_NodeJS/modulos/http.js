const http = require('http');

http.createServer(router).listen(3000);

function router(req, res) {
    console.info('Nueva Petición');
    console.log(req.url);
    res.writeHead(200, { 'Content-Type': 'text/plain' });
    res.write('Ruta => ' + req.url);
    switch (req.url) {
        case '/hola':
            res.write('\n\rHola desde la ruta ' + req.url);
            res.end();
            break;
        default:
            res.write('\n\rError 404');
            res.end();
    }
}
console.info('escuchando 3000')