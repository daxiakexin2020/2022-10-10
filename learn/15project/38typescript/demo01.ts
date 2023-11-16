interface Demo {
    Eat(type: string)

    Name: string
}


export class Test {
    firstname: String
    age: number

    public constructor(f: string, s: string, age: number) {
        this.firstname = f + s
        this.age = age
    }

    public demo01() {
    }

    protected demo02() {
    }

    private demo03() {
    }
}


var test = new Test("zz", "kx", 18);
test.demo01()

export function greeter(person) {
    return "Hello, " + person;
}

let user = "Jane User";

document.body.textContent = greeter(user);

function greeter2(person2: number[]) {

}

let user2 = [0, 1, 2];

greeter2(user2);

