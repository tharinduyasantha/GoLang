# Go Crash Course

### Declaring Variables
___
#### Data Types
    1. int
    2. float32
    3. string
    4. bool

#####  Using var
    var variablename type = value
    var x int = 1

##### Using :=
    variablename := value
    a := "String"
    var a, b, c, d int = 1, 3, 5, 7

### Constants
    const CONSTNAME type = value
    const PI = 3.14

### Arrays
___
    Arrays => Fixed , Must be defined

##### With var

```go
var array_name = [length]datatype{values} // here length is defined
var my_array = [5]string{"1", "2", "3", "4", "5"}
or
var array_name = [...]datatype{values} // here length is inferred
var my_array = [...]string{"1", "2", "3", "4", "5"}
```
##### With :=

```go
array_name := [length]datatype{values} // here length is defined
my_array := [5]string{"1", "2", "3", "4", "5"}
or
array_name := [...]datatype{values} // here length is inferred
my_array := [...]string{"1", "2", "3", "4", "5"}
```

### Slices
___
    Slices=> Not Fixed , Must be defined using :=
#### Syntax
```go
slice_name := []datatype{values}
myslice := []int{1,2,3}
```

#### Append and Modify
    Append
    myslice1 := []int{1, 2, 3, 4, 5, 6}
    myslice1 = append(myslice1, 20, 21) => myslice1 = [1 2 3 4 5 6 20 21]
    --
    Modify
    prices := []int{10,20,30}
    prices[2] = 50 => prices = [10,20,50]

### Operators
___
| Operator          | Description                             | Example                         | Result |
|-------------------|-----------------------------------------|---------------------------------|--------|
| `+`               | Addition                                | `5 + 3`                         | `8`    |
| `-`               | Subtraction                             | `5 - 3`                         | `2`    |
| `*`               | Multiplication                          | `5 * 3`                         | `15`   |
| `/`               | Division                                | `6 / 3`                         | `2`    |
| `%`               | Modulus (remainder)                     | `5 % 3`                         | `2`    |
| `++`              | Increment                               | `var a = 5; a++`                | `6`    |
| `--`              | Decrement                               | `var a = 5; a--`                | `4`    |
| `==`              | Equal to                                | `5 == 3`                        | `false`|
| `!=`              | Not equal to                            | `5 != 3`                        | `true` |
| `>`               | Greater than                            | `5 > 3`                         | `true` |
| `<`               | Less than                               | `5 < 3`                         | `false`|
| `>=`              | Greater than or equal to                | `5 >= 3`                        | `true` |
| `<=`              | Less than or equal to                   | `5 <= 3`                        | `false`|
| `&&`              | Logical AND                             | `true && false`                 | `false`|
| `||`              | Logical OR                              | `true || false`                 | `true` |
| `!`               | Logical NOT                             | `!true`                         | `false`|
| `&`               | Bitwise AND                             | `5 & 3`                         | `1`    |
| `|`               | Bitwise OR                              | `5 | 3`                         | `7`    |
| `^`               | Bitwise XOR                             | `5 ^ 3`                         | `6`    |
| `<<`              | Left shift                              | `5 << 1`                        | `10`   |
| `>>`              | Right shift                             | `5 >> 1`                        | `2`    |
| `+=`              | Add and assign                          | `var a = 5; a += 3`             | `8`    |
| `-=`              | Subtract and assign                     | `var a = 5; a -= 3`             | `2`    |
| `*=`              | Multiply and assign                     | `var a = 5; a *= 3`             | `15`   |
| `/=`              | Divide and assign                       | `var a = 6; a /= 3`             | `2`    |
| `%=`              | Modulus and assign                      | `var a = 5; a %= 3`             | `2`    |
| `&=`              | Bitwise AND and assign                  | `var a = 5; a &= 3`             | `1`    |
| `|=`              | Bitwise OR and assign                   | `var a = 5; a |= 3`             | `7`    |
| `^=`              | Bitwise XOR and assign                  | `var a = 5; a ^= 3`             | `6`    |
| `<<=`             | Left shift and assign                   | `var a = 5; a <<= 1`            | `10`   |
| `>>=`             | Right shift and assign                  | `var a = 5; a >>= 1`            | `2`    |

### Conditions
___
#### If
```go
if condition {
  // code to be executed if condition is true
}
```
#### IF Else
```go
if condition {
    // code to be executed if condition is true
} else{
    // code to be executed if condition is false
}
```

#### Nested If
```go
if condition1 {
   // code to be executed if condition1 is true
  if condition2 {
     // code to be executed if both condition1 and condition2 are true
  }
}
```
#### Switch
```go
switch expression {
case x:
// code block
default:
// code block
}
```
#### Multi-case Switch
```go
switch expression {
case x,y:
// code block if expression is evaluated to x or y
default:
// code block
}
```
### Loops (Only loop available)
___
#### Default For Loop
```go
for statement1; statement2; statement3 {
   // code to be executed for each iteration
}
```
#### Default For Loop
```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
```
#### Collection Iteration For Loop

```go
my_array := [...]string{"1", "2", "3", "4", "5"}
for i, v := range my_array {
    fmt.Println(i, v)
}
prints
0 1
1 2
2 3
. .
```
### Functions
___
```go
func FunctionName() {
  // code to be executed
}
```
#### Functions with params
```go
func FunctionName(param1 type, param2 type, param3 type) {
  // code to be executed
}
```
#### Functions with return types
```go
func FunctionName(param1 type, param2 type) type {
  // code to be executed
  return output
}
```

#### Named returns fuctions
```go
func rectangleDimensions(length, width float64) (area, perimeter float64) {
    area = length * width
    perimeter = 2 * (length + width)
    return // Named return values are automatically returned
}
area, perimeter := rectangleDimensions(5.0, 3.0)
fmt.Println("Area:", area, "Perimeter:", perimeter)

prints
Area: 15 Perimeter: 16
```
#### Anonymous Functions
```go
message := func(name string) string {
		return "Hello, " + name
	}("Go")
fmt.Println(message)
```
