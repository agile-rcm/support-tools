package tools

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/tealeg/xlsx"
	"regexp"
	"strconv"
	"strings"
)

type outputer func(s string, n string)
type outputerIndex func(s string)

// Generate a new CSV file from an Excel sheet identified by filename and sheet index
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

// Generate a new CSV file for all sheets in an Excel that match the given regex.
func GenerateNewCsvAll(filepath string, delimiter string, outputf outputer, regex string) error {

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	numSheets := f.SheetCount

	fmt.Printf("Sheetcount: %s \n", strconv.Itoa(numSheets))

	for i := 0; i <= numSheets; i++ {
		sheetname := f.GetSheetName(i)

		if !validateSheetName(sheetname, regex) {
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

// Get a table "col-block" out of a given Excel sheet. The block will contain every row from the given row/col to the next
// (really) empty row without containing the initially given row/col.
// Example: [TestName1 TestName2 TestName3 TestName4] while given "Block2" and "Name" as search coords. (Refer to test/TestWorkbook.xlsx)
func GetBlockByCoords(filepath string, findRow string, findCol string, sheet string) ([]string, error) {

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return nil, err
	}

	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	cols, err := f.GetCols(sheet)
	if err != nil {
		return nil, err
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

// Validate the given sheetname against a given regex
func validateSheetName(sheetname string, regex string) bool {

	RegExp := regexp.MustCompile(regex)

	return RegExp.MatchString(sheetname)
}

// Return a list of sheet names that are already validated against the given regex
// e.g. find all sheet starting with "SomeName..."
func GetAllSheetsRegex(filepath string, regex string) ([]string, error) {

	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	sheets := f.GetSheetList()

	var validatedSheets  []string

	for _, sheet := range sheets {
		if validateSheetName(sheet, regex) {
			validatedSheets = append(validatedSheets, sheet)
		}
	}

	return validatedSheets, nil
}