const fs = require("fs");
const readLine = require('readline');


async function readFile(){
    const fileStream = fs.createReadStream('output.txt');
    const rl = readLine.createInterface({
        input: fileStream,
        crlfDelay: Infinity
    })

    let lineNumber = 1;
    for await (const line of rl){
        console.log(`${lineNumber}: ${line}`);
        lineNumber++;
    }
}

readFile();
