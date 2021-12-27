package graceful

import (
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New())
}

func TestWait(t *testing.T) {
	sigChan := New()
	go func() {
		time.Sleep(200 * time.Millisecond)
		sigChan.channel <- syscall.SIGINT
	}()

	sigChan.Wait()
}
