const fs = require("fs/promises");

async function copyFile(){
    try{
        const res = await fs.readFile("source.txt") 
            try{
                fs.writeFile("destination.txt", res);
                console.log("File copied Succesfully~!~");
            }
            catch (err) {
                console.log(err);
            }
    }
    catch (err){
        console.log(err);
    }
}

copyFile()