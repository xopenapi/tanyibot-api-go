/*
 * Tanyi Bot open api
 *
 * tanyibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tanyibot-api-go
// PolicyGroup struct for PolicyGroup
type PolicyGroup struct {
	CreateUserId int32 `json:"createUserId,omitempty"`
	UpdateUserId int32 `json:"updateUserId,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
	// 外呼策略组ID
	CallPolicyGroupId int32 `json:"callPolicyGroupId,omitempty"`
	// 公司ID
	TenantId int64 `json:"tenantId,omitempty"`
	Name string `json:"name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 外呼策略组类型
	PhoneType string `json:"phoneType,omitempty"`
	// 优先级策略，(SORT_FIRST, \"排序优先级优先\"),(LOCATION_MATCH_FIRST, \"归属地匹配优先\")
	DispatchType string `json:"dispatchType,omitempty"`
	DetailList []CallPolicyGroupDetail `json:"detailList,omitempty"`
}
