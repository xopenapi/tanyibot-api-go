/*
 * Tanyi Bot open api
 *
 * tanyibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tanyibot-api-go
// GetJobsRsp struct for GetJobsRsp
type GetJobsRsp struct {
	// 响应码
	Code int32 `json:"code,omitempty"`
	// 请求id
	RequestId string `json:"requestId,omitempty"`
	// 响应说明
	ResultMsg string `json:"resultMsg,omitempty"`
	ErrorStackTrace string `json:"errorStackTrace,omitempty"`
	Data JobsPage `json:"data,omitempty"`
}
