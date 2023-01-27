package service

func setStringFromMap(m map[string]interface{}, key string, field *string) {
	if value, ok := m[key]; ok {
		if v, ok := value.(string); ok {
			*field = v
		}
	}
}

func setIntFromMap(m map[string]interface{}, key string, field *int) {
	if value, ok := m[key]; ok {
		if v, ok := value.(int); ok {
			*field = v
		}
	}
}
