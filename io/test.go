package hello

import "testing"

func TestSay(t *testing.T) {
	subtests := []struct {
		items []string
		result string
	
} {
	{
		result: "hello world",
	},
	{
		items: []string{"matt"},
		result: "hello, matt",
	},
	{
		items: []string{"kartik", "sharma", "gautam"},
		result: "hello, kartik, sharma, gautam",
	},
}

for _,st := range subtests {
	if s := Say(st.items); s!= st.result{
		t.Errorf("got %s, gave %v, wanted %s", s,
		st.items, st.result)
	}
}
}