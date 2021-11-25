package excel

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"testing"

	"github.com/xuri/excelize/v2"
)

func Test_Prepare(t *testing.T) {
	f := excelize.NewFile()
	imgTypes := []string{"png", "jpg", "gif"}
	f.SetColWidth("Sheet1", "B", "B", 50)
	for i, tt := range imgTypes {
		r := i + 1
		f.SetRowHeight("Sheet1", r, 250)

		c, _ := excelize.CoordinatesToCellName(1, r)
		f.SetCellValue("Sheet1", c, fmt.Sprintf("%s image", tt))

		c, _ = excelize.CoordinatesToCellName(2, r)
		if err := f.AddPicture("Sheet1", c, fmt.Sprintf("../testdata/static/gopher.%s", tt), ""); err != nil {
			t.Error(err)
		}
	}
	f.SaveAs("../testdata/gophers.xlsx")
}
