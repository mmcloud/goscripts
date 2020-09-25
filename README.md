# goscripts
A general collection of scripts


## Layout

All projects are found in the src file.

# Projects

## hello

Module Reference: [golang.org](https://blog.golang.org/using-go-modules)

Workflow Notes: 

- Create new module, init go.mod file
    ```
    go mod init mmcloud.com/hello
    ```
- Package building commands add to go.mod
    ```
    go build
    go test
    ```
- List module dependencies
    ```
    go list -m all
    ```
- Get module dependency
    ```
    go get
    ```
- Remove unused dependency
    ```
    go mod tidy
    ```