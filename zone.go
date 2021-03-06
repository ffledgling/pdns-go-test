/*
 * PowerDNS Authoritative HTTP API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.13
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package pdnsapi

// This represents an authoritative DNS Zone.
type Zone struct {

	// Opaque zone id (string), assigned by the server, should not be interpreted by the application. Guaranteed to be safe for embedding in URLs.
	Id string `json:"id,omitempty"`

	// Name of the zone (e.g. “example.com.”) MUST have a trailing dot
	Name string `json:"name,omitempty"`

	// Set to “Zone”
	Type_ string `json:"type,omitempty"`

	// API endpoint for this zone
	Url string `json:"url,omitempty"`

	// Zone kind, one of “Native”, “Master”, “Slave”
	Kind string `json:"kind,omitempty"`

	// RRSets in this zone
	Rrsets []RrSet `json:"rrsets,omitempty"`

	// The SOA serial number
	Serial int32 `json:"serial,omitempty"`

	// The SOA serial notifications have been sent out for
	NotifiedSerial int32 `json:"notified_serial,omitempty"`

	//  List of IP addresses configured as a master for this zone (“Slave” type zones only)
	Masters []string `json:"masters,omitempty"`

	// Whether or not this zone is DNSSEC signed (inferred from presigned being true XOR presence of at least one cryptokey with active being true)
	Dnssec bool `json:"dnssec,omitempty"`

	// The NSEC3PARAM record
	Nsec3param string `json:"nsec3param,omitempty"`

	// Whether or not the zone uses NSEC3 narrow
	Nsec3narrow bool `json:"nsec3narrow,omitempty"`

	// Whether or not the zone is pre-signed
	Presigned bool `json:"presigned,omitempty"`

	// The SOA-EDIT metadata item
	SoaEdit string `json:"soa_edit,omitempty"`

	// The SOA-EDIT-API metadata item
	SoaEditApi string `json:"soa_edit_api,omitempty"`

	//  Whether or not the zone will be rectified on data changes via the API
	ApiRectify bool `json:"api_rectify,omitempty"`

	// MAY contain a BIND-style zone file when creating a zone
	Zone string `json:"zone,omitempty"`

	// MAY be set. Its value is defined by local policy
	Account string `json:"account,omitempty"`

	// MAY be sent in client bodies during creation, and MUST NOT be sent by the server. Simple list of strings of nameserver names, including the trailing dot. Not required for slave zones.
	Nameservers []string `json:"nameservers,omitempty"`
}
