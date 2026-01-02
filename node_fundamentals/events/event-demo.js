const events = require("events");

const EventEmitar = new events();
const data = {
    name: "John",
    id: 123
}

for (let i=0; i<3; i++){
    EventEmitar.on("userLoggedIn", (data =>{
        console.log(`User ${data.name} logged in with id ${data.id}`)
    }))
}

EventEmitar.emit("userLoggedIn", data);
