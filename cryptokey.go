/*
 * PowerDNS Authoritative HTTP API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.13
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Describes a DNSSEC cryptographic key
type Cryptokey struct {

	// set to \"Cryptokey\"
	Type_ string `json:"type,omitempty"`

	// The internal identifier, read only
	Id string `json:"id,omitempty"`

	Keytype string `json:"keytype,omitempty"`

	// Whether or not the key is in active use
	Active bool `json:"active,omitempty"`

	// The DNSKEY record for this key
	Dnskey string `json:"dnskey,omitempty"`

	// An array of DS records for this key
	Ds []string `json:"ds,omitempty"`

	// The private key in ISC format
	Privatekey string `json:"privatekey,omitempty"`

	// The name of the algorithm of the key, should be a mnemonic
	Algorithm string `json:"algorithm,omitempty"`

	// The size of the key
	Bits int32 `json:"bits,omitempty"`
}
