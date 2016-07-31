// The MIT License (MIT)
//
// Copyright (c) 2016 Mervin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package godjan

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	// AlgorithmPBKDF2SHA1 defined algorithm with PBKDF2, hmac and SHA1
	AlgorithmPBKDF2SHA1 = "pbkdf2_sha1"

	// AlgorithmPBKDF2SHA256 defined algorithm with PBKDF2, hmac and SHA256
	AlgorithmPBKDF2SHA256 = "pbkdf2_sha256"

	// AlgorithmPBKDF2SHA512 defined algorithm with PBKDF2, hmac and SHA512
	AlgorithmPBKDF2SHA512 = "pbkdf2_sha512"

	// AlgorithmPBKDF2MD5 defined algorithm with PBKDF2, hmac and MD5
	AlgorithmPBKDF2MD5 = "pbkdf2_md5"
)

const (
	// DefaultIter defined default iter value
	DefaultIter = 10000

	// DefaultSaltLength defined default salt length
	DefaultSaltLength = 12

	// DefaultAlgorithm defined default algorithm
	DefaultAlgorithm = AlgorithmPBKDF2SHA256
)

// GetAlgorithm returns support hash function with specified algorithm name.
func GetAlgorithm(algorithm string) (func() hash.Hash, error) {
	switch algorithm {
	case AlgorithmPBKDF2SHA1:
		return sha1.New, nil
	case AlgorithmPBKDF2SHA256:
		return sha256.New, nil
	case AlgorithmPBKDF2SHA512:
		return sha512.New, nil
	case AlgorithmPBKDF2MD5:
		return md5.New, nil
	}
	return nil, errors.New("specified algorithm is not support.")
}

// GetEncodedHash secure password hashing using the PBKDF2 algorithm
func GetEncodedHash(password, salt string, iter int, algorithm string) (string, error) {

	digest, err := GetAlgorithm(algorithm)

	if err != nil {
		return "", err
	}

	rawHash := pbkdf2.Key([]byte(password), []byte(salt), iter, 32, digest)
	hashBase64 := make([]byte, base64.StdEncoding.EncodedLen(len(rawHash)))
	base64.StdEncoding.Encode(hashBase64, rawHash)
	return string(hashBase64), nil
}

// MakePassword turn a plain-text password into a hash with random salt value, default iter and algorithm.
func MakePassword(src string) (string, error) {
	return MakePasswordWithSaltIterAndAlgorithm(src, GetRandomString(DefaultSaltLength), DefaultIter, DefaultAlgorithm)
}

// MakePasswordWithSaltIterAndAlgorithm turn a plain-text password into a hash with salt value, iter and algorithm.
func MakePasswordWithSaltIterAndAlgorithm(source, salt string, iter int, algorithm string) (string, error) {
	encodedHash, err := GetEncodedHash(source, salt, iter, algorithm)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s$%s$%s$%s", DefaultAlgorithm, strconv.FormatInt(int64(iter), 10), salt, encodedHash), nil
}

// CheckPassword returns a boolean of whether the raw password matches the three
// part encoded digest.
func CheckPassword(src, encoded string) bool {
	srcSlice := strings.Split(encoded, "$")

	if len(srcSlice) != 4 {
		return false
	}

	algorithm := srcSlice[0]
	iter := srcSlice[1]
	salt := srcSlice[2]
	encodedHash := srcSlice[3]

	iterInt, err := strconv.Atoi(iter)

	if err != nil {
		return false
	}

	encodedResult, err := GetEncodedHash(src, salt, iterInt, algorithm)
	if err != nil {
		return false
	}
	return encodedHash == encodedResult
}
