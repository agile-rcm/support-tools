package crowd

import (
	"github.com/agile-rcm/crowd-go"
)

func AddUserToGroup(crowdUrl, crowdApp, crowdPassword, user, group string, skipVerify bool) error {

	api, err := crowd.NewAPI(crowdUrl, crowdApp, crowdPassword, skipVerify)

	if err != nil {
		return err
	}

	err = api.AddUserToGroup(user, group)

	if err != nil {
		return err
	}

	return nil

}

func CreateGroup(crowdUrl, crowdApp, crowdPassword, groupName, groupDescription string, skipVerify, isActive bool) error {
	api, err := crowd.NewAPI(crowdUrl, crowdApp, crowdPassword, skipVerify)

	if err != nil {
		return err
	}

	err = api.CreateGroup(groupName, groupDescription, isActive)

	if err != nil {
		return err
	}

	return nil

}

func RemoveGroup(crowdUrl, crowdApp, crowdPassword, groupName string, skipVerify bool) error {
	api, err := crowd.NewAPI(crowdUrl, crowdApp, crowdPassword, skipVerify)

	if err != nil {
		return err
	}

	err = api.RemoveGroup(groupName)

	if err != nil {
		return err
	}

	return nil

}

func AddChildGroupMembership(crowdUrl, crowdApp, crowdPassword, parentGroupName, childGroupName string, skipVerify bool) error {
	api, err := crowd.NewAPI(crowdUrl, crowdApp, crowdPassword, skipVerify)

	if err != nil {
		return err
	}

	err = api.AddChildGroupMembership(parentGroupName, childGroupName)

	if err != nil {
		return err
	}

	return nil

}

func AddParentGroupMembership(crowdUrl, crowdApp, crowdPassword, parentGroupName, childGroupName string, skipVerify bool) error {
	api, err := crowd.NewAPI(crowdUrl, crowdApp, crowdPassword, skipVerify)

	if err != nil {
		return err
	}

	err = api.AddParentGroupMembership(parentGroupName, childGroupName)

	if err != nil {
		return err
	}

	return nil

}

func AddUser(crowdUrl, crowdApp, crowdPassword, userName, userPassword, userFirstName, userLastName, userDisplayName, userEmail string, skipVerify, isActive bool) error {

	api, err := crowd.NewAPI(crowdUrl, crowdApp, crowdPassword, skipVerify)

	if err != nil {
		return err
	}

	err = api.AddUser(userName, userPassword, userFirstName, userLastName, userDisplayName, userEmail, isActive)

	if err != nil {
		return err
	}

	return nil

}

func GetUser(crowdUrl, crowdApp, crowdPassword, userName string, skipVerify bool) (*crowd.User, error) {

	api, err := crowd.NewAPI(crowdUrl, crowdApp, crowdPassword, skipVerify)

	if err != nil {
		return nil, err
	}

	user, err := api.GetUser(userName)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func UpdateUser(crowdUrl, crowdApp, crowdPassword, userName, userFirstName, userLastName, userDisplayName, userEmail string, skipVerify, isActive bool) error {

	api, err := crowd.NewAPI(crowdUrl, crowdApp, crowdPassword, skipVerify)

	if err != nil {
		return err
	}

	err = api.UpdateUser(userName, userFirstName, userLastName, userDisplayName, userEmail, isActive)

	if err != nil {
		return err
	}

	return nil

}

func RemoveUser(crowdUrl, crowdApp, crowdPassword, userName string, skipVerify bool) error {
	api, err := crowd.NewAPI(crowdUrl, crowdApp, crowdPassword, skipVerify)

	if err != nil {
		return err
	}

	err = api.RemoveUser(userName)

	if err != nil {
		return err
	}

	return nil

}

func GetUserAttributes(crowdUrl, crowdApp, crowdPassword, userName string, skipVerify bool) (*crowd.Attributes, error) {

	api, err := crowd.NewAPI(crowdUrl, crowdApp, crowdPassword, skipVerify)

	if err != nil {
		return nil, err
	}

	attributes, err := api.GetUserAttributes(userName)

	if err != nil {
		return nil, err
	}

	return attributes, nil

}

func StoreUserAttribute(crowdUrl, crowdApp, crowdPassword, userName, attributeName, attributeValue string, skipVerify bool) error {

	api, err := crowd.NewAPI(crowdUrl, crowdApp, crowdPassword, skipVerify)

	if err != nil {
		return err
	}

	attributes := crowd.Attributes{
		Attributes: []*crowd.Attribute{
			{
				Name:   attributeName,
				Values: []string{attributeValue},
			},
		},
	}

	err = api.StoreUserAttributes(userName, &attributes)

	if err != nil {
		return err
	}

	return nil

}

func RemoveUserAttribute(crowdUrl, crowdApp, crowdPassword, userName, attributeName string, skipVerify bool) error {
	api, err := crowd.NewAPI(crowdUrl, crowdApp, crowdPassword, skipVerify)

	if err != nil {
		return err
	}

	err = api.RemoveUserAttribute(userName, attributeName)

	if err != nil {
		return err
	}

	return nil

}

func RemoveUserFromGroup(crowdUrl, crowdApp, crowdPassword, userName, groupName string, skipVerify bool) error {
	api, err := crowd.NewAPI(crowdUrl, crowdApp, crowdPassword, skipVerify)

	if err != nil {
		return err
	}

	err = api.RemoveUserFromGroup(userName, groupName)

	if err != nil {
		return err
	}

	return nil

}