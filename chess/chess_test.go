package chess

import "testing"

var tests = []struct {
	white  string
	black  string
	attack bool
}{
	{"f1", "g2", false},
	{"d4", "f4", false},
	{"a8", "b6", true},
	{"b7", "d8", true},
	{"c6", "b8", true},
	{"d1", "f2", true},
	{"e3", "g2", true},
	{"f8", "h7", true},
	{"g1", "h3", true},
	{"h4", "g2", true},
}

func TestCanKnightAttack(t *testing.T) {
	for _, test := range tests {
		attack := canKnightAttack(test.white, test.black)
		if attack != test.attack {
			t.Fatalf("CanKnightAttack(%s, %s) = %t, want %t.",
				test.white, test.black, attack, test.attack)
		}
	}
}

func BenchmarkCanKnightAttack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			canKnightAttack(test.white, test.black)
		}
	}
}
