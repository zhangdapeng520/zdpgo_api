//go:build !go1.9 || safe || codec.safe || appengine || !gc
// +build !go1.9 safe codec.safe appengine !gc

package codec

import "reflect"

// This files contains safe versions of the code where the unsafe versions are not supported
// in either gccgo or gollvm.
//
// - rvType:
//   reflect.toType is not supported in gccgo, gollvm.

func rvType(rv reflect.Value) reflect.Type {
	return rv.Type()
}
