# qmongo
Mongo DB query builder for go, works with mgo driver

# Example
```go

func (c *YourController) ShowBooks() mvc.Response{
	var books []Book
	q := qm.Builder()
	if value:=c.FormVar("publish_before"); value != ""{
		q.And(Lte("publish_date",value))
	}
	if value:=c.FormVar("title"); value != ""{
		q.And(Eq("title",value))
	}
	...
	c.Db.C("books").Find(q).All(&books)
	return c.SendView("showbooks.go.html",books)
}

```
