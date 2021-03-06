/*
 * Tanyi Bot open api
 *
 * tanyibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tanyibot
// ImportCustomerReq struct for ImportCustomerReq
type ImportCustomerReq struct {
	// 任务Id
	RobotCallJobId int64 `json:"robotCallJobId"`
	// 是否已呼列表去重 默认false
	CallRecordDup bool `json:"callRecordDup,omitempty"`
	CustomerPersons []CustomerPerson `json:"customerPersons,omitempty"`
}
