package main

import (
	"context"
	"log"
	"os"

	"github.com/tenkoh/exsql/db"
	"github.com/tenkoh/exsql/excel"
)

const imgDir = "testdata/tmp"

func init() {
	if _, err := os.Stat(imgDir); err != nil {
		if err := os.Mkdir(imgDir, 0755); err != nil {
			panic(err)
		}
	}
}

func main() {
	// open file
	doc := "testdata/gophers.xlsx"
	r, err := excel.NewReader(doc, "Sheet1")
	if err != nil {
		panic(err)
	}

	client, err := db.Initialize()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	// save document data to db
	d, err := client.Document.Create().SetFilename(doc).Save(ctx)
	if err != nil {
		log.Printf("could not save document data to db; %s\n", err)
	}

	// operate each row
	for i := 1; i < 4; i++ {
		title, err := r.GetValue(i, 1)
		if err != nil {
			continue
		}
		imgPath, err := r.SaveImage(i, 1, imgDir, nil)
		if err != nil {
			continue
		}
		_, err = client.Post.Create().
			SetTitle(title).
			SetImgPath(imgPath).
			AddDocument(d).
			Save(ctx)
		if err != nil {
			continue
		}
	}
}
