/*
 * Tanyi Bot open api
 *
 * tanyibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tanyibot
// StaffGroupListRspAllOf struct for StaffGroupListRspAllOf
type StaffGroupListRspAllOf struct {
	Number int32 `json:"number,omitempty"`
	PageSize int32 `json:"pageSize,omitempty"`
	TotalElements int32 `json:"totalElements,omitempty"`
	Pages int32 `json:"pages,omitempty"`
	CsStaffGroupId int64 `json:"csStaffGroupId,omitempty"`
	GroupName string `json:"groupName,omitempty"`
}
