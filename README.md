# GoInterpreter

This project is a personal implementation of the code examples found in the book "Writing An Interpreter In Go" by Thorsten Ball.

The book serves as a comprehensive guide to building an interpreter from scratch, and this project aims to provide a hands-on experience by recreating the concepts and code discussed in the book.

---

## Examples

### Functions

```
>> let multiply = fn(x, y) { x * y };
>> multiply(50 / 2, 1 * 2)
50
```

### Support for hashes and arrays

```
>>let people = [{"name": "Alice", "age": 24}, {"name": "Anna", "age": 28}];
>>people[0]["name"];
Alice
```

```
>> sum([1, 2, 3, 4, 5]);
15
```

### Builtin functions

```
>> let myArray = ["one", "two", "three"];
>> len(myArray)
3
```

---

##### Acknowledgments
This project is heavily inspired by the book "Writing An Interpreter In Go" by Thorsten Ball. 