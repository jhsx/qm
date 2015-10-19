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

func (doc *Builder ) Set(k string, v interface{}) *Builder {
	doc.doc = append(doc.doc, bson.DocElem{k, v})
	return doc
}


// -- eq --
type eq struct {
	Eq interface{} `bson:"$eq"`
}

func (doc *Builder) Eq(k string, v interface{}) *Builder {
	doc.Set(k, eq{v})
	return doc
}

func Eq(k string, v interface{}) *Builder {
	return NewBuilder().Eq(k, v)
}


// -- neq --
type neq struct {
	Neq interface{} `bson:"$ne"`
}

func (doc *Builder) Neq(k string, v interface{}) *Builder {
	doc.Set(k, bson.D{{"$neq", v}})
	return doc
}

func Neq(k string, v interface{}) *Builder {
	return NewBuilder().Neq(k, v)
}


// -- RegEx --
func (doc *Builder) RegEx(k, pattern, options string) *Builder {
	doc.Set(k, bson.RegEx{Pattern:pattern, Options:options})
	return doc
}

func RegEx(k, pattern, options string) *Builder {
	return NewBuilder().RegEx(k, pattern, options)
}


// -- And --
func (doc *Builder ) And(docs ...interface{}) *Builder {
	doc.doc = append(doc.doc, bson.DocElem{"$and", docs})
	return doc
}

func And(docs ...interface{}) *Builder {
	return NewBuilder().And(docs...)
}


// -- Or --
func (doc *Builder ) Or(docs ...interface{}) *Builder {
	doc.doc = append(doc.doc, bson.DocElem{"$or", docs})
	return doc
}

func Or(docs ...interface{}) *Builder {
	return NewBuilder().Or(docs...)
}


// -- NotOr --
func (doc *Builder ) NotOr(docs ...interface{}) *Builder {
	doc.doc = append(doc.doc, bson.DocElem{"$nor", docs})
	return doc
}

func NotOr(docs ...interface{}) *Builder {
	return NewBuilder().Or(docs...)
}