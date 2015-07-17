package limiter

type MemStatus struct {
	All uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
	Self uint64 `json:"self"`
}

type DiskStatus struct {
	All uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}