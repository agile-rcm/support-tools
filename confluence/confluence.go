package confluence

// TODO: Godoc nutzen -> mehr kommentieren -> https://blog.golang.org/godoc

import (
	"fmt"
	"github.com/virtomize/confluence-go-api"
	"log"
	"os"
	"time"
)

func GetUser(userId string, password string, endpoint string, insecure bool, debug bool) error {

	// initialize a new api instance
	api, err := goconfluence.NewAPI(endpoint, userId, password)

	// TODO - see messages
	api.VerifyTLS(insecure)

	goconfluence.SetDebug(debug)

	if err != nil {
		log.Fatal(err)
	}

	// get current user information
	currentUser, err := api.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%-20s%s\n%-20s%s\n%-20s%s\n%-20s%s\n%-20s%s\n\n",
		"Type:", currentUser.Type,
		"Userid:", currentUser.Username,
		"AccountId:", currentUser.AccountID,
		"UserKey:", currentUser.UserKey,
		"DisplayName:", currentUser.DisplayName)

	return nil
}

func GetContent(userId string, password string, endpoint string, insecure bool, debug bool, title string, spacekey string) (*goconfluence.ContentSearch, error) {
	api, err := goconfluence.NewAPI(endpoint, userId, password)

	// TODO - see messages
	api.VerifyTLS(insecure)

	goconfluence.SetDebug(debug)

	if err != nil {
		log.Fatal(err)
	}

	c, err := api.GetContent(goconfluence.ContentQuery{
		SpaceKey: spacekey,
		Title:    title,
	})
	if err != nil {
		log.Fatal(err)
	}

	return c, nil
}

func CreatePage(
	userId string,
	password string,
	endpoint string,
	insecure bool,
	debug bool,
	parentTitle string,
	title string,
	newPageContent string,
	minorEdit bool,
	spacekey string,
	timestamp bool) error {

	api, err := goconfluence.NewAPI(endpoint, userId, password)

	// TODO - see messages
	api.VerifyTLS(insecure)

	goconfluence.SetDebug(debug)

	if err != nil {
		log.Fatal(err)
	}

	pageParentTitle, err := GetContent(userId, password, endpoint, insecure, debug, parentTitle, spacekey)

	if err != nil {
		log.Fatal(err)
	}

	if len(pageParentTitle.Results) == 0 {
		log.Fatal("Error : Parentpage \"" + parentTitle + "\" not found!")
	}

	if len(pageParentTitle.Results) > 1 {
		log.Fatal("Error : Parentpage exists more than one times!")
	}

	parentTitleId := pageParentTitle.Results[0].ID

	pageTitle, err := api.GetContent(goconfluence.ContentQuery{
		SpaceKey: spacekey,
		Title:    title,
	})

	if err != nil {
		log.Fatal(err)
	}

	newPageContent = getContentStringWithTimestamp(newPageContent, timestamp)

	if len(pageTitle.Results) == 0 {

		result, err := api.CreateContent(&goconfluence.Content{
			Type:  "page", // can also be blogpost
			Title: title,  // page title
			Ancestors: []goconfluence.Ancestor{
				{
					ID: parentTitleId, // ancestor-id optional if you want to create sub-pages
				},
			},
			Body: goconfluence.Body{
				Storage: goconfluence.Storage{
					Value:          newPageContent,
					Representation: "storage",
				},
			},
			Version: goconfluence.Version{
				Number:    1,
				MinorEdit: true,
			},
			Space: goconfluence.Space{
				Key: spacekey,
			},
		})

		if err != nil {
			fmt.Printf("%+v\n", result)
			log.Fatal(err)
		}

	} else {

		pageId := pageTitle.Results[0].ID
		history, err := api.GetHistory(pageId)

		// TODO - see messages
		newVersion := history.LastUpdated.Number

		data := &goconfluence.Content{
			Type:  "page", // can also be blogpost
			Title: title,  // page title
			ID:    pageId,
			Ancestors: []goconfluence.Ancestor{
				{
					ID: parentTitleId, // ancestor-id optional if you want to create sub-pages
				},
			},
			Body: goconfluence.Body{
				Storage: goconfluence.Storage{
					Value:          newPageContent,
					Representation: "storage",
				},
			},

			Version: goconfluence.Version{
				Number:    newVersion + 1,
				MinorEdit: minorEdit,
			},
			Space: goconfluence.Space{
				Key: spacekey, // Space
			},
		}

		cc, err := api.UpdateContent(data)
		if err != nil {
			fmt.Printf("%+v\n", cc)
			log.Fatal(err)
		}

	}
	return nil
}

func CreateAttachment(userId string, password string, endpoint string, insecure bool, debug bool, filepath string, pageTitle string, spaceKey string, attachmentName string) error {

	// Open API
	api, err := goconfluence.NewAPI(endpoint,userId,password)

	if err != nil {
		return err
	}

	api.VerifyTLS(insecure)
	goconfluence.SetDebug(debug)

	// Search Title ID
	pageParentTitle, err := GetContent(userId, password, endpoint, insecure, debug, pageTitle, spaceKey)

	if err != nil {
		log.Fatal(err)
	}

	if len(pageParentTitle.Results) == 0 {
		log.Fatal("Error : Page \"" + pageTitle + "\" not found!")
	}

	if len(pageParentTitle.Results) > 1 {
		log.Fatal("Error : Page exists more than one times!")
	}

	pageTitleId := pageParentTitle.Results[0].ID

	// Rename File
	file, err := os.Open(filepath)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = api.UploadAttachment(pageTitleId, attachmentName, file)

	if err != nil {
		return err
	}

	return nil
}

func DeleteAttachment(userId string, password string, endpoint string, insecure bool, debug bool, pageTitle string, spaceKey string, attachmentName string) error {

	// Open API
	api, err := goconfluence.NewAPI(endpoint,userId,password)

	if err != nil {
		return err
	}

	api.VerifyTLS(insecure)
	goconfluence.SetDebug(debug)

	// Search Title ID
	pageParentTitle, err := GetContent(userId, password, endpoint, insecure, debug, pageTitle, spaceKey)

	if err != nil {
		log.Fatal(err)
	}

	if len(pageParentTitle.Results) == 0 {
		log.Fatal("Error : Page \"" + pageTitle + "\" not found!")
	}

	if len(pageParentTitle.Results) > 1 {
		log.Fatal("Error : Page exists more than one times!")
	}

	pageTitleId := pageParentTitle.Results[0].ID

	attachments, err := api.GetAttachments(pageTitleId)

	for _, result := range attachments.Results{

		if result.Title == attachmentName {

			_, err := api.DelContent(result.ID)

			if err != nil {
				fmt.Println(err)
			}
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func AddLabel(userId string, password string, endpoint string, insecure bool, debug bool, pageTitle string, spaceKey string, label string) error {

	labelToAdd := goconfluence.Label{
		Prefix: "global",
		Name:   label,
		Label:  label,
	}

	var labels []goconfluence.Label

	labels = append(labels, labelToAdd)

	// Open API
	api, err := goconfluence.NewAPI(endpoint,userId,password)

	if err != nil {
		return err
	}

	api.VerifyTLS(insecure)
	goconfluence.SetDebug(debug)

	// Search Title ID
	pageParentTitle, err := GetContent(userId, password, endpoint, insecure, debug, pageTitle, spaceKey)

	if err != nil {
		log.Fatal(err)
	}

	if len(pageParentTitle.Results) == 0 {
		log.Fatal("Error : Page \"" + pageTitle + "\" not found!")
	}

	if len(pageParentTitle.Results) > 1 {
		log.Fatal("Error : Page exists more than one times!")
	}

	pageTitleId := pageParentTitle.Results[0].ID

	_, err = api.AddLabels(pageTitleId, &labels)

	if err != nil {
		return err
	}

	return nil
}

func getContentStringWithTimestamp(contentString string, timestamp bool) string {

	if timestamp {
		timeStampTxt := time.Now().String()
		contentString = "<p>" + timeStampTxt + "</p>" + contentString
	}
	return contentString
}
