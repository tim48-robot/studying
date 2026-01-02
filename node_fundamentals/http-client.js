// const http = require("http");
// const axios = require("axios");
// axios.get("https://jsonplaceholder.typicode.com/users/1")
// .then(res => {
//     console.log(res);
// }).catch(err => {
//     console.log(err);
// })

const https = require("https");
const options = {
    hostname: "jsonplaceholder.typicode.com",
    path: "/users/1",
    method: "GET"
}

const req = https.request(options, res => {
    console.log(res.statusCode);
    let body = "";
    
    res.on("data", d => {
        body += d;
    })

    res.on("end", d=> {
        process.stdout.write(body);
    })
}
)

req.on("error", msg => {
    console.log(msg);
})

req.end();