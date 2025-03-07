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
    Email:  "glennjacob@example.com",
    Status: "Subscribed" | "Unsubscribed",
    Data: yonoma.ContactData{
        FirstName: string,
        LastName:  string,
        Phone:     string,
        Gender:    string,
        Address:   string,
        City:      string,
        State:     string,
        Country:   string,
        Zipcode:   string,
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
response, err := client.AddContactTag("Contact Id", contactData)
```
#### Remove tag from contact
```go
contactData := yonoma.TagId{
	TagId: "Tag Id",
}
response, err := client.RemoveContactTag("Contact Id", contactData)

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


