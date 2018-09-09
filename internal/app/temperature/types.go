package temperature

type SensorData struct {
	Unit  string `json:"unit,omitempty"`
	Value string `json:"value,omitempty"`
}
