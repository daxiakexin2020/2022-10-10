var Test = /** @class */ (function () {
    function Test(f, s, age) {
        this.firstname = f + s;
        this.age = age;
    }
    Test.prototype.demo01 = function () {
    };
    Test.prototype.demo02 = function () {
    };
    Test.prototype.demo03 = function () {
    };
    return Test;
}());
var test = new Test("zz", "kx", 18);
test.demo01();
function greeter(person) {
    return "Hello, " + person;
}
var user = "Jane User";
document.body.textContent = greeter(user);
function greeter2(person2) {
}
var user2 = [0, 1, 2];
greeter2(user2);
