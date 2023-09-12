package utils_test

import (
	"testing"

	"github.com/cnrancher/tcr-access-control/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func Test_EncryptAES(t *testing.T) {
	e, err := utils.EncryptAES(utils.AesEncryptKey, "example")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%v", e)
}

func Test_DecryptAES(t *testing.T) {
	e, err := utils.EncryptAES(utils.AesEncryptKey, "example")
	if err != nil {
		t.Error(err)
		return
	}
	d, err := utils.DecryptAES(utils.AesEncryptKey, e)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, "example", d)
}
