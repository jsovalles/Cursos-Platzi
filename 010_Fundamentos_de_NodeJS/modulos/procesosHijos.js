const { exec, spawn } = require('child_process')

exec('node modulos/consola.js', (e, stdout) => {
    (e) ? console.log(e) : console.log(stdout)
})

let process = spawn('cmd.exe', ['/c','dir']);
process.stdout.on("data", data => console.log(data.toString()));
process.on("exit", () => console.log("process end"));