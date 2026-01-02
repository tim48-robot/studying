const http = require("http");
const url = require("node:url");

http.createServer((req, res) => {
    const parsedUrl = new URL(req.url, "https://localhost:3000");
    const pathname = parsedUrl.pathname;
    

    if (req.method === "GET" && pathname === "/"){
        res.writeHead(200, {
            "Content-Type": "text/plain"
        })
        res.write("Home Page");
        res.end();
    } else if (req.method === "GET" && pathname === "/about"){
        res.writeHead(200, {
            "Content-Type": "text/plain"
        })
        res.write("About-Page");
        res.end();
    } else if (req.method === "GET" && pathname === "/api/time"){
        res.writeHead(200, {
            "Content-Type": "application/json"
        })
        res.write(JSON.stringify(new Date()))
        res.end();
    } else if (pathname != "/" && pathname != "/about" && pathname != "/api/time") {
        res.writeHead(404, {
            "Content-Type": "text/plain"
        })
        res.write("404 - PAGE NOT FOUND");
        res.end();
    }
}).listen(3000);
