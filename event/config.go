package event

type ConfigEvent struct {
	Config interface{} `json:"config,omitempty"`
}
