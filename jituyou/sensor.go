package main

type SensorData struct {
	SensorType string
	ModelID    string
	Value      float32
}

// まず型を作っておいて、分岐をこんな感じで書くと読みやすい
func (r SensorData) ReadValue() float32 {
	if r.SensorType == "Fahrenheit" {
		return (r.Value * 9 / 5) + 32
	}
	return r.Value
}
