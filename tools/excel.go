package tools

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/tealeg/xlsx"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type outputer func(s string, n string)
type outputerIndex func(s string)

func GenerateNewCsvByIndex(filepath string, sheetIndex int, outputf outputerIndex, delimiter string) error {

	xlFile, error := xlsx.OpenFile(filepath)
	if error != nil {
		return error
	}
	sheetLen := len(xlFile.Sheets)
	switch {
	case sheetLen == 0:
		return errors.New("This XLSX file contains no sheets.")
	case sheetIndex >= sheetLen:
		return fmt.Errorf("No sheet %d available, please select a sheet between 0 and %d\n", sheetIndex, sheetLen-1)
	}
	sheet := xlFile.Sheets[sheetIndex]
	for _, row := range sheet.Rows {
		var vals []string
		if row != nil {
			for _, cell := range row.Cells {
				str, err := cell.FormattedValue()
				if err != nil {
					vals = append(vals, err.Error())
				}
				vals = append(vals, fmt.Sprintf("%q", str))
			}
			outputf(strings.Join(vals, delimiter) + "\n")
		}
	}
	return nil

}

func GenerateNewCsvAll(filepath string, delimiter string, outputf outputer, regex string) error {

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	numSheets := f.SheetCount

	fmt.Printf("Sheetcount: %s \n", strconv.Itoa(numSheets))

	for i := 1; i <= numSheets; i++ {
		sheetname := f.GetSheetName(i)

		//TODO: Was passiert hier? Warum "false &&..." ?
		if false && !validateSheetName(sheetname, regex) {
			fmt.Printf("Sheetname: %s with index %s does not match the regexp -> skipping \n", sheetname, strconv.Itoa(i))
		} else {
			fmt.Printf("Sheetname: %s with index %s matches regexp -> generating CSV file \n", sheetname, strconv.Itoa(i))

			xlFile, error := xlsx.OpenFile(filepath)
			if error != nil {
				return error
			}

			sheet := xlFile.Sheet[sheetname]
			for _, row := range sheet.Rows {
				var vals []string
				if row != nil {
					for _, cell := range row.Cells {
						str, err := cell.FormattedValue()
						if err != nil {

							vals = append(vals, err.Error())
						}

						fmt.Printf("Read Value: %s \n", str)
						vals = append(vals, fmt.Sprintf("%q", str))
					}
					outputf(strings.Join(vals, delimiter)+"\n", sheetname)
				}
			}

		}

	}

	return nil
}

func GetBlockByCoords(filepath string, findRow string, findCol string, sheet string) ([]string, error) {

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rows, err := f.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
		log.Fatal("cannot open sheet")
	}

	cols, err := f.GetCols(sheet)
	if err != nil {
		fmt.Println(err)
		log.Fatal("cannot open sheet")
	}

	var returnArray []string
	var targetCol int

	for x, row := range rows {

		if len(row) != 0 && row[0] == findRow {

			for a, col := range cols {
				if len(col) != 0 && col[x] == findCol {
					targetCol = a
				}
			}

			for i, innerRow := range rows {

				if i <= x {
					continue
				}

				if len(innerRow) != 0 {
					returnArray = append(returnArray, innerRow[targetCol])
					fmt.Print(innerRow[targetCol], "\n")
				} else {
					break
				}
			}
		} else {
			continue
		}
	}

	return returnArray, nil

}

func validateSheetName(sheetname string, regex string) bool {

	RegExp := regexp.MustCompile(regex)

	return RegExp.MatchString(sheetname)
}

func GetAllSheetsRegex(filepath string, regex string) (error, []string) {

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		return err, nil
	}

	sheets := f.GetSheetList()

	var validatedSheets  []string

	for _, sheet := range sheets {
		if validateSheetName(sheet, regex) {
			validatedSheets = append(validatedSheets, sheet)
		}
	}

	return nil, validatedSheets
}

func testForEmptyCells(filepath string, regex string) error {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	numSheets := f.SheetCount

	fmt.Printf("Sheetcount: %s \n", strconv.Itoa(numSheets))

	for i := 1; i <= numSheets; i++ {
		sheetname := f.GetSheetName(i)

		if !validateSheetName(sheetname, regex) {
			fmt.Printf("Sheetname: %s with index %s does not match the regexp -> skipping \n", sheetname, strconv.Itoa(i))
		} else {
		}

	}
	return nil

}
