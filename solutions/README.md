# Intro to Go Lab 1: Imperative Programming - Solutions

### Question 1a

See `hello/hello.go`

### Question 1a

See `hello/hello.go`

### Question 2

See `quiz/quiz.go`

### Question 3a

See `sequences/sequences.go`

### Question 3b

Arrays are passed by value and have fixed length, where as sdlices reference a section of an underlying array and are dynamic. 

### Question 3c

`mapArray` no longer works because the type `[5]int` doesn't match type `[3]int`. 

### Question 3d

```go
newSlice := intsSlice[1:3] // = [3, 4]
mapSlice(square, newSlice)
fmt.Println(newSlice)  // = [9, 16]
fmt.Println(intsSlice) // = [2 9 16 5 6]
```

The slice returned by (`:`) returns a new slice which still references the same underlying array.

### Question 3e

As slices contain pointers to arrays, if we modify the contents of a slice in a way that needs to change the underlying array such as changing the length with `double`, we need to reassign the slice to see those changes, this can be done by returning from `double`

### Question 3f

- When Passing to functions arrays are passed by value i.e. the entire array data is copied. Slices are passed by reference i.e. you pass a run-data structure which points to an underlying array. 
- With append() we cannot append to arrays as they have fixed size upon creation, but we can append to slices and the result needs to be returned, because append may allocate a new larger underlying array.
- In general arrays provide safety via immutability so are useful for fixed sized collections like a fixed transform matrix, and slices provide flexibility via mutability so are useful for dunamically sized collections like batch data tjat needs to be processed, overall arrays have limited use cases and 99% of the time we will want to use a slice.

### Question 4

See `gol/gol.go`


