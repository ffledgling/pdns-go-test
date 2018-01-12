/*
 * PowerDNS Authoritative HTTP API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.13
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// A comment about an RRSet.
type Comment struct {

	// The actual comment
	Content string `json:"content,omitempty"`

	// Name of an account that added the comment
	Account string `json:"account,omitempty"`

	// Timestamp of the last change to the comment
	ModifidedAt int32 `json:"modifided_at,omitempty"`
}
