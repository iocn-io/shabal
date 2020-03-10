package shabal

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestShabal256(t *testing.T) {
	h := New()
	if i, err := h.Write([]byte("hello")); i != 5 || err != nil {
		t.Fatal(i, err)
	}
	v := h.Sum(nil)
	p, err := hex.DecodeString("6563a2d36f2f541e38aaa3f5375bfae8ce1dd2811cdf0993216669d48618aa9a")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.EqualFold(v, p) {
		t.Fatal("invalid bytes of shabal256")
	}
}
