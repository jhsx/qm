package qmongo
import "gopkg.in/mgo.v2/bson"


type Builder struct {
	doc bson.D
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (doc *Builder) GetBSON() (interface{}, error) {
	return doc.doc, nil
}

// Set sets bson field k with the value v row[k]=v
func (doc *Builder ) Set(k string, v interface{}) *Builder {
	numDocs := len(doc.doc)
	for i := 0; i < numDocs; i++ {
		d := &doc.doc[i]
		if d.Name == k {
			d.Value = v
			return doc
		}
	}
	doc.doc = append(doc.doc, bson.DocElem{Name:k, Value:v})
	return doc
}


// -- eq --
type eq struct {
	Eq interface{} `bson:"$eq"`
}

// Eq sets bson field k with row[k][$eq]=v which can be express like row[k] == v
func (doc *Builder) Eq(k string, v interface{}) *Builder {
	doc.Set(k, eq{v})
	return doc
}

// Eq creates a new query object and sets bson field k with row[k][$eq]=v which can be express like row[k] == v
func Eq(k string, v interface{}) *Builder {
	return NewBuilder().Eq(k, v)
}


// -- neq --
type neq struct {
	Neq interface{} `bson:"$ne"`
}

// Neq sets bson field k with row[k][$ne]=v which can be express like row[k] != v
func (doc *Builder) Neq(k string, v interface{}) *Builder {
	doc.Set(k, bson.D{{"$neq", v}})
	return doc
}

// Neq creates a new query object and sets bson field k with row[k][$ne]=v which can be express like row[k] != v
func Neq(k string, v interface{}) *Builder {
	return NewBuilder().Neq(k, v)
}

// RegEx sets bson field k with a regex pattern and options
func (doc *Builder) RegEx(k, pattern, options string) *Builder {
	doc.Set(k, bson.RegEx{Pattern:pattern, Options:options})
	return doc
}

// RegEx creates a new query object and sets bson field k with regex pattern and options
func RegEx(k, pattern, options string) *Builder {
	return NewBuilder().RegEx(k, pattern, options)
}


// And is equivalent to expr0 && expr1 && expr...
func (doc *Builder ) And(docs ...interface{}) *Builder {
	numDocs := len(doc.doc)
	for i := 0; i < numDocs; i++ {
		d := &doc.doc[i]
		if d.Name == "$and" {
			d.Value = append(d.Value.([]interface{}), docs...)
			return doc
		}
	}
	doc.doc = append(doc.doc, bson.DocElem{Name:"$and", Value: docs})
	return doc
}

// And is equivalent to expr0 && expr1 && expr...
func And(docs ...interface{}) *Builder {
	return NewBuilder().And(docs...)
}

// Or is equivalent to expr0 || expr1 || expr...
func (doc *Builder ) Or(docs ...interface{}) *Builder {
	numDocs := len(doc.doc)
	for i := 0; i < numDocs; i++ {
		d := &doc.doc[i]
		if d.Name == "$or" {
			d.Value = append(d.Value.([]interface{}), docs...)
			return doc
		}
	}
	doc.doc = append(doc.doc, bson.DocElem{Name:"$or", Value: docs})
	return doc
}

// Or is equivalent to expr0 || expr1 || expr...
func Or(docs ...interface{}) *Builder {
	return NewBuilder().Or(docs...)
}

// NotOr is equivalent to !expr0 || !expr1 || !expr...
func (doc *Builder ) NotOr(docs ...interface{}) *Builder {
	numDocs := len(doc.doc)
	for i := 0; i < numDocs; i++ {
		d := &doc.doc[i]
		if d.Name == "$nor" {
			d.Value = append(d.Value.([]interface{}), docs...)
			return doc
		}
	}
	doc.doc = append(doc.doc, bson.DocElem{Name:"$nor", Value: docs})
	return doc
}

// NotOr is equivalent to !expr0 || !expr1 || !expr...
func NotOr(docs ...interface{}) *Builder {
	return NewBuilder().Or(docs...)
}
