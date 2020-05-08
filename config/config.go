package config

var (
	Port       = ":8087"
	VersionApp = "0.1.1"
)

var JsonDefault = `{}`

var (
	NumCounters int64 = 1e7     // Num keys to track frequency of (10M).
	MaxCost     int64 = 1 << 30 // Maximum cost of cache (1GB).
	BufferItems int64 = 64      // Number of keys per Get buffer.
)
