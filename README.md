# Custom Testing Framework for Go

This is a lightweight and flexible testing framework for Go, built on top of the standard `testing` package. It adds features such as enhanced assertions, simple mocking, fixtures for setup and teardown, asynchronous testing utilities, and logging support—all without external dependencies.

## Features

- **Enhanced Assertions:** Simplified methods for condition checks.
- **Mocking:** Easy-to-implement mock functions.
- **Fixtures:** Setup and cleanup for reusable test environments.
- **Asynchronous Testing:** Tools to wait for conditions with timeouts.
- **Logging:** Debugging support during test execution.
- **Nil and Panic Assertions:** Utilities to handle `nil` checks and panic validations.

---

## Installation

Clone or copy the `testframework` folder into your project directory:

```bash
cp -r testframework/ <your_project_directory>/
```
## Directory Structure
```
testframework/
├── assert.go       # Assertion functions
├── test.go         # Base test object and common methods
├── mock.go         # Mocking implementation
├── fixtures.go     # Setup and teardown utilities
├── async.go        # Asynchronous testing tools
├── log.go          # Logging utilities for test output
```

## Usage
Importing the Framework <br>
To use the framework, import it in your test files:
```
import (
    "your_project/testframework"
    "testing"
)
```

## Creating a Test Object
Wrap your testing.T object with the Test object provided by the framework:

```
func TestExample(t *testing.T) {
    test := testframework.New(t)
    // Use test methods for assertions, logging, etc.
}
```
