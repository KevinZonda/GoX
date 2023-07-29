package mapx

import "strconv"

func GetString(m map[string]any, key string, ifFailed string) string {
	if v, ok := m[key]; ok {
		if str, ok := v.(string); ok {
			return str
		}
		return ifFailed
	}
	return ifFailed
}

func GetBool(m map[string]any, key string, ifFailed bool) bool {
	if v, ok := m[key]; ok {
		switch v.(type) {
		case bool:
			return v.(bool)
		case string:
			if v.(string) == "true" {
				return true
			}
			return false
		default:
			return ifFailed
		}
	}
	return ifFailed
}

func GetInt(m map[string]any, key string, ifFailed int) int {
	if v, ok := m[key]; ok {
		switch v.(type) {
		case int:
			return v.(int)
		case string:
			if i, err := strconv.Atoi(v.(string)); err == nil {
				return i
			}
			return ifFailed
		default:
			return ifFailed
		}
	}
	return ifFailed
}

func GetInt64(m map[string]any, key string, ifFailed int64) int64 {
	if v, ok := m[key]; ok {
		switch v.(type) {
		case int64:
			return v.(int64)
		case string:
			if i, err := strconv.ParseInt(v.(string), 10, 64); err == nil {
				return i
			}
			return ifFailed
		default:
			return ifFailed
		}
	}
	return ifFailed
}

func GetFloat64(m map[string]any, key string, ifFailed float64) float64 {
	if v, ok := m[key]; ok {
		switch v.(type) {
		case float64:
			return v.(float64)
		case string:
			if i, err := strconv.ParseFloat(v.(string), 64); err == nil {
				return i
			}
			return ifFailed
		default:
			return ifFailed
		}
	}
	return ifFailed
}
