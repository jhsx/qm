package qm

import "gopkg.in/mgo.v2/bson"

type Document struct {
	doc bson.D
}

func Builder() *Document {
	return &Document{}
}

func New(k string, v interface{}) *Document {
	return Builder().Set(k, v)
}

func (doc *Document) GetBSON() (interface{}, error) {
	return doc.doc, nil
}

// Set sets bson field k with the value v row[k]=v
func (doc *Document) Set(k string, v interface{}) *Document {
	numDocs := len(doc.doc)
	for i := 0; i < numDocs; i++ {
		d := &doc.doc[i]
		if d.Name == k {
			d.Value = v
			return doc
		}
	}
	doc.doc = append(doc.doc, bson.DocElem{Name: k, Value: v})
	return doc
}

// -- lt --
type lt struct {
	Lt interface{} `bson:"$lt"`
}

// Lt sets bson field k with row[k][$lt]=v which can be express like row[k] == v
func (doc *Document) Lt(k string, v interface{}) *Document {
	doc.Set(k, lt{v})
	return doc
}

// Lt creates a new query object and sets bson field k with row[k][$eq]=v which can be express like row[k] == v
func Lt(k string, v interface{}) *Document {
	return Builder().Lt(k, v)
}

// -- lte --
type lte struct {
	Lte interface{} `bson:"$lte"`
}

// Lte sets bson field k with row[k][$lte]=v which can be express like row[k] == v
func (doc *Document) Lte(k string, v interface{}) *Document {
	doc.Set(k, lte{v})
	return doc
}

// Lte creates a new query object and sets bson field k with row[k][$eq]=v which can be express like row[k] == v
func Lte(k string, v interface{}) *Document {
	return Builder().Gt(k, v)
}

// -- gt --
type gt struct {
	Gt interface{} `bson:"$gt"`
}

// Gt sets bson field k with row[k][$gt]=v which can be express like row[k] == v
func (doc *Document) Gt(k string, v interface{}) *Document {
	doc.Set(k, gt{v})
	return doc
}

// Gt creates a new query object and sets bson field k with row[k][$eq]=v which can be express like row[k] == v
func Gt(k string, v interface{}) *Document {
	return Builder().Gt(k, v)
}

// -- gte --
type gte struct {
	Gte interface{} `bson:"$gte"`
}

// Gte sets bson field k with row[k][$gte]=v which can be express like row[k] == v
func (doc *Document) Gte(k string, v interface{}) *Document {
	doc.Set(k, gte{v})
	return doc
}

// Gte creates a new query object and sets bson field k with row[k][$eq]=v which can be express like row[k] == v
func Gte(k string, v interface{}) *Document {
	return Builder().Gt(k, v)
}

// -- eq --
type eq struct {
	Eq interface{} `bson:"$eq"`
}

// Eq sets bson field k with row[k][$eq]=v which can be express like row[k] == v
func (doc *Document) Eq(k string, v interface{}) *Document {
	doc.Set(k, eq{v})
	return doc
}

// Eq creates a new query object and sets bson field k with row[k][$eq]=v which can be express like row[k] == v
func Eq(k string, v interface{}) *Document {
	return Builder().Eq(k, v)
}

// -- neq --
type neq struct {
	Neq interface{} `bson:"$ne"`
}

// Neq sets bson field k with row[k][$ne]=v which can be express like row[k] != v
func (doc *Document) Neq(k string, v interface{}) *Document {
	doc.Set(k, neq{v})
	return doc
}

// Neq creates a new query object and sets bson field k with row[k][$ne]=v which can be express like row[k] != v
func Neq(k string, v interface{}) *Document {
	return Builder().Neq(k, v)
}

// -- in --
type in struct {
	In interface{} `bson:"$in"`
}

// In sets bson field k with row[k][$ne]=v which can be express like row[k] != v
func (doc *Document) In(k string, v ...interface{}) *Document {
	doc.Set(k, in{v})
	return doc
}

// In creates a new query object and sets bson field k with row[k][$ne]=v which can be express like row[k] != v
func In(k string, v ...interface{}) *Document {
	return Builder().In(k, v)
}

// -- nin --
type nin struct {
	Nin interface{} `bson:"$nin"`
}

// Nin sets bson field k with row[k][$ne]=v which can be express like row[k] != v
func (doc *Document) Nin(k string, v ...interface{}) *Document {
	doc.Set(k, nin{v})
	return doc
}

// Nin creates a new query object and sets bson field k with row[k][$ne]=v which can be express like row[k] != v
func Nin(k string, v ...interface{}) *Document {
	return Builder().Nin(k, v)
}

// RegEx sets bson field k with a regex pattern and options
func (doc *Document) RegEx(k, pattern, options string) *Document {
	doc.Set(k, bson.RegEx{Pattern: pattern, Options: options})
	return doc
}

// RegEx creates a new query object and sets bson field k with regex pattern and options
func RegEx(k, pattern, options string) *Document {
	return Builder().RegEx(k, pattern, options)
}

// And is equivalent to expr0 && expr1 && expr...
func (doc *Document) And(docs ...interface{}) *Document {
	numDocs := len(doc.doc)
	for i := 0; i < numDocs; i++ {
		d := &doc.doc[i]
		if d.Name == "$and" {
			d.Value = append(d.Value.([]interface{}), docs...)
			return doc
		}
	}
	doc.doc = append(doc.doc, bson.DocElem{Name: "$and", Value: docs})
	return doc
}

// And is equivalent to expr0 && expr1 && expr...
func And(docs ...interface{}) *Document {
	return Builder().And(docs...)
}

// Or is equivalent to expr0 || expr1 || expr...
func (doc *Document) Or(docs ...interface{}) *Document {
	numDocs := len(doc.doc)
	for i := 0; i < numDocs; i++ {
		d := &doc.doc[i]
		if d.Name == "$or" {
			d.Value = append(d.Value.([]interface{}), docs...)
			return doc
		}
	}
	doc.doc = append(doc.doc, bson.DocElem{Name: "$or", Value: docs})
	return doc
}

// Or is equivalent to expr0 || expr1 || expr...
func Or(docs ...interface{}) *Document {
	return Builder().Or(docs...)
}

// NotOr is equivalent to !expr0 || !expr1 || !expr...
func (doc *Document) NotOr(docs ...interface{}) *Document {
	numDocs := len(doc.doc)
	for i := 0; i < numDocs; i++ {
		d := &doc.doc[i]
		if d.Name == "$nor" {
			d.Value = append(d.Value.([]interface{}), docs...)
			return doc
		}
	}
	doc.doc = append(doc.doc, bson.DocElem{Name: "$nor", Value: docs})
	return doc
}

// NotOr is equivalent to !expr0 || !expr1 || !expr...
func NotOr(docs ...interface{}) *Document {
	return Builder().NotOr(docs...)
}


// -- nin --

// DocSet
func (doc *Document) DocSet(k string, v interface{}) *Document {
	numDocs := len(doc.doc)
	for i := 0; i < numDocs; i++ {
		d := &doc.doc[i]
		if d.Name == "$set" {
			d.Value.(*Document).Set(k, v)
		}
	}

	doc.doc = append(doc.doc, bson.DocElem{Name: "$set", Value: Builder().Set(k, v)})
	return doc
}

// DocSet is equivalent to !expr0 || !expr1 || !expr...
func DocSet(k string, v interface{}) *Document {
	return Builder().DocSet(k, v)
}


// Exists
func (doc *Document) Exists(k string, v interface{}) *Document {
	numDocs := len(doc.doc)
	for i := 0; i < numDocs; i++ {
		d := &doc.doc[i]
		if d.Name == "$exists" {
			d.Value.(*Document).Set(k, v)
		}
	}

	doc.doc = append(doc.doc, bson.DocElem{Name: "$exists", Value: Builder().Set(k, v)})
	return doc
}

// Exists is equivalent to !expr0 || !expr1 || !expr...
func Exists(k string, v interface{}) *Document {
	return Builder().Exists(k, v)
}