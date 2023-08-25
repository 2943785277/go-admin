package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type User struct {
	Token string
}

var loggedInUsers map[int]*User

func main() {
	loggedInUsers = make(map[int]*User)
	loggedInUsers[32] = &User{
		Token: "asdas",
	}
	loggedInUsers[3] = &User{
		Token: "666",
	}
	if _, ok := loggedInUsers[3]; ok {
		fmt.Println("User already logged in")
	} else {
		fmt.Println("66")
	}

	return
	f := excelize.NewFile()
	// Create a new sheet.
	Time := time.Now().Unix()
	Fname := strconv.FormatInt(Time, 10) + "Book.xlsx"
	index := f.NewSheet("Sheet1")
	n := 1
	for n < 10 {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(n), "Hello world."+strconv.Itoa(n))
		n++
	}
	// name := "A" + string(n)
	// fmt.Println(n)
	// Set value of a cell.
	// f.SetCellValue("Sheet2", "A2", "Hello world.")
	// f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs(Fname); err != nil {
		fmt.Println(err)
	}

}
