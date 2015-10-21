# qmongo
Mongo DB query builder for go, works with mgo driver

# Example
```go

import(
	"github.com/jhsx/qm"
)

func (c *bookController) ShowBook() Response {
	var book Book
	
	
	c.DB.C("books").Find(
	  // same as bson.D{{"id":bson.D{{"$eq":c.P.Get("bookId")}}}}
	  qm.Eq("id", c.P.Get("bookId")),
  	).One(&book)
	
	if book.Id == "" {
		return c.Goto("main.bookController.ShowBooks")
	}
	
	return c.SendViewLayout("book.go.html", &book)
	
}


func (c *bookController) SearchBooks() Response {
	var books []*Book

	query := qm.NewBuilder()

	if value := c.R.FormValue("id"); value != "" {
		query.And(qm.RegEx("id",value,"ig"))
	}

	if value := c.R.FormValue("name"); value != "" {
		query.And(qm.RegEx("name",value,"ig"))
	}

	if value := c.R.FormValue("description"); value != "" {
		query.And(qm.RegEx("description",value,"ig"))
	}
	
	c.DB.C("books").Find(query).All(&books)
	
	return c.SendViewLayout("books.go.html", books)
}


```
