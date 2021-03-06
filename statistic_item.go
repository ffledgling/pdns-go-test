/*
 * PowerDNS Authoritative HTTP API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.13
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package pdnsapi

type StatisticItem struct {

	// set to “StatisticItem”
	Name string `json:"name,omitempty"`

	// The name of this item (e.g. ‘uptime’)
	Type_ string `json:"type,omitempty"`

	// The value of item
	Value string `json:"value,omitempty"`
}
