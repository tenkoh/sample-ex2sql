package excel

import (
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tenkoh/exsql/util"
	"github.com/xuri/excelize/v2"
)

var ErrInvalidSheetName = errors.New("the specified sheet is not in the file")

type Reader struct {
	file  *excelize.File
	sheet string
}

func NewReader(path, sheet string) (*Reader, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}
	sheets := f.GetSheetList()
	for _, s := range sheets {
		if sheet == s {
			return &Reader{f, sheet}, nil
		}
	}
	return nil, ErrInvalidSheetName
}

// GetValue returns the cell value. The origin of the input row and col is (1,1).
func (r *Reader) GetValue(row, col int) (string, error) {
	cell, err := excelize.CoordinatesToCellName(col, row)
	if err != nil {
		return "", fmt.Errorf("error at (%d, %d); %w", row, col, err)
	}
	v, err := r.file.GetCellValue(r.sheet, cell)
	if err != nil {
		return "", fmt.Errorf("error at (%d, %d); %w", row, col, err)
	}
	return v, nil
}

// SaveImage saves a inline image, then returns the saved file's name.
// Set the func if you want to define save-name. Default is yyyymmdd-hhmmdd_hash(md5).ext.
// Set the func as nil, if you do not define save-name.
func (r *Reader) SaveImage(row, col int, savedir string, fn func([]byte) string) (string, error) {
	cell, err := excelize.CoordinatesToCellName(col, row)
	if err != nil {
		return "", fmt.Errorf("error at (%d, %d); %w", row, col, err)
	}
	t, i, err := r.file.GetPicture(r.sheet, cell)
	if err != nil {
		return "", fmt.Errorf("error at (%d, %d); %w", row, col, err)
	}

	saveBase := defaultSaveName(i)
	if fn != nil {
		saveBase = fn(i)
	}
	saveExt := filepath.Base(t)
	saveName := filepath.Join(savedir, saveBase+saveExt)
	f, err := os.Create(saveName)
	if err != nil {
		return "", fmt.Errorf("error at (%d, %d); %w", row, col, err)
	}
	defer f.Close()
	if _, err := f.Write(i); err != nil {
		return "", fmt.Errorf("error at (%d, %d); %w", row, col, err)
	}
	return saveName, nil
}

func defaultSaveName(b []byte) string {
	n := util.LocalNow().Format("20060102-150405")
	return fmt.Sprintf("%s_%x", n, md5.Sum(b))
}
