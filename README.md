# Learning Go

## Resources used:

- [boot.dev](https://boot.dev)
- Learn Go with tests - [quii.gitbook.io](https://quii.gitbook.io)

### Development Cycle

- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor

### Things to remeber while testing

- A test function needs to be in a file with a name like `xxx_test.go`
- The test function must start with the word `Test`
- The test function takes one argument only `t *testing.T`
- To use the `*testing.T` type, you need to `import "testing"`
