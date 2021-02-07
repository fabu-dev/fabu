package global

type AppInfo struct {
	Name       string `json:"name"`
	BundleId   string `json:"bundle_id"`
	Version    string `json:"version"`
	Build      string `json:"build"`
	Platform   uint8  `json:"platform"`
	Env        uint8  `json:"env"`
	Icon       string `json:"icon"`
	Size       uint64 `json:"size"`
	Path       string `json:"path"`
	Identifier string `json:"identifier"`
	ShortKey   string `json:"short_key"`
}
