package nowait

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNowait(t *testing.T) {
	wg := NewNowaitGroup()
	wg.Add(2)

	count := 1
	go func() {
		defer wg.Done()
		count++
	}()

	go func() {
		defer wg.Done()
		count++
	}()

	wg.Wait()

	assert.Equal(t, count, 3)
}
