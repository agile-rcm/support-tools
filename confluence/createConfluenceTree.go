package confluence

// TODO: Godoc nutzen -> mehr kommentieren -> https://blog.golang.org/godoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

func WalkDirAndCreate(
	userId string, password string, endpoint string, insecure bool, debug bool, directoryName string,
	confluenceParentPage string, minorEdit bool, timestamp bool, spacekey string) error {
	if directoryName != "/" && directoryName[len(directoryName)-1:] == "/" {
		directoryName = directoryName[:len(directoryName)-1]
	}

	a := []string{}
	err := WalkDirAndCreateRecursive(
		userId, password, endpoint, insecure, debug, directoryName, a, 0,
		confluenceParentPage, minorEdit, timestamp, spacekey)
	if err != nil {
		return err
	}
	return nil
}

func WalkDirAndCreateRecursive(
	userId string, password string, endpoint string, insecure bool, debug bool, directoryName string,
	dirs []string, counter int, anchestor string, minorEdit bool, timestamp bool, spacekey string) error {

	// formattedOutput := fmt.Sprintf("%-60s%-50s%-50s", directoryName, anchestor, dirs)
	// formattedOutput := fmt.Sprintf("%-60s%-50s", directoryName, anchestor)
	// fmt.Println(formattedOutput)

	counter++
	files, _ := ioutil.ReadDir(directoryName)

	for _, f := range files {
		var newPath string
		if directoryName != "/" {
			newPath = fmt.Sprintf("%s/%s", directoryName, f.Name())
		} else {
			newPath = fmt.Sprintf("%s%s", directoryName, f.Name())
		}

		if f.IsDir() {

			if !processed(newPath, dirs) {
				dirs = append(dirs, newPath)
				// formattedOutput := fmt.Sprintf("X: %-60s%-50s", f.Name(),anchestor)
				// fmt.Println(formattedOutput)

				pageName := f.Name()
				contentString := getContentStringFromFile(newPath + "/" + pageName + ".html")

				// create page for folder in the source-structure
				// IMPORTANT: the folder requires a page with the same name in it
				// - + folderA
				//   |- folderA.html    <<========
				//   |- name2.html
				//   | + folderB
				//     | - folderB.html   <<========
				//
				err := CreatePage(
					userId, password, endpoint, insecure, debug,
					anchestor, pageName, contentString, minorEdit, spacekey, timestamp)
				if err != nil {
					return err
				}
				// go into this folder
				anchestorOld := anchestor
				anchestor = pageName
				err = WalkDirAndCreateRecursive(
					userId, password, endpoint, insecure,
					debug, newPath, dirs, counter, anchestor, minorEdit, timestamp, spacekey)
				if err != nil {
					return err
				}
				anchestor = anchestorOld
			}
		} else {

			file := filepath.Base(newPath)
			pageName := strings.TrimSuffix(file, filepath.Ext(file))

			r, _ := regexp.Compile(".*/")
			fileDir := r.ReplaceAllString(filepath.Dir(newPath), "")

			// formattedOutput := fmt.Sprintf("\n  ==> %-60s%-50s", file, fileDir)
			// fmt.Println(formattedOutput)

			// IMPORTANT:
			// dont create a page with the same name as the foldername, this page was already created
			// in the previous step during the folder-Creation (while treversing the folder recursive
			if strings.Compare(pageName, fileDir) != 0 {

				contentString := getContentStringFromFile(newPath)

				formattedOutput := fmt.Sprintf("%-60s%-50s", anchestor, pageName)
				fmt.Println(formattedOutput)

				//			formattedOutput = fmt.Sprintf("\n\na/p: %-60s%-50s\n\n", anchestor, pageName)
				//			fmt.Println(formattedOutput)
				err := CreatePage(
					userId, password, endpoint, insecure, debug, anchestor, pageName, contentString,
					minorEdit, spacekey, timestamp)
				if err != nil {
					return err
				}
			}

		}
	}
	return nil
}

func processed(fileName string, processedDirectories []string) bool {
	for i := 0; i < len(processedDirectories); i++ {
		if processedDirectories[i] != fileName {
			continue
		}
		return true
	}
	return false
}

func getContentStringFromFile(fileName string) string {

	contentBuffer, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	contentString := string(contentBuffer)
	return contentString
}
