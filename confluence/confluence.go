package confluence

// TODO: Godoc nutzen -> mehr kommentieren -> https://blog.golang.org/godoc

import (
	"fmt"
	"github.com/virtomize/confluence-go-api"
	"log"
)

func GetUser(userId string, password string, endpoint string, insecure bool, debug bool) error {

	// initialize a new api instance
	api, err := goconfluence.NewAPI(endpoint, userId, password)
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

func CreatePage(userId string, password string, endpoint string, insecure bool, debug bool, parentTitle string, title string, newPageContent string, minorEdit bool, spacekey string) error {

	api, err := goconfluence.NewAPI(endpoint, userId, password)
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
			log.Fatal(err)
			fmt.Printf("%+v\n", result)
		}

	} else {

		pageId := pageTitle.Results[0].ID
		history, err := api.GetHistory(pageId)
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
			log.Fatal(err)
			fmt.Printf("%+v\n", cc)
		}

	}
	return nil
}
