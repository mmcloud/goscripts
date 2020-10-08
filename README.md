# goscripts
A general collection of scripts


## Layout

All projects are found in the src file.


# Notes


Workflow Notes: 

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
