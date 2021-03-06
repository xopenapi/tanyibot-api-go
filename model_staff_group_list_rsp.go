/*
 * Tanyi Bot open api
 *
 * tanyibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tanyibot
// StaffGroupListRsp struct for StaffGroupListRsp
type StaffGroupListRsp struct {
	// 响应码
	Code int32 `json:"code,omitempty"`
	// 请求id
	RequestId string `json:"requestId,omitempty"`
	// 响应说明
	ResultMsg string `json:"resultMsg,omitempty"`
	ErrorStackTrace string `json:"errorStackTrace,omitempty"`
	Number int32 `json:"number,omitempty"`
	PageSize int32 `json:"pageSize,omitempty"`
	TotalElements int32 `json:"totalElements,omitempty"`
	Pages int32 `json:"pages,omitempty"`
	CsStaffGroupId int64 `json:"csStaffGroupId,omitempty"`
	GroupName string `json:"groupName,omitempty"`
}
