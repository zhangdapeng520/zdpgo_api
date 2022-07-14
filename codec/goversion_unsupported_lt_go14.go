//go:build !go1.4
// +build !go1.4

package codec

import "errors"

// This codec package will only work for go1.4 and above.
// This is for the following reasons:
//   - go 1.4 was released in 2014
//   - go runtime is written fully in go
//   - interface only holds pointers
//   - reflect.Value is stabilized as 3 words

var errCodecSupportedOnlyFromGo14 = errors.New("codec: go 1.3 and below are not supported")

func init() {
	panic(errCodecSupportedOnlyFromGo14)
}
