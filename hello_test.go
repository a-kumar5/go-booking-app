package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items:  []string{"Ayush"},
			result: "Hello, Ayush!",
		},
		{
			items:  []string{"Ayush", "Aish"},
			result: "Hello, Ayush, Aish!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("Wanted %s (%v), got %s", st.result, st.items, s)
		}
	}
}
