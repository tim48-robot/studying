const fs = require("fs");
const http = require("http");
const url = require("url");


http.createServer((req,res)=>{
    let realURL = new URL(`http://${req.headers.host}${req.url}`)
    let path = realURL.pathname
    console.log(path)
    if (path === "/"){
        console.log(path)
        fs.readFile("views/index.html", (err, data) => {
            if (err){
                res.writeHead(404, {
                    "Content-Type" :"text/plain"
                })
                res.end("error")
            } else {
                res.writeHead(200, {
                    "Content-Type": "text/html"
                })
                res.write(data);
                res.end()
            }
        })
    } else if (path === "/about") {
        fs.readFile("views/about.html", (err, data) => {
            if (err){
                res.writeHead(404, {
                    "Content-Type" :"text/plain"
                })
                res.end("error")
            } else {
                res.writeHead(200, {
                    "Content-Type": "text/html"
                })
                res.write(data);
                res.end()
            }
        })
    } else if (path === "/contact-me"){
        fs.readFile("views/contact-me.html", (err, data) => {
            if (err){
                res.writeHead(404, {
                    "Content-Type" :"text/plain"
                })
                res.end("error")
            } else {
                res.writeHead(200, {
                    "Content-Type": "text/html"
                })
                res.write(data);
                res.end()
            }
        })
    } else {
        fs.readFile("views/404.html", (err, data) => {
            if (err){
                res.writeHead(404, {
                    "Content-Type" :"text/plain"
                })
                res.end("error")
            } else {
                res.writeHead(200, {
                    "Content-Type": "text/html"
                })
                res.write(data);
                res.end()
            }
        })
    }

})
.listen(8080);

