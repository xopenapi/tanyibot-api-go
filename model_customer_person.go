/*
 * Tanyi Bot open api
 *
 * tanyibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tanyibot
// CustomerPerson struct for CustomerPerson
type CustomerPerson struct {
	// 客户名称
	Name string `json:"name,omitempty"`
	// 性别
	Gender string `json:"gender,omitempty"`
	// 客户电话
	PhoneNumber string `json:"phoneNumber,omitempty"`
	// 话术中自定义的语句内容
	Properties map[string]interface{} `json:"properties,omitempty"`
}