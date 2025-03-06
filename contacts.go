package yonoma

import "fmt"

type Contact struct {
	email  string `json:"contact_email"`
	Status string `json:"sub_status"`
	Data   struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"mobile_number"`
		Gender    string `json:"gender"`
		Address   string `json:"address"`
		City      string `json:"city"`
		State     string `json:"state"`
		Country   string `json:"country"`
		Zipcode   string `json:"zipcode"`
	} `json:"data"`
}

func (c *Client) CreateContact(listID string, contact Contact) ([]byte, error) {
	endpoint := fmt.Sprintf("contacts/%s/create", listID)
	return c.request("POST", endpoint, contact)
}

func (c *Client) UnsubscribeContact(listID string, contactId string, contact Contact) ([]byte, error) {
	endpoint := fmt.Sprintf("contacts/%s/status/%s", listID, contactId)
	return c.request("POST", endpoint, contact)
}

func (c *Client) AddTag(contactId string, contact Contact) ([]byte, error) {
	endpoint := fmt.Sprintf("contacts/tags/%s/add", contactId)
	return c.request("POST", endpoint, contact)
}

func (c *Client) RemoveContactTag(contactId string, contact Contact) ([]byte, error) {
	endpoint := fmt.Sprintf("contacts/tags/%s/remove", contactId)
	return c.request("POST", endpoint, contact)
}
