package models

import "time"

type LFSOp string

const (
	LFSOpUpload   LFSOp = "upload"
	LFSOpDownload LFSOp = "download"
)

type LFSBatchRequest struct {
	Operation LFSOp    `json:"operation" validate:"oneof=upload download"`
	Transfers []string `json:"transfers" validate:"unique"`
	Ref       *struct {
		Name string `json:"name"`
	} `json:"ref"`
	Objects []struct {
		OID  string `json:"oid"`
		Size int    `json:"size"`
	} `json:"objects"`
	HashAlgo string `json:"hash_algo" validate:"omitempty,eq=sha256"`
}

type LFSBatchResponse struct {
	Transfer string                   `json:"transfer"`
	Objects  []LFSBatchObjectResponse `json:"objects"`
	HashAlgo string                   `json:"hash_algo"`
}

type LFSBatchErrorResponse struct {
	Message          string  `json:"message"`
	DocumentationUrl *string `json:"documentation_url"`
	RequestId        *string `json:"request_id"`
}

type LFSBatchAction struct {
	Href      string            `json:"href"`
	Header    map[string]string `json:"header"`
	ExpiresAt time.Time         `json:"expires_at"`
}

type LFSBatchObjectResponse struct {
	OID           string `json:"oid"`
	Size          int    `json:"size"`
	Authenticated *bool  `json:"authenticated"`
	Actions       *struct {
		Download *LFSBatchAction `json:"download"`
		Upload   *LFSBatchAction `json:"upload"`
		Verify   *LFSBatchAction `json:"verify"`
	} `json:"actions"`
	Error *LFSBatchObjectError `json:"error"`
}

type LFSBatchObjectError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
