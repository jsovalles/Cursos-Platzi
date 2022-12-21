const myName = 'Juan';
const myAge = 12;
const suma = (a: number, b: number) => {
  return a + b;
};
suma(12, 23);

class Persona {
  constructor(private age: number, private name: string) {}

  getSummary() {
    return `My name is: ${this.name}, ${this.age}`;
  }
}

const juan = new Persona(14, 'Juan');
juan.getSummary();
