package translator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_translate(t *testing.T) {
	s := []byte(`{"wallet_id":"abc"}`)
	m := translate(s)
	expectedMap := map[string]string{"wallet_id": "abc"}
	assert.Equal(t, expectedMap, m)
}
