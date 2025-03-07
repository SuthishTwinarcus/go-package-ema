## [Yonoma](https://yonoma.io/) Go Email Marketing SDK
### Install
```bash
go get github.com/YonomaHQ/yonoma
```
### Setup
First, you need to get an API key:
```go
client := yonoma.NewClient("api-key") 
```
### Usage
## Contacts
#### Create new contact
```go
contactData := yonoma.Contact{
    Email:  "johndavid12@gmail.com",
    Status: "Subscribed",
    Data: struct {
        FirstName   string `json:"firstName"`
        LastName    string `json:"lastName"`
        Phone       string `json:"phone"`
        DateOfBirth string `json:"dateOfBirth"`
        Address     string `json:"address"`
        City        string `json:"city"`
        State       string `json:"state"`
        Country     string `json:"country"`
        Zipcode     string `json:"zipcode"`
    }{
        FirstName:   "Johndavid",
        LastName:    "12",
        Phone:       "+1234567890",
        DateOfBirth: "2-08-1994",
        Address:     "123 Main St",
        City:        "New York",
        State:       "NY",
        Country:     "USA",
        Zipcode:     "10001",
    },
}
response, err := client.CreateContact("List Id", contactData) 
```
#### Update contact
```go
contactData := yonoma.Status{
	Status: "Subscribed" | "Unsubscribed",
}
response, err := client.UnsubscribeContact("List Id", "Contact Id", contactData)
```
#### Add tag to contact
```go
contactData := yonoma.TagId{
	TagId: "Tag Id",
}
response, err := client.AddTag("Contact Id", contactData)
```
#### Remove tag from contact
```go
contactData := yonoma.TagId{
	TagId: "Tag Id",
}
response, err := client.RemoveTag("Contact Id", contactData)

```
## Managing Tags
#### Create a Tag
```go
tagData := yonoma.Tag{
	Name: "Tag Name",
}
response, err := client.CreateTag(tagData)
```
#### Update a Tag
```go
tagData := yonoma.Tag{
	Name: "Tag Name",
}
response, err := client.UpdateTag(tagData)
```
#### Delete a Tag
```go
response, err := client.DeleteTag("Tag Id")
```
#### Retrieve a Specific Tag
```go
response, err := client.RetrieveTag("Tag Id")
```
#### List All Tags
```go
response, err := client.ListTags()
```
## Managing Lists
#### Create a List
```go
listData := yonoma.List{
	Name: "List Name",
}
response, err := client.CreateList(listData)

```
#### Update a List
```go
listData := yonoma.List{
	Name: "List Name",
}
response, err := client.UpdateList("List Id", listData)
```
#### Delete a List
```go
response, err := client.DeleteList("List Id")
```
#### Retrieve a Specific List
```go
response, err := client.RetrieveList("List Id")
```
#### List All Lists
```go
response, err := client.ListLists()
```


