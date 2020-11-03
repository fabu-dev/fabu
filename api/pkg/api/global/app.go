package global

type AppInfo struct {
	Name       string `json:"name"`
	BundleId   string `json:"bundle_id"`
	Version    string `json:"version"`
	Build      string `json:"build"`
	Icon       string `json:"icon"`
	Size       int64  `json:"size"`
	Identifier string `json:"identifier"`
}
