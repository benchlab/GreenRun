package rocketeerset

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"runtime/debug"
	"strings"

	"github.com/benchlab/benchcore/types"
)

// captureAnyPanic is present for use because we are dealing
// with randomized input from the greenrunner, the code in
// "benchchain/types".RocketeerSet panics on encountering
// any decoding/encoding errors. Plainly letting the greenrunner
// panic on such cases wouldn't produce new information
// and produces really noisy output.
func captureAnyPanic(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case error:
				err = e
			case string:
				err = errors.New(e)
			default:
				err = errors.New(string(debug.Stack()))
			}
		}
	}()
	fn()
	return err
}

func decodingError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Panicked on a Crisis")
}

func GreenRuns(data []byte) int {
	vs := new(types.RocketeerSet)
	vs1 := *vs
	if err := captureAnyPanic(func() { vs1.FromBytes(data) }); err != nil {
		if decodingError(err) {
			// Nothing interesting to do here
			return 0
		}
		// We've found a potential crasher
		panic(err)
	}
	if reflect.DeepEqual(&vs1, vs) {
		// Nothing interesting to do here
		return -1
	}

	asData := vs1.ToBytes()
	if !bytes.Equal(asData, data) {
		// We've found an interesting case where the input/serializing
		// data and finally serialized data aren't the same.
		panic("input and out are different")
	}

	// Now redo the deserialization
	vs2 := *vs
	vs2.FromBytes(asData)
	if !reflect.DeepEqual(vs2, vs1) {
		// We've found an interesting case where the
		// re-deserialized and last deserialized
		// data don't match.
		panic(fmt.Sprintf("input %#v and output %#v are different", vs1, vs2))
	}

	// Valid input was passed, interesting case, move on.
	return 1
}
