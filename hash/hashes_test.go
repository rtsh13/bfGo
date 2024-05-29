package hash

import "testing"

func TestFnv1a(t *testing.T) {
	tests := []struct {
		input    []byte
		expected uint
	}{
		{[]byte(""), 0xcbf29ce484222325},
		{[]byte("abc"), 0x0e71fa2190541574b},
		{[]byte("message"), 0x546401b5d2a8d2a4},
		{[]byte("hello world"), 0x779a65e7023cd2e7},
	}

	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			got := Fnv1a(test.input)
			if got != test.expected {
				t.Errorf("Fnv1a(%q) = %x; want %x", test.input, got, test.expected)
			}
		})
	}
}

func TestMurmum3_128(t *testing.T) {
	tests := []struct {
		input      []byte
		expectedH1 uint
		expectedH2 uint
	}{
		{[]byte(""), 0x0000000000000000, 0x0000000000000000},
		{[]byte("a"), 0x85555565f6597889, 0xe6b53a48510e895a},
		{[]byte("message"), 0xb520427404685fa1, 0x2500955e3b9cab83},
		{[]byte("1234567890"), 0xecfa4ae68079870a, 0xc1d017c820ebd22b},
	}

	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			gotH1, gotH2 := Murmum3_128(test.input)
			if gotH1 != test.expectedH1 || gotH2 != test.expectedH2 {
				t.Errorf("Murmum3_128(%q) = (%x, %x); want (%x, %x)", test.input, gotH1, gotH2, test.expectedH1, test.expectedH2)
			}
		})
	}
}
