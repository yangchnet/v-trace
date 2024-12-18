package passwd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	passwd := "producerUser"
	crypto_passwd, err := HashPassword(passwd)
	require.NoError(t, err)
	fmt.Println(crypto_passwd)

	err = CheckPassword(passwd, crypto_passwd)
	require.NoError(t, err)
}
