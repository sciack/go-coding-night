# Coding night

During the coding night we explore go lang so this is the suggested software to have installed:

* go lang (any available, should not make a difference)
* Ide 
    * VSCode + Go plugin
    * JetBrain go (or IntellijUltimate + go)
* You

I did some training with Mac and Linux, don't know how go works in Windows, last hope WLS2

## Installation
Official guide: https://golang.org/doc/install

### Mac instruction

Official site tell you to download the executables from the site and unpack them, but a most handy way is to use brew.

Installing go with brew:

```|shell
brew install go
```

### Linux
As for mac official site tell you to download and install the tarball, but usually is 
already packaged with the distribution

### Windows
Don't know if choco works well with windows, but could be a different way to install it.

## Create a module

Each directory in go is a module and in order to simplify building etc. is better to start with a module:

```
go mod init coding-night
```

could be an example.

after you can create a file (usually `main.go`) that contains your code and `main_test.go` that contains the test

To run the code:

```
go run main.go
```

or 

```
go build 
```
and execute the generated executables.

To execute the tests:

```
go test
```

## Documentation


* https://golang.org/doc/
* https://golangdocs.com/
* https://devhints.io/go