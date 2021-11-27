# sample-ex2sql
Sample application to convert excel database into sql database.

## Functions
- Import texts and images on a excel file
- Insert texts and images' path into a sql db (this application uses only in-memory db.)
- Show imported items on localhost. (This function is just a bonus)

## PreRequirements
This application depends on awesome "github.com/mattn/go-sqlite3". If you could not compile this application, please try to confirm below.

https://github.com/mattn/go-sqlite3#compilation


## Usage
```shell
git clone https://github.com/tenkoh/sample-ex2sql
cd sample-ex2sql
go build
./exsql      # or exsql.exe (on Windows)

# If you want to see imported data, let's access below.
# http://localhost:1323/posts
```

## License of codes
MIT

## License of images
Testdata includes images of Gopher (a character of Go). Gopher was designed by Renee French (http://reneefrench.blogspot.com/). 

The included images were made by tenkoh (https://twitter.com/tenkoh88), are under Creative Commons 4.0 Attributions license.