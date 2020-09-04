/*
 * Tanyi Bot open api
 *
 * tanyibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tanyibot-api-go
// IntentLevelTag struct for IntentLevelTag
type IntentLevelTag struct {
	// 意向标签分组Id
	IntentLevelTagId int32 `json:"intentLevelTagId,omitempty"`
	// 意向标签分组名称
	Name string `json:"name,omitempty"`
	Details []IntentLevelTagDetails `json:"details,omitempty"`
}
