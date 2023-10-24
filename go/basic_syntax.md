# Go: Basic syntax

### Variables

The `:=` syntax is shorthand for declaring and initializing a variable

```go
f := "apple"
// equal to
var f string = "apple"
```

### For

`for` is Go’s only looping construct.

```go
i := 1
for i <= 3 {
    fmt.Println(i)
    i = i + 1
}

for j := 7; j <= 9; j++ {
    fmt.Println(j)
}

for {
    fmt.Println("loop")
    break
}

for n := 0; n <= 5; n++ {
    if n%2 == 0 {
        continue
    }
    fmt.Println(n)
}
```

### If/Else

Any variables declared in this statement are available in the current and all subsequent branches.

```go
if num := 9; num < 0 {
    fmt.Println(num, "is negative")
} else if num < 10 {
    fmt.Println(num, "has 1 digit")
} else {
    fmt.Println(num, "has multiple digits")
}

// Output: 9 has 1 digit
```

### Switch

`switch` without an expression is an alternate way to express if/else logic

```go
// Normal
i := 2
switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
}

// without an expression
t := time.Now()
switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
}
```

### Array

An array type definition specifies a length and an element typ

```go
var a [4]int
a[0] = 1
i := a[0] // i == 1

// An array literal can be specified like so:
b := [2]string{"Penn", "Teller"}
// Or, you can have the compiler count the array elements for you:
b := [...]string{"Penn", "Teller"}
```

### Slice

Unlike an array type, a slice type has no specified length.

```go
letters := []string{"a", "b", "c", "d"}
```

Or can be created with the built-in function called `make`

```go
func make([]T, len, cap) []T
```

A slice can also be formed by “slicing” an existing slice or array

```go
b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
// b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
// b[:2] == []byte{'g', 'o'}
// b[2:] == []byte{'l', 'a', 'n', 'g'}
// b[:] == b
```

#### Growing slices (the copy and append functions)

```go
// built-in copy function
func copy(dst, src []T) int
// built-in append function
func append(s []T, x ...T) []T
```

```go
t := make([]byte, len(s), (cap(s)+1)*2)
copy(t, s)
s = t

// use `append`
a := make([]int, 1)
// a == []int{0}
a = append(a, 1, 2, 3)
// a == []int{0, 1, 2, 3}
```

### Maps

To create an empty map, use the builtin `make`: `make(map[key-type]val-type)`

```go
m := make(map[string]int)
m["k1"] = 7
m["k2"] = 13
fmt.Println(m) // map[k1:7 k2:13]

// Or with initialize values
n := map[string]int{"foo": 1, "bar": 2}
fmt.Println("map:", n) // map: map[bar:2 foo:1]
```

If the key doesn’t exist, the zero value of the value type is returned.

```go
fmt.Println(m["k3"]) // 0
```

The builtin delete removes key/value pairs from a map.

```go
delete(m, "k2")
fmt.Println("map:", m) // map: map[k1:7]
```

To remove all key/value pairs from a map, use the clear builtin.

```go
clear(m)
fmt.Println("map:", m) // map: map[]
```

### Range

`range` iterates over elements in a variety of data structures.

```go
// use range to sum the numbers in a slice
nums := []int{2, 3, 4}
sum := 0
// If we don’t need the index, ignore it with the blank identifier `_`.
for _, num := range nums {
    sum += num
}
fmt.Println("sum:", sum) // sum: 9
```

`range` on map iterates over key/value pairs.

```go
kvs := map[string]string{"a": "apple", "b": "banana"}
for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
    // a -> apple
    // b -> banana
}
```

`range` on strings iterates over Unicode code points

```go
for i, c := range "hello" {
    fmt.Println(i, c)
    // 0 104
    // 1 101
    // 2 108
    // 3 108
    // 4 111
}
```

### Multiple Return Values

```go
func vals() (int, int) {
    return 3, 7
}

func main() {
    a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
}
```

The `(int, int)` in this function signature shows that the function returns 2 ints.

### Variadic Functions

can be called with any number of trailing arguments

```go
func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("len:", len(nums), "total:", total)
}

func main() {
	sum(1, 2, 3) // len: 3 total: 6
}
```

Apply multiple args in a slice to a variadic function using `func(slice...)`.

```go
func main() {
	nums := []int{1, 2, 3, 4}
	sum(nums...) // len: 4 total: 10
}
```

### Closures

```go
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
```

We call `intSeq`, assigning the result (a function) to nextInt.

```go
func main() {
	// This function value captures its own i value, which will be updated each time we call nextInt.
	nextInt := intSeq()
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	// The state is unique to that particular function.
	newInts := intSeq()
	fmt.Println(newInts()) // 1
}
```

### Pointers

```go
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1

	zeroval(i)
	fmt.Println("zeroval:", i) // 1 -> value was not impacted

	zeroptr(&i)
	fmt.Println("zeroptr:", i) // 0 -> value was changed

	fmt.Println("pointer:", &i)
}

```

The `&i` syntax gives the memory address of `i`

### Strings and Runes

A Go string is a read-only slice of bytes.

In Go, the concept of a character is called a `rune` - it’s an integer that represents a Unicode code point.

Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.

```go
const s = "สวัสดี"
fmt.Println(len(s)) // 18, instead of 4
```

Indexing into a string produces the raw byte values at each index

```go
for i := 0; i < len(s); i++ {
    fmt.Printf("%x ", s[i])
    // e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5
}
```

Use `RuneCountInString` to count how many `runes` are in string

```go
 fmt.Println(utf8.RuneCountInString(s)) // 6
```

Why 6 instead of 4? Some characters are represented by multiple UTF-8 code point.

```go
for i, w := 0, 0; i < len(s); i += w {
    runeValue, width := utf8.DecodeRuneInString(s[i:])
    fmt.Printf("%#U starts at %d\n", runeValue, i)
    w = width
}
// U+0E2A 'ส' starts at 0
// U+0E27 'ว' starts at 3
// U+0E31 'ั' starts at 6
// U+0E2A 'ส' starts at 9
// U+0E14 'ด' starts at 12
// U+0E35 'ี' starts at 15
```

### Structs

Go’s structs are typed collections of fields.

```go
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}
```

This syntax creates a new struct.

```go
person{"Bob", 20}                   // {Bob 20}
person{name: "Alice", age: 30}      // {Alice 30}
newPerson("Jon")                    // &{Jon 42}
```

Structs are mutable.

### Methods

Go supports methods defined on struct types.

```go
type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println(r.area())  // 50
	fmt.Println(r.perim()) // 30
}
```

### Interfaces

Interfaces are named collections of method signatures.

```go

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
```

To implement an interface in Go, we just need to implement all the methods in the interface

```go
func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    measure(r)

    c := circle{radius: 5}
    measure(c)
}
```

### Struct Embedding

Go supports embedding of `structs` and `interfaces` to express a more seamless composition of types.

```go
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base        // An embedding looks like a field without a name.
	str string
}
```

```go
co := container{
    base: base{num: 1},
    str:  "some name",
}

fmt.Println(co.num, co.base.num) // both return 1

fmt.Println(co.describe()) // base with num=1

```

### Generics

Example of a generic function

```go
func MapKeys[K comparable, V any](m map[K]V) []K {
    r := make([]K, 0, len(m))
    for k := range m {
        r = append(r, k)
    }
    return r
}
```

- `MapKeys` takes a map of any type and returns a slice of its keys
- K has the `comparable` constraint => can compare values of this type with the `==` and `!=` operators.
- V has the `any` constraint, meaning that it’s not restricted in any way.

```go
var m = map[int]string{1: "2", 2: "4", 4: "8"}

fmt.Println(MapKeys(m)) // [4 1 2]
```

Example of a generic type, List is a singly-linked list with values of any type.

```go
type element[T any] struct {
    next *element[T]
    val  T
}

type List[T any] struct {
    head, tail *element[T]
}
```

The type is `List[T]`, not `List`

```go
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}
func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	fmt.Println(lst.GetAll()) // [10 13]
}
```

### Errors

By convention, errors are the `last return value` and have type error, a built-in interface.

```go
import "errors"

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
}

// f1 worked: 10
// f1 failed: can't work with 42
```

It’s possible to use custom types as errors by implementing the `Error()` method on them

```go
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}

	return arg + 3, nil
}

func main() {
	_, e := f2(42)
	if err, ok := e.(*argError); ok {
		fmt.Println(err.arg)    // 42
		fmt.Println(err.prob)   // can't work with it
	}
}
```
