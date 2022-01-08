# Go-Learning

### Packages
Go programs are constructed by linking together packages. A package in turn is constructed from one or more source files that together declare constants, types, variables and functions belonging to the package and which are accessible in all files of the same package. Those elements may be exported and used in another package.

### Modules
A module is a collection of Go packages stored in a file tree with a go.mod file at its root.The go.mod file defines the module’s module path, which is also the import path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build.  

#### *Creating a new module*
`go mod init example.com/hello`

#### *Adding a dependency*
The primary motivation for Go modules was to improve the experience of using (that is, adding a dependency on) code written by other developers.
we can add the dependecy with `import` statment. 
eg:`import "rsc.io/quote"` 

` go list -m all` : lists the current module and all its dependencies

#### *Upgrading dependencies*
`go list -m -versions rsc.io/sampler` : show versions
`go get rsc.io/sampler@v1.3.1` : install specific version 
