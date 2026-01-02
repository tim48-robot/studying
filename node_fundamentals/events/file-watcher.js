const event = require("events");
const fs = require("fs");
const EventEmitter = new event();

EventEmitter.on("fileCreated", (res) => {
    console.log(`File ${res.name} berhasil ditulis, pada ${res.timestamp}`)
})


EventEmitter.on("fileRead", (res) => {
    console.log(`File ${res.name} berhasil dibaca, dengan text ${res.text} pada ${res.timestamp}`)
})

function writeFile(name, text){
    fs.writeFile(name, text, (err) =>{
        if (err){
            console.log(err);
        } else {
            EventEmitter.emit("fileCreated", {
                name: name,
                text: text,
                timestamp: new Date().toISOString()
            });
        }
    })
}

function readFile(name){
    fs.readFile(name, (err,data) =>{
        if (err){
            console.log(err);
        } else {
            EventEmitter.emit("fileRead", {
                name: name,
                text: data,
                timestamp: new Date().toISOString()
            });
        }
    })
}

writeFile("keren.txt", "kapan kapan ya")

setTimeout(readFile, 1000, "keren.txt")
// or
// setTimeout(() => {
//   readFile("keren.txt")}, 1000);