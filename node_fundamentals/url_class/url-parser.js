const url = require("url");
const readline = require("readline");
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
})


rl.on("line", (input =>{
    let myurl = "";
    try {
        myurl = new URL(input);
    } catch (err) {
        console.log("Invalid URL");
        rl.close();
        return;
    }
    console.log(`Protocol: ${myurl.protocol}`)
    console.log(`Hostname: ${myurl.hostname}`)
    console.log(`Port: ${myurl.port}`)
    console.log(`Pathname: ${myurl.pathname}`)

    console.log("Search params: ")
    const objectifiedParams = Object.fromEntries(myurl.searchParams);
    const allKeys = Object.keys(objectifiedParams);

    if (myurl.searchParams){
        for (let i=0; i<allKeys.length; i++){
            console.log(`  ${allKeys[i]}: ${objectifiedParams[allKeys[i]]}`);
        }
    }
    // console.log(myurl.searchParams);


    // let perParam = myurl.search.split("?")[1];
    // if (perParam.includes("&")){
    //     perParam = perParam.split("&");
    // }
    // for (let i = 0; i<perParam.length; i++){
    //     console.log(`${perParam[i].split("=")[0]}: ${perParam[i].split("=")[1]}`)
    // }

    rl.close();
}))

