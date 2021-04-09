# passwordHash
[![GoDoc](https://godoc.org/github.com/malixsys/passwordHash?status.svg)](https://godoc.org/github.com/malixsys/passwordHash)
[![Build Status](https://travis-ci.org/malixsys/passwordHash.svg?branch=master)](https://travis-ci.org/malixsys/passwordHash)
[![Go Report Card](https://goreportcard.com/badge/github.com/malixsys/passwordHash)](https://goreportcard.com/report/github.com/malixsys/passwordHash)

Password hashing package compatible with Node's `password-hash`


## Usage

```go
package main
import "github.com/malixsys/passwordHash"

 func main() {
	 // Generate a hashed password
	 testPassword := `test1234`
	 hashedPassword := passwordHash.Generate(testPassword, nil)

	 // Test correct password in constant time
	 valid := passwordHash.Verify(hashedPassword, testPassword)
	 log.Printf("The password validity is %t against the hash", valid)

	 // Test incorrect password in constant time
	 valid, err = passwordHash.Verify(hashedPassword, "badPass")
	 log.Printf("The password validity is %t against the hash", valid)
 }

```

## Development

### Tests

Some tests are included

To run, use:

```bash
go test
```

## TODO

- Extend to other algorithms 
- More development info
- More tests

## Additional info

- Used in Node project: [password-hash](https://github.com/davidwood/node-password-hash)
- Package layout inspiration: [argon2pw](https://github.com/raja/argon2pw)
