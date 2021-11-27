package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/labstack/echo/v4"
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

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func showPosts(c echo.Context) error {
	client, err := db.Connect()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer client.Close()
	posts, err := client.Post.Query().All(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.Render(http.StatusOK, "posts", posts)
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
		imgPath, err := r.SaveImage(i, 2, imgDir, nil)
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

	// show data
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	e.Renderer = t
	e.Static("/testdata", "testdata")
	e.GET("/posts", showPosts)
	e.Logger.Fatal(e.Start(":1323"))
}
