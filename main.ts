const wait = (ms: number) => new Promise(r => setTimeout(r, ms));

let value = 0
const ticker =  async() => {
    await wait(200);
    console.clear();
    if (value < 10) { 
        value++;
        ticker();
        console.log(value);
    }
}

ticker();