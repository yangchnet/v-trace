package authz

import (
	"context"

	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"github.com/casbin/casbin/v2"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewAuthzServer)

type AuthzServer interface {
	CanDo(context.Context, *AuthzRequest) bool
	BatchCanDo(context.Context, []*AuthzRequest) []bool
}

var _ AuthzServer = (*CasbinAuthz)(nil)

type CasbinAuthz struct {
	Enforcer *casbin.Enforcer
}

type AuthzRequest struct {
	Sub string
	Obj string
	Act string
}

// NewAuthzServer create authz server with given location of model and policy
// if the location of model and policy is not given, will use default model and policy.
func NewAuthzServer(ctx context.Context) *CasbinAuthz {
	var (
		modelLoc, policyLoc string
		locs                []string
	)

	if ctx.Value("casbin") != nil {
		locs = ctx.Value("casbin").([]string)
	}

	if len(locs) <= 0 {
		modelLoc = "app/vtrace/config/authz/rbac_model.conf"
		policyLoc = "app/vtrace/config/authz/policy.csv"
	} else {
		modelLoc = locs[0]
		policyLoc = locs[1]
	}

	e, err := casbin.NewEnforcer(modelLoc, policyLoc) // TODO config location
	if err != nil {
		logger.Panic("failed to create authz server: ", err)
	}

	return &CasbinAuthz{
		Enforcer: e,
	}
}

// CanDo check if AuthzRequest is valid.
func (s *CasbinAuthz) CanDo(ctx context.Context, req *AuthzRequest) bool {
	ok, err := s.Enforcer.Enforce(req.Sub, req.Obj, req.Act)
	if err != nil {
		logger.Error(err)

		return false
	}

	return ok
}

// BatchCanDo check if multiple AuthzRequests is valid.
func (s *CasbinAuthz) BatchCanDo(ctx context.Context, reqs []*AuthzRequest) []bool {
	batchreq := make([][]interface{}, 0)
	for _, req := range reqs {
		batchreq = append(batchreq, []interface{}{req.Sub, req.Obj, req.Act})
	}
	results, err := s.Enforcer.BatchEnforce(batchreq)
	if err != nil {
		logger.Error(err)

		return make([]bool, len(reqs))
	}

	return results
}
