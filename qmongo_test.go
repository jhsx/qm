package qm

import (
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func tt() interface{} {
	return bson.M{"$and": []bson.M{
		{"id": bson.M{"$eq": "p1"}},
		{"name": bson.M{"$eq": "p2"}},
		{"description": bson.M{"$regex": "asfasdfas"}},
	}}
}

func ttd() interface{} {
	return bson.D{
		{"$and", []bson.D{
			{{"id", bson.D{{"$eq", "p1"}}}},
			{{"name", bson.D{{"$eq", "p2"}}}},
			{{"description", bson.D{{"$regex", "asfasdfas"}}}},
		}},
	}
}

func ttdb() interface{} {
	return And(
		Eq("id", "p1"),
		Eq("name", "p2"),
		Eq("description", "p3"),
	)
}

func BenchmarkDBMapObjects(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tt()
	}
}

func BenchmarkDBRawObjects(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ttd()
	}
}

func BenchmarkDBDocBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ttdb()
	}
}
