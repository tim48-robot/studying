const fs = require("fs/promises");

async function log (message){
    let today = new Date();
    today = today.toISOString();
    const [date, tim] = today.split("T");
    const time = tim.split(".")[0];
    fs.appendFile("applog.txt", `[${date} ${time}] ${message}\n`);
}

log("Server Started");
log("User logged in");
log("Database connected");