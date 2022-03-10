package main

import "testing"

func TestUnpack(t *testing.T) {

	tests := map[int]struct {
		packedStr string
		want      string
	}{
		0: {
			packedStr: "",
			want:      "",
		},
		1: {
			packedStr: "a4bc2d5e",
			want:      "aaaabccddddde",
		},
		2: {
			packedStr: "abcd",
			want:      "abcd",
		},
		3: {
			packedStr: "qwe\\4\\5",
			want:      "qwe45",
		},
		4: {
			packedStr: "qwe\\45",
			want:      "qwe44444",
		},
		5: {
			packedStr: "qwe\\\\5",
			want:      "qwe\\\\\\\\\\",
		},
	}

	for index, test := range tests {
		if Unpack(test.packedStr) != test.want {
			t.Errorf("test id: %d | Unpack(%s) != %s", index, test.packedStr, test.want)
		}
	}
}
