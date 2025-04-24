[![wakatime](https://wakatime.com/badge/user/0d75cfc5-da70-41b7-b8c8-661ef9d8338b/project/9795795c-2f45-46f6-8cc7-6de33ef8f8f9.svg)](https://wakatime.com/badge/user/0d75cfc5-da70-41b7-b8c8-661ef9d8338b/project/9795795c-2f45-46f6-8cc7-6de33ef8f8f9)

# Learning Go

## Resources used:

- [boot.dev](https://boot.dev)
- Learn Go with tests - [quii.gitbook.io](https://quii.gitbook.io)

### Test Driven Development Cycle

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
