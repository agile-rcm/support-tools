package tools_test

import (
	"bufio"
	"fmt"
	"git.agiletech.de/AgileRCM/support-tools/tools"
	"os"
	"testing"
)

func TestGetAllSheetsRegex(t *testing.T) {
	sheets, err := tools.GetAllSheetsRegex("test/TestWorkbook.xlsx", `^([Aa][nother])`)

	if err != nil {
		t.Errorf("Error while getting sheets: %s", err)
	}

	if sheets[0] != "AnotherSheet" || sheets[1] != "AnotherSheet2" {
		t.Errorf("Expected sheetnames to be AnotherSheet and AnotherSheet2 but got %s and %s", sheets[0], sheets[1])
	}
}

func TestGetBlockByCoords(t *testing.T) {
	block, err := tools.GetBlockByCoords("test/TestWorkbook.xlsx", "Block2", "Name", "Testsheet")

	if err != nil {
		t.Errorf("Error getting block contents: %s", err)
	}

	if block[0] != "TestName1" {
		t.Errorf("Expected block[0] to be \"TestName1\" but got: %s", block[0])
	}
}

func TestGenerateNewCsvAll(t *testing.T) {
	printer := func(s string, n string) {
		f, err := os.OpenFile("./test/"+n+".csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			// return cli.NewExitError(err.Error(), 1)
			panic(err)
		}

		defer f.Close()

		write, err := f.WriteString(s)
		fmt.Printf("wrote %d bytes\n", write)

	}

	err := tools.GenerateNewCsvAll("test/TestWorkbook.xlsx", ";", printer, `^([Aa][nother])`)

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if _, err := os.Stat("./test/AnotherSheet.csv"); err == nil {
		fmt.Printf("File ./test/AnotherSheet.csv was created\n");

		file, err := os.Open("./test/AnotherSheet.csv")
		if err != nil {
			t.Errorf("Error reading CSV file: %s", err)
		}
		defer file.Close()

		var lines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if lines[0] != "\"AnotherBlock1\";\"Key\";\"Name\";\"Something\"" {
			t.Errorf("CSV file does not have expected content. Content found: %s", lines[0])
		}
		if lines[1] != "\"TestAnother\";\"TestKeyBlk1\";\"TestNameBlk1\";\"SomethingBlk1\"" {
			t.Errorf("CSV file does not have expected content. Content found: %s", lines[0])
		}

		os.Remove("./test/AnotherSheet.csv")
	} else {
		t.Errorf("File ./test/AnotherSheet.csv was not created as expected\n")
	}

	if _, err := os.Stat("./test/AnotherSheet2.csv"); err == nil {
		fmt.Printf("File ./test/AnotherSheet2.csv was created\n");

		file, err := os.Open("./test/AnotherSheet2.csv")
		if err != nil {
			t.Errorf("Error reading CSV file: %s", err)
		}
		defer file.Close()

		var lines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if lines[0] != "\"AnotherBlock2\";\"Key\";\"Name\";\"Something\"" {
			t.Errorf("CSV file does not have expected content. Content found: %s", lines[0])
		}
		if lines[1] != "\"TestAnother\";\"TestKeyBlk1\";\"TestNameBlk1\";\"SomethingBlk1\"" {
			t.Errorf("CSV file does not have expected content. Content found: %s", lines[0])
		}

		os.Remove("./test/AnotherSheet2.csv")
	} else {
		t.Errorf("File ./test/AnotherSheet2.csv was not created as expected\n")
	}

}

func TestGenerateNewCsvByIndex(t *testing.T) {
	printer := func(s string) {
		f, err := os.OpenFile("./test/GenerateCSVbyIndex.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			// return cli.NewExitError(err.Error(), 1)
			panic(err)
		}

		defer f.Close()

		write, err := f.WriteString(s)
		fmt.Printf("wrote %d bytes\n", write)

	}

	err := tools.GenerateNewCsvByIndex("test/TestWorkbook.xlsx", 1, printer, ";")

	if err != nil {
		t.Errorf("Error generating CSV by Index: %s", err)
	}

	if _, err := os.Stat("./test/GenerateCSVbyIndex.csv"); err == nil {
		fmt.Printf("File ./test/GenerateCSVbyIndex.csv was created\n");

		file, err := os.Open("./test/GenerateCSVbyIndex.csv")
		if err != nil {
			t.Errorf("Error reading CSV file: %s", err)
		}
		defer file.Close()

		var lines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if lines[0] != "\"AnotherBlock1\";\"Key\";\"Name\";\"Something\"" {
			t.Errorf("CSV file does not have expected content. Content found: %s", lines[0])
		}
		if lines[1] != "\"TestAnother\";\"TestKeyBlk1\";\"TestNameBlk1\";\"SomethingBlk1\"" {
			t.Errorf("CSV file does not have expected content. Content found: %s", lines[0])
		}

		os.Remove("./test/GenerateCSVbyIndex.csv")
	}else{
		t.Errorf("Error testing CSV file: %s", err)
	}
}

func TestGetLinesByKeyword(t *testing.T) {
	result, err := tools.GetLinesByKeyword("test/TestWorkbook.xlsx", "Testdata", 3, "Testsheet")

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println(result)

}