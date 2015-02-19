package authboss

// HTMLData is used to render templates with.
type HTMLData map[string]interface{}

// NewHTMLData creates HTMLData from key-value pairs. The input is a key-value
// slice, where odd elements are keys, and the following even element is their value.
func NewHTMLData(data ...interface{}) HTMLData {
	if len(data)%2 != 0 {
		panic("It should be a key value list of arguments.")
	}

	h := make(HTMLData)

	for i := 0; i < len(data)-1; i += 2 {
		k, ok := data[i].(string)
		if !ok {
			panic("Keys must be strings.")
		}

		h[k] = data[i+1]
	}

	return h
}

// Merge adds extra key-values to the HTMLData. The input is a key-value
// slice, where odd elements are keys, and the following even element is their value.
func (h HTMLData) Merge(data ...interface{}) HTMLData {
	if len(data)%2 != 0 {
		panic("It should be a key value list of arguments.")
	}

	for i := 0; i < len(data)-1; i += 2 {
		k, ok := data[i].(string)
		if !ok {
			panic("Keys must be strings.")
		}
	}

	return h
}
