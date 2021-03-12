package misc

// Float64ToFloat32 float64を float32ポインタ に
func Float64ToFloat32(v float64) *float32 {
	f := float32(v)
	return &f
}

// Float64ToFloat32Value float64ポインタを float32 に
func Float64ToFloat32Value(v *float64) float32 {
	if v != nil {
		return float32(*v)
	}
	return 0
}

// Float32ToFloat64 float32を float64ポインタ に
func Float32ToFloat64(v float32) *float64 {
	f := float64(v)
	return &f
}

// Float32ToFloat64Value float32ポインタを float64 に
func Float32ToFloat64Value(v *float32) float64 {
	if v != nil {
		return float64(*v)
	}
	return 0
}
