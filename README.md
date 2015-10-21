# qmongo
Mongo DB query builder for go, works with mgo driver

# Example
```go


func (c *bookController) ShowBook() Response {
	var book Book
	
	
	c.DB.C("books").Find(
	  // same as bson.D{{"id":bson.D{{"$eq":c.P.Get("bookId")}}}}
	  Eq("id", c.P.Get("bookId")),
  ).One(&book)
	
	if book.Id == "" {
		return c.Goto("main.bookController.ShowBooks")
	}
	
	return c.SendViewLayout("book.go.html", &book)
	
}


func (c *bookController) SearchBooks() Response {
	var books []*Book

	var exprs []interface{}

	if value := c.R.FormValue("id"); value != "" {
		exprs = append(exprs, RegEx("id", value, "ig"))
	}
	
	if value := c.R.FormValue("name"); value != "" {
		exprs = append(exprs, RegEx("name", value, "ig"))
	}
	
	if value := c.R.FormValue("description"); value != "" {
		exprs = append(exprs, RegEx("description", value, "ig"))
	}
	
	c.DB.C("books").Find(And(
		exprs...
	)).All(&books)
	
	return c.SendViewLayout("books.go.html", books)
}


```
