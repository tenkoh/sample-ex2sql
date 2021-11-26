package excel

import (
	"fmt"
	"os"
	"testing"
)

func Test_DefaultSaveName(t *testing.T) {
	data := []byte("These pretzels are making me thirsty.")
	n := defaultSaveName(data)
	fmt.Println(n)
}

func Test_SaveImage(t *testing.T) {
	saveDir := "../testdata/tmp"
	if _, err := os.Stat(saveDir); err != nil {
		if err := os.Mkdir(saveDir, 0755); err != nil {
			panic(err)
		}
	}

	r, _ := NewReader("../testdata/gophers.xlsx", "Sheet1")
	for _, i := range []int{1, 2, 3} {
		n, err := r.SaveImage(i, 2, "../testdata/tmp", nil)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(n)
	}
}
