package dic

import "testing"

func TestType(t *testing.T) {
	var _ IDictionary[int, string] = NewDictionary[int, string]()
	var _ IDictionary[int, string] = NewConcurrentDictionary[int, string]()
}
