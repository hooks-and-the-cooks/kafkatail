package translator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Translate(t *testing.T) {
	s := []byte(`{"wallet_id":"abc"}`)
	m := Translate(s)
	expectedMap := map[string]string{"wallet_id": "abc"}
	assert.Equal(t, expectedMap, m)
}
