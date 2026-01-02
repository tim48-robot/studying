const readline = require("readline");
const url = require("url");
const params = new URLSearchParams({
    name: "John", age: 30, city: "NYC"
});

console.log(params.toString());