# Go Interpreter and Compiler

This project is a personal implementation of the code examples found in the book "Writing An Interpreter In Go" and "Writing A Compiler In Go" by Thorsten Ball.

The book serves as a comprehensive guide to building an interpreter and compiler from scratch, and this project aims to provide a hands-on experience by recreating the concepts and code discussed in the book.

---
# Interpreter

## Examples

### Functions

```bash
>> let multiply = fn(x, y) { x * y };
>> multiply(50 / 2, 1 * 2)
50
```

### Support for hashes and arrays

```bash
>>let people = [{"name": "Alice", "age": 24}, {"name": "Anna", "age": 28}];
>>people[0]["name"];
Alice
```

```bash
>> sum([1, 2, 3, 4, 5]);
15
```

### Builtin functions

```bash
>> let myArray = ["one", "two", "three"];
>> len(myArray)
3
```

# Compiler

## Examples

### Functions

```bash
>>let one = fn() { 1; };
>>let two = fn() { 2; };
>>one() + two()
3
```

### Support for hashes and arrays

```bash
>>{"one": 1, "two": 2, "three": 3}["o" + "ne"]
1 
```

### Builtin functions
    
```bash
>>let array = [1, 2, 3];
>>first(rest(push(array, 4)))
2
```

### Closures

```bash
>>let newAdder = fn(x) { fn(y) { x + y }; };
>>let addTwo = newAdder(2);
>>addTwo(3)
5
```
 
---
##### Acknowledgments
This project is heavily inspired by the book "Writing An Interpreter In Go" by Thorsten Ball. 