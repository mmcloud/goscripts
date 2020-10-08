# goscripts
A general collection of scripts used to teach myself go.


## Layout

All projects are found in the src file.

# Projects

## hello

Module Reference: [golang.org](https://blog.golang.org/using-go-modules)

Workflow Notes: 

- Create new module, init go.mod file
    ```
    go mod init github.com/mmcloud/goscript/src/hello
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

## Binary Tree
[wiki](https://en.wikipedia.org/wiki/Binary_search_tree)