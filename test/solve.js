solve = (input) => {
    let map = {"x": 0, "A": 0, "B": 1, "C": 3, "D": 5};
    let data = input.split("");
    this.answer = data.map(x => map[x]).reduce((x, y) => x + y);
    for (let i = 0; i < data.length; i += 3) {
        let x = data[i] === 'x' ? 1 : 0;
        x += data[i + 1] === 'x' ? 1 : 0;
        x += data[i + 2] === 'x' ? 1 : 0;
        if (x === 1) {
            this.answer += 2;
        } else if (x === 0) {
            this.answer += 6;
        }
    }
}

solve()
