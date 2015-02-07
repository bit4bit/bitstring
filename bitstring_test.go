package bitstring

import (
	"testing"
)

func TestBitStringAt(t *testing.T) {
	bitString := BitString([]byte{0x0, 0x0F, 0x0}, 8*3)
	if bitString.At(8) != 0 {
		t.Fatal("at")
	}
	if bitString.At(12) != 1 {
		t.Fatal("at")
	}
	if bitString.At(13) != 1 {
		t.Fatal("at")
	}
	if bitString.At(14) != 1 {
		t.Fatal("at")
	}
	if bitString.At(15) != 1 {
		t.Fatal("at")
	}
	if bitString.At(16) != 0 {
		t.Fatal("at")
	}

	if bitString.By(12, 4) != 0x0F {
		t.Fatalf("atGroup %b", bitString.By(12, 4))
	}

	//10011101
	bitString = BitString([]byte{0x9D}, 8)
	//extract from 0 counting it self, so extract 100
	if bitString.By(0, 3) != 0x04 {
		t.Fatalf("atGroup got %b", bitString.By(0, 3))
	}

	if bitString.By(3, 3) != 0x07 {
		t.Fatalf("atGroup got %b", bitString.By(3, 3))
	}

	if bitString.By(5, 3) != 0x05 {
		t.Fatalf("atGroup got %b", bitString.By(5, 3))
	}
}
