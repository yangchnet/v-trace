package constants

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_role(t *testing.T) {
	require.Equal(t, "admin", ShouldRole(
		[]string{"producer", "normal", "admin"},
	))

	require.Equal(t, "producer", ShouldRole(
		[]string{"producer", "normal", "examiner"},
	))

	require.Equal(t, "examiner", ShouldRole(
		[]string{"normal", "examiner"},
	))

	require.Equal(t, "transporter", ShouldRole(
		[]string{"normal", "examiner", "transporter"},
	))

	require.Equal(t, "producer", ShouldRole(
		[]string{"normal", "transporter", "producer"},
	))

	require.Equal(t, "boss", ShouldRole(
		[]string{"normal", "transporter", "producer", "boss"},
	))
}
