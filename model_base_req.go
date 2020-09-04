/*
 * Tanyi Bot open api
 *
 * tanyibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tanyibot
// BaseReq 
type BaseReq struct {
	AppKey string `json:"appKey,omitempty"`
	AppSecret string `json:"appSecret,omitempty"`
	TenantSign string `json:"tenantSign,omitempty"`
	Version string `json:"version,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Signature string `json:"signature,omitempty"`
}
