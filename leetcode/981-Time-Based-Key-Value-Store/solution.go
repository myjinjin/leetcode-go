package leetcode

type Value struct {
	value     string
	timestamp int
}

type TimeMap struct {
	values map[string][]Value
}

func Constructor() TimeMap {
	return TimeMap{values: make(map[string][]Value)}
}

func (m *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := m.values[key]; ok {
		m.values[key] = append(m.values[key], Value{value: value, timestamp: timestamp})
		return
	}
	m.values[key] = []Value{{value: value, timestamp: timestamp}}
}

func (m *TimeMap) Get(key string, timestamp int) string {
	if _, ok := m.values[key]; !ok {
		return ""
	}

	values := m.values[key]
	left := 0
	right := len(values) - 1

	result := ""

	for left <= right {
		mid := left + (right-left)/2
		if values[mid].timestamp <= timestamp {
			result = values[mid].value
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
