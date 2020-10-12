package confluence

// TODO: Godoc nutzen -> mehr kommentieren -> https://blog.golang.org/godoc

import (
	"fmt"
	"git.agiletech.de/AgileRCM/support-tools/context"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func WalkDirAndCreate(ctx context.Context, directoryName string, confluenceParentPage string, minorEdit bool, timestamp bool) error {
	if directoryName != "/" && directoryName[len(directoryName)-1:] == "/" {
		directoryName = directoryName[:len(directoryName)-1]
	}

	a := []string{}
	WalkDirAndCreateRecursive(ctx, directoryName, a, 0, confluenceParentPage, minorEdit, timestamp)
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

func getContentString(pageName string, timestamp bool) string {

	contentBuffer, err := ioutil.ReadFile(pageName)
	if err != nil {
		log.Fatal(err)
	}
	contentString := string(contentBuffer)
	if timestamp {
		timeStampTxt := time.Now().String()
		contentString = "<p>" + timeStampTxt + "</p>" + contentString
	}
	return contentString
}

func WalkDirAndCreateRecursive(
	ctx context.Context, directoryName string, dirs []string, counter int, anchestor string, minorEdit bool, timestamp bool) {

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
				contentString := getContentString(newPath+"/"+pageName+".html", timestamp)

				// create page for folder in the source-structure
				// IMPORTANT: the folder requires a page with the same name in it
				// - + folderA
				//   |- folderA.html    <<========
				//   |- name2.html
				//   | + folderB
				//     | - folderB.html   <<========
				//
				CreatePage(ctx, anchestor, pageName, contentString, minorEdit)
				// go into this folder
				anchestorOld := anchestor
				anchestor = pageName
				WalkDirAndCreateRecursive(ctx, newPath, dirs, counter, anchestor, minorEdit, timestamp)
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

				contentString := getContentString(newPath, timestamp)

				formattedOutput := fmt.Sprintf("%-60s%-50s", anchestor, pageName)
				fmt.Println(formattedOutput)

				//			formattedOutput = fmt.Sprintf("\n\na/p: %-60s%-50s\n\n", anchestor, pageName)
				//			fmt.Println(formattedOutput)
				CreatePage(ctx, anchestor, pageName, contentString, minorEdit)
			}

		}
	}
}
