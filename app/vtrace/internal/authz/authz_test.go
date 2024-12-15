package authz

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AuthzServer(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "casbin", []string{"../../config/authz/rbac_model.conf", "../../config/authz/policy.csv"})
	server := NewAuthzServer(ctx)

	// require.Equal(t, true, server.CanDo(ctx, &AuthzRequest{
	// 	Sub: "alice",
	// 	Obj: "data2",
	// 	Act: "read",
	// }))

	// require.Equal(t, false, server.CanDo(ctx, &AuthzRequest{
	// 	Sub: "bob",
	// 	Obj: "data2",
	// 	Act: "read",
	// }))

	require.Equal(t, true, server.CanDo(ctx, &AuthzRequest{
		Sub: "admin",
		Obj: "/api/v1/goods/class",
		Act: "create",
	}))
}
