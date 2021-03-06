/*
 * Tanyi Bot open api
 *
 * tanyibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tanyibot
// DialogFlowCallJobRelatedInfoAllOf struct for DialogFlowCallJobRelatedInfoAllOf
type DialogFlowCallJobRelatedInfoAllOf struct {
	// 是否有变量名
	PlaceholderExist bool `json:"placeholderExist,omitempty"`
	// 是否转人工标识
	JumpToHumanServiceExist bool `json:"jumpToHumanServiceExist,omitempty"`
	// 是否有人工介入标识
	HumanInterventionExist bool `json:"humanInterventionExist,omitempty"`
}
