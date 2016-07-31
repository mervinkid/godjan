# Godjan

![Go](https://img.shields.io/badge/Go-1.6.3-blue.svg?style=flat)
![Build](https://img.shields.io/badge/Build-Passing-brightgreen.svg?style=flat)
![License](https://img.shields.io/badge/License-MIT-lightgray.svg?style=flat)
![Release](https://img.shields.io/badge/Release-1.0.0-blue.svg?style=flat)

```
   ____           _  _             
  / ___| ___   __| |(_) __ _ _ __
 | |  _ / _ \ / _` || |/ _` | '_ \
 | |_| | (_) | (_| || | (_| | | | |
  \____|\___/ \__,_|/ |\__,_|_| |_|
                  |__/
```

**Godjan** implement the build in password crypto algorithm of **django** framework with google **go** language.

Django is a popular web framework for python with build in orm and lot of features.<br> 
The build in password crypto algorithm which django have is awesome. <br>
Now, I implement it with go language.<br> 

## Installation

Begin by installing `godjan` using `go get` command.

```
$ go get github.com/mofei2816/godjan
```

If you already have `godjan` installed, updating `godjan` is simple:

```
$ go get -u github.com/mofei2816/godjan
```

## Usage

### Import

Import to your go file.

```
import "github.com/mofei2816/godjan"
```

### Algorithms

There are 4 algorithms supported by `godjan`.

- AlgorithmPBKDF2SHA1
- AlgorithmPBKDF2SHA256
- AlgorithmPBKDF2SHA512
- AlgorithmPBKDF2MD5

### Make password

Make password with default algorithm and random salt by using function `MakePassword`.

```
import (
    "github.com/mofei2816/godjan"
)

func main() {
    password := "mypassword"
    encoded := godjan.MakePassword(password)
    // the var encoded is the encrypted password
}
```

Make password with specified algorithm, salt an iter by using function `MakePasswordWithSaltIterAndAlgorithm`.

```
import (
    "github.com/mofei2816/godjan"
)

func main() {
    password := "mypassword"
    salt := "mysalt"
    iter := 1000
    encoded := godjan.MakePasswordWithSaltIterAndAlgorithm(
        password, 
        salt, 
        iter, 
        godjan.AlgorithmPBKDF2SHA512,
    )
    // the var encoded is the encrypted password
}
```

### Check Password

Check password by using function `CheckPassword`

```
import (
    "github.com/mofei2816/godjan"
)

func main() {
    pain := "painpassword"
    encoded := "encodedpassword"
    if godjan.CheckPassword(pain, encoded) {
        // password is valid
    }
}
```

## Dependencies

```
golang.org/x/crypto/pbkdf2
```

## Contributing

1. Fork it.
2. Create your feature branch. (`$ git checkout feature/my-feature-branch`)
3. Commit your changes. (`$ git commit -am 'What feature I just added.'`)
4. Push to the branch. (`$ git push origin feature/my-feature-branch`)
5. Create a new Pull Request

## Authors

[@Mervin](https://github.com/mofei2816) 

## License

The MIT License (MIT). For detail see [LICENSE](LICENSE).



