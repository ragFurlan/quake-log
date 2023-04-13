package quakeLog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	rightURL = "https://gist.githubusercontent.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8/raw/df6ef4a9c0b326ce3760233ef24ae8bfa8e33940/qgames.log"
	wrongURL = "http://teste"
)

func TestQuakeLogTest_Success(t *testing.T) {
	quakeLog := NewQuakeLog(rightURL)
	log, err := quakeLog.Get()

	assert.NoError(t, err)
	assert.NotEmpty(t, log)
}
