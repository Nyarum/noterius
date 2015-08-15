package pills

import (
	"crypto/md5"
	"encoding/hex"
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	md5Init := md5.Sum([]byte("password"))
	passMd5 := hex.EncodeToString(md5Init[:])
	_, err := EncryptPassword(string([]byte(passMd5)[:24]), "[8-15 17:9:15:743]")

	if err != nil {
		t.Error("Error encrypt password, err - ", err)
	}
}
