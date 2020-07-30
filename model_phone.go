/*
 * Tanyi Bot open api
 *
 * tanyibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package aibot
// Phone Phone
type Phone struct {
	// 电话id
	TenantPhoneNumberId int64 `json:"tenantPhoneNumberId,omitempty"`
	// 电话号码
	PhoneNumber string `json:"phoneNumber,omitempty"`
	// 电话号码名称
	PhoneName string `json:"phoneName,omitempty"`
	// (MOBILE, \"手机\"),(LANDLINE, \"固话\"),(UNFIXED_CALL, \"无主叫固话\")
	PhoneType string `json:"phoneType,omitempty"`
	// 本地话费（单位：厘）
	LocalBillRate string `json:"localBillRate,omitempty"`
	// 外地话费（单位：厘）
	OtherBillRate string `json:"otherBillRate,omitempty"`
	// (FINANCE, \"金融\"),(OTHER, \"其他\")
	CallOutIndustry string `json:"callOutIndustry,omitempty"`
	// 归属地区号
	AreaCode string `json:"areaCode,omitempty"`
	// 归属地省
	Province string `json:"province,omitempty"`
	// 归属地市
	City string `json:"city,omitempty"`
	// 无法外呼地区
	DeadArea string `json:"deadArea,omitempty"`
}
