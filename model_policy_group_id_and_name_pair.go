/*
 * Tanyi Bot open api
 *
 * tanyibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tanyibot-api-go
// PolicyGroupIdAndNamePair struct for PolicyGroupIdAndNamePair
type PolicyGroupIdAndNamePair struct {
	// 外呼策略组id
	CallPolicyGroupId string `json:"callPolicyGroupId,omitempty"`
	// 公司Id
	TenantId int64 `json:"tenantId,omitempty"`
	// 外呼策略组名称
	Name string `json:"name,omitempty"`
}
