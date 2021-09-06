## Exercise: Errors

Copy your `Sqrt` function from the earlier exercise and modify it to return an error value.
`Sqrt` should return a non-nil error value when given a negative number, as it doesn't support complex numbers.
Create a new type

```go
type ErrNegativeSqrt float64
```

and make it an `error` by giving it a

```go
func (e ErrNegativeSqrt) Error() string
```

method such that `ErrNegativeSqrt(-2).Error()` returns "cannot Sqrt negative number: -2".
Change your `Sqrt` function to return an `ErrNegativeSqrt` value when given a negative number.

**Note**: A call to `fmt.Sprint(e)` inside the Error method will send the program into an infinite loop. You can avoid this by converting `e` first: `fmt.Sprint(float64(e))`. **Why?**

**Answer**: `fmt.Sprint(e)` will call `e.Error()` to convert the value `e` to a string.
If the `Error()` method calls `fmt.Sprint(e)`, then the program recurses until out of memory.
You can break the recursion by converting the `e` to a value without a `String` or `Error` method.
*Note* using `fmt.Sprintf("%f",e)` works, because this string format converts the `e` to float and does not call the `Error()` method to print it out.
