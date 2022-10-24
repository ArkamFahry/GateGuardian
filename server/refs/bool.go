package refs

// NewBoolRef returns a ref to a bool with given value
func NewBoolRef(v bool) *bool {
	return &v
}

// BoolValue returns the value of the given bool ref
func BoolValue(r *bool) bool {
	if r == nil {
		return false
	}
	return *r
}
