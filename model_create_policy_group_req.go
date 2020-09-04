/*
 * Tanyi Bot open api
 *
 * tanyibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tanyibot-api-go
// CreatePolicyGroupReq struct for CreatePolicyGroupReq
type CreatePolicyGroupReq struct {
	// 外呼策略组名称
	Name string `json:"name"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 优先级策略，(SORT_FIRST, \"排序优先级优先\"),(LOCATION_MATCH_FIRST, \"归属地匹配优先\");
	DispatchType string `json:"dispatchType"`
	// 用户线路ID
	TenantPhoneNumberId int64 `json:"tenantPhoneNumberId"`
}
