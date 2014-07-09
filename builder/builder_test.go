package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	expect := "# Title\n## String\n- Item1\n- Item2\n"
	builder := NewTextBuilder()
	result := builder.MakeTitle("Title").
		MakeContent("String").
		MakeItems([]string{"Item1","Item2",}).
		Build().
		GetString()

	if result != expect {
		t.Errorf("Expect result to %s, but %s", result, expect)
	}
}
