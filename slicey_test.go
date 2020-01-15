package slicey

import "testing"

func TestSanitize(t *testing.T) {
	input := []interface{}{"ayam", "bawang", "lemonilo", "kari", "bebek"}
	got := Splitl(input, 2)

	if len(got) != 2 {
		t.Errorf("len of return should be 2")
	}

	ret := make([][]interface{}, 2)
	ret[0] = append(ret[0], []interface{}{"ayam", "bawang", "lemonilo"})
	ret[1] = append(ret[1], []interface{}{"kari", "bebek"})

	if got[0][0] != "ayam" {
		t.Errorf("return should be ayam")
	}
}
