/*
 * Tanyi Bot open api
 *
 * tanyibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tanyibot-api-go
// CopyDialogFlowRspAllOf struct for CopyDialogFlowRspAllOf
type CopyDialogFlowRspAllOf struct {
	// 话术id
	DialogFlowId int64 `json:"dialogFlowId,omitempty"`
	// 话术名称
	Name string `json:"name,omitempty"`
	// 意向分组id
	IntentLevelTagId int64 `json:"intentLevelTagId,omitempty"`
	// 话术状态(DRAFT 草稿 PENDING_APPROVAL 等待审核 REJECTED 拒绝 APPROVED 审核通过)
	Status string `json:"status,omitempty"`
	// 话术创建时间
	CreateTime string `json:"createTime,omitempty"`
	Industry Industry `json:"industry,omitempty"`
	SubIndustry Industry `json:"subIndustry,omitempty"`
	// 描述
	Description string `json:"description,omitempty"`
	// 话术类型(NORMAL 客户定制 TEMPLATE 模板)
	Type string `json:"type,omitempty"`
}
