package encodingx

import "testing"

func TestFromJsonTo(t *testing.T) {
	type Test struct {
		Hello string
	}
	var test Test
	err := FromJsonTo([]byte(`{"Hello": "World"}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.Hello != "World" {
		t.Fatal("Expected 'World', got", test.Hello)
	}
}

func TestToJson(t *testing.T) {
	type Test struct {
		Hello string
	}
	test := Test{Hello: "World"}
	s, err := ToJson(test)
	if err != nil {
		t.Fatal(err)
	}
	if s != `{"Hello":"World"}` {
		t.Fatal("Expected '{\"Hello\":\"World\"}', got", s)
	}
}
