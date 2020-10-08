# goscripts

A general collection of scripts

## Layout

All projects are found in the src file.

## Workflow Notes

- Create new module, init go.mod file

    ```bash
    go mod init mmcloud.com/hello
    ```

- Package building commands add to go.mod

    ```bash
    go build
    go test
    ```

- List module dependencies

    ```bash
    go list -m all
    ```

- Get module dependency

    ```bash
    go get
    ```

- Remove unused dependency

    ```bash
    go mod tidy
    ```

## Testing and Benchmarking

Store in tests and benchmarks *_test file of module.

- Create UnitCase

```go
func TestRepeat(t *testing.T) {
    repeated := Repeat("a", 5)  
    expected := "aaaaa"

    if repeated != expected {
        t.Errorf("expected '%q' but got '%q'", expected, repeated)
    }
}
```

- Create BenchMarks

```go
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ { // b.N to iterate for benchmark
        Repeat("a", 5)
    }
}

```

- Parallel test, use t.Run

```go
func TestArea(t *testing.T) {

    t.Run("rectangles", func(t *testing.T) {
        rectangle := Rectangle{12, 6}
        got := Area(rectangle)
        want := 72.0

        if got != want {
            t.Errorf("got %g want %g", got, want)
        }
    })

    t.Run("circles", func(t *testing.T) {
        circle := Circle{10}
        got := Area(circle)
        want := 314.1592653589793

        if got != want {
            t.Errorf("got %g want %g", got, want)
        }
    })
}
```

- Go does not let you use equality operators. Instead reflect.

```go
got := SumAll([]int{1, 2}, []int{0, 9})
want := []int{3, 9}

if !reflect.DeepEqual(got, want) {
    t.Errorf("got %v want %v", got, want)
}
```

## Syntax & comprehension

- Use make to create [slices](https://tour.golang.org/moretypes/13)

```go 
a := make([]int, 5) //Makes an zero array of length 5
```

- To specify a capacity, pass a third argument to make:

```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
```

- Use make to create [maps](https://blog.golang.org/maps)

```go
m = make(map[string]int)
```

- Interfaces

```go
type Shape interface {
    Area() float64
}
```

- Testing Multiple Structs

```go
    func TestAreaInterfaceRefactored(t *testing.T) {
        areaTests := []struct {
            shape Shape
            want  float64
        }{
            {Rectangle{10, 12}, 120},
            {Circle{10}, 314.1592653589793},
        }
        for _, tt := range areaTests {
            got := tt.shape.Area()
            if got != tt.want {
                t.Errorf("got %g want %g", got, tt.want)
            }
        }
    }
```