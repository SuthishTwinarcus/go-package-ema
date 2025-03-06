package yonoma

import "fmt"

type Contact struct {
	Email  string `json:"email"`
	Status string `json:"status"`
	Data   struct {
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		Phone       string `json:"phone"`
		DateOfBirth string `json:"dateOfBirth"`
		Address     string `json:"address"`
		City        string `json:"city"`
		State       string `json:"state"`
		Country     string `json:"country"`
		Zipcode     string `json:"zipcode"`
	} `json:"data"`
}
type Status struct {
	Status string `json:"status"`
}
type ContactTag struct {
	TagId string `json:"tag_id"`
}

func (c *Client) CreateContact(listID string, contact Contact) ([]byte, error) {
	endpoint := fmt.Sprintf("contacts/%s/create", listID)
	return c.request("POST", endpoint, contact)
}

func (c *Client) UnsubscribeContact(listID string, contactId string, status Status) ([]byte, error) {
	endpoint := fmt.Sprintf("contacts/%s/status/%s", listID, contactId)
	return c.request("POST", endpoint, status)
}

func (c *Client) AddTag(contactId string, contactTag ContactTag) ([]byte, error) {
	endpoint := fmt.Sprintf("contacts/tags/%s/add", contactId)
	return c.request("POST", endpoint, contactTag)
}

func (c *Client) RemoveContactTag(contactId string, contactTag ContactTag) ([]byte, error) {
	endpoint := fmt.Sprintf("contacts/tags/%s/remove", contactId)
	return c.request("POST", endpoint, contactTag)
}
