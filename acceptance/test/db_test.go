package acceptace

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	resp, err := http.Get("http://localhost:8081/write")
	assert.Nil(t, err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, "success\n", fmt.Sprintf("%s", body))
}

func TestGet(t *testing.T) {
	resp, err := http.Get("http://localhost:8081/get")
	assert.Nil(t, err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, "success: id=1\n", fmt.Sprintf("%s", body))
}
