package consumer

import (
	actcli "hcm/cmd/task-server/logics/action/cli"
	"hcm/pkg/api/core"
	"hcm/pkg/dal/dao/tools"
	"hcm/pkg/kit"
)

// listTenants 获取所有租户ID
func listTenants(kt *kit.Kit) ([]string, error) {
	result, err := actcli.GetDataService().Global.Tenant.List(kt, &core.ListReq{
		Page:   core.NewDefaultBasePage(), //page为空支持无分页限制吗
		Fields: []string{"tenant_id"},
		Filter: tools.AllExpression(),
	})
	if err != nil {
		return nil, err
	}

	tenantIDs := make([]string, 0)
	for _, t := range result.Details {
		tenantIDs = append(tenantIDs, t.TenantID)
	}

	return tenantIDs, nil
}
