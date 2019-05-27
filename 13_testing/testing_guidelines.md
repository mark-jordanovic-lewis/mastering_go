# Overview
- import the testing package
- use equalities and type comparisons from relevant libraries to perform asserts/expects
- output test results as strings
- benchmark troublesome functions
- run `go test` or `go test -bench={function to benchmark}` `.` means run all benchmark 
- split your test files into functionality or functions if the library is large (ie/ too fine grained > can't find examples) 

# Testing
- Testing a library includes writing a file called `<name>_test.go`
- Test files must be in the same package as functions to be tested
- Test functions must begin with with `Test`. Good practice is to call them `Test<Functionality>()`
    - these work like examples in rspec or like small isolated mains, however any go processes spawned during
    the test run will continue to run until the entire suite is complete
- A call to `Fatal("string")` causes a test to fail
- A call to `Errorf("string")` prints the string, causes a test to fail and continues running of the test function 
- The testing package has more like: `Fail` and `FailNow`
- Any number of helper functions can be made but will not be run as tests

# Benchmarking
- Benchmark files end in `_test.go` (ie/ bench in the same file you test the functionality in)
- Benchmarking functions are named `Benchmark<Functionality>`
- COME BACK TO BENCHMARKING AFTER DB HAS BEEN COMPLETED