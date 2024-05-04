## Go Learning

1. Initiate the go application
   `go mod init <module path>`
   This creates a new module

   - Initialized a go.mod file
   - go.mod describes the module: with name/modulepath and go version used in the program
   - The module path is also the import path

2. All our code must belong to a package

   - The first statement in Go file must be "package ..."

3. Where to start the program? Where is the entrypoint? Go program needs.

   ```
   func main() {
       Print("Hello World")
   }
   ```

   you can have only one main function in a go program.

4. to use Print or any other functions, we need to import some inbuild packages.

   - Go programs are organized into packages
   - Go's standard library, provides different core packages for us to use
   - "fmt" is one of these, which you can use by importing it.

   `import "fmt"`

5. to execute the program `go run main.go`
