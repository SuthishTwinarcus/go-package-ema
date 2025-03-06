package yonoma

type List struct {
	Name string `json:"list_name"`
}

func (c *Client) ListLists() ([]byte, error) {
	return c.request("GET", "lists/list", nil)
}

func (c *Client) CreateList(list List) ([]byte, error) {
	return c.request("POST", "lists/create", list)
}

func (c *Client) UpdateList(listID string, list List) ([]byte, error) {
	return c.request("POST", "lists/"+listID+"/update", list)
}

func (c *Client) RetrieveList(listID string) ([]byte, error) {
	return c.request("GET", "lists/"+listID, nil)
}

func (c *Client) DeleteList(listID string) ([]byte, error) {
	return c.request("POST", "lists/"+listID+"/delete", nil)
}
