package main

import (
	"errors"
	"testing"
)

func TestExtractParams(t *testing.T) {
	tt := []struct {
		desc      string
		path      string
		typeParam string
		id        int
		err       error
	}{
		{
			desc:      "success",
			path:      "/serve/footype/123",
			typeParam: "footype",
			id:        123,
		},
		{
			desc: "bad path",
			path: "/serve/123/footype",
			err:  errors.New("bad path: /serve/123/footype"),
		},
		{
			desc: "out-of-range ID",
			path: "/serve/footype/123456789012345678901234567890",
			err:  errors.New(`bad ID: strconv.ParseInt: parsing "123456789012345678901234567890": value out of range`),
		},
	}
	for _, tc := range tt {
		errorf := makeErrorf(t, tc.desc)
		id, typeParam, err := extractParams(tc.path)
		if err != tc.err {
			if tc.err == nil {
				errorf("expected no error, but got %v", err)
				continue
			}
			if expErrstr, actualErrstr := tc.err.Error(), err.Error(); expErrstr != actualErrstr {
				errorf("expected error %q but got %q", tc.err, err)
				continue
			}
			continue
		}
		if tc.id != id {
			errorf("expected ID %d, but got %d", tc.id, id)
		}
		if tc.typeParam != typeParam {
			errorf("expected Type Param %q, but got %q", tc.typeParam, typeParam)
		}
	}
}

func BenchmarkExtractParams(b *testing.B) {
	var (
		id        int
		typeParam string
		err       error
	)
	for i := 0; i < b.N; i++ {
		id, typeParam, err = extractParams("/serve/foofoofoo/12345")
	}
	b.Logf("benchmark finished: id=%d, type=%s, err=%v", id, typeParam, err)
}

// makeErrorf adds a description parameter as the first output of Errorf.  This
// is useful for table driven tests that include a description property in the
// test table, to help clarify test output, while not having to remember to
// pass the description as the first substitution on every call to t.Errorf.
func makeErrorf(t *testing.T, desc string) func(string, ...interface{}) {
	return func(format string, args ...interface{}) {
		fargs := make([]interface{}, 0, len(args)+1)
		fargs = append(fargs, desc)
		fargs = append(fargs, args...)
		t.Errorf("[%s]: "+format, fargs...)
	}
}
