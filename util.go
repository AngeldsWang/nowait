package nowait

import (
	"flag"
)

const (
	defaultTestFlag = "test.v"
	ginkgoTestFlag  = "ginkgo.v"
)

var testFlags = []string{
	defaultTestFlag,
	ginkgoTestFlag,
}

func runInTest() bool {
	for _, f := range testFlags {
		if flag.Lookup(f) != nil {
			return true
		}
	}

	return false
}
