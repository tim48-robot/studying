import fs from 'fs/promises';
try {
    await fs.writeFile("output.txt", "This is line 1\n");
    await fs.appendFile("output.txt", "This is line 2\n");
    await fs.appendFile("output.txt", "This is line 3\n");
} catch (err) {
    console.log(err);
}


console.log(typeof fs.writeFile);