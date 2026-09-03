# Comma-Ok Syntax in Go

- Comma-ok syntax returns a value and a boolean.
- The general form is:

```go
value, ok := operation
```

- `value` contains the result.
- `ok` tells whether the operation succeeded.
- The meaning of `ok` depends on the operation.

## Map Lookup

- Map lookup uses comma-ok to check whether a key exists.

```go
scores := map[string]int{
    "Ada":  0,
    "John": 90,
}

score, ok := scores["Ada"]
if ok {
    fmt.Println("Score:", score)
} else {
    fmt.Println("Student not found")
}
```

- `ok == true` means the key exists.
- `ok == false` means the key does not exist.
- A missing key returns the zero value of the map value type.
- For an `int` map, the zero value is `0`.
- This is why checking only the value is not enough.
- Both of these lookups return `0`, but their `ok` values differ:

```go
existing, exists := scores["Ada"]   // 0, true
missing, exists := scores["Grace"]  // 0, false
```

- Use `_` when the value is not needed:

```go
_, exists := scores["Ada"]
if exists {
    fmt.Println("Key exists")
}
```

- A lookup in a nil map is safe.
- A nil map lookup returns the zero value and `false`.

## Channel Receive

- Channel receive uses comma-ok to detect whether a channel is closed.

```go
jobs := make(chan int, 1)
jobs <- 10
close(jobs)

job, ok := <-jobs
fmt.Println(job, ok) // 10 true

job, ok = <-jobs
fmt.Println(job, ok) // 0 false
```

- `ok == true` means a value was received.
- `ok == false` means the channel is closed and no values remain.
- A closed channel returns the element type's zero value.
- An open channel with no value does not return `false`.
- Receiving from an open empty channel blocks until a value is sent or the channel is closed.
- Buffered values are received before `ok` becomes `false`.
- Only the sender should close a channel.
- A channel must not be closed more than once.

- A `for range` loop automatically stops when the channel is closed:

```go
for job := range jobs {
    fmt.Println("Processing job:", job)
}
```

## Type Assertion

- Type assertion checks the concrete type stored inside an interface.

```go
var value any = "Hello"

text, ok := value.(string)
if ok {
    fmt.Println("String value:", text)
} else {
    fmt.Println("Value is not a string")
}
```

- `ok == true` means the value has the requested type.
- `ok == false` means the value has a different type.
- The safe form does not panic when the type is incorrect.

```go
var value any = "Hello"

number, ok := value.(int)
fmt.Println(number, ok) // 0 false
```

- A type assertion is not a type conversion.
- An `int` does not become a `string` through a type assertion.
- The boolean checks the type, not whether the value is nil.

## If Statement Scope

- Comma-ok values can be declared inside an `if` statement.
- The variables are available only inside that `if` statement.

```go
if score, ok := scores["Ada"]; ok {
    fmt.Println("Score:", score)
}
```

- Declare the variables outside when they are needed later:

```go
score, ok := scores["Ada"]
if ok {
    fmt.Println(score)
}
```

## Common Rules

- Use `:=` when declaring new variables.
- Use `=` when the variables already exist.
- The number and order of variables must match the returned values.
- Use `_` to ignore a value that is not needed.
- Always check `ok` when a zero value could be a valid result.
- Comma-ok is not the same as error handling.
- Functions that need to explain a failure usually return an error:

```go
result, err := doWork()
if err != nil {
    fmt.Println("Error:", err)
}
```

## Summary

- `value, ok := myMap[key]`
  - Checks whether a map key exists.
- `value, ok := <-channel`
  - Checks whether a channel value was received before closure.
- `value, ok := item.(Type)`
  - Checks whether an interface contains a specific type.
