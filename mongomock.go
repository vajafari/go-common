package gocommon

// import (
// 	mgo "gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"
// )

// // ISession is an interface to access to the Session struct.
// type ISession interface {
// 	DB(name string) IDataLayer
// 	Close()
// 	Clone() ISession
// 	Copy() ISession
// }

// // MongoSession is currently a Mongo session.
// type MongoSession struct {
// 	dbSession *mgo.Session
// }

// // DB shadows *mgo.DB to returns a IDataLayer interface instead of *mgo.Database.
// func (s MongoSession) DB(name string) IDataLayer {
// 	return &MongoDatabase{Database: s.dbSession.DB(name)}
// }

// // Clone clone session
// func (s MongoSession) Clone() ISession {
// 	//return session.clone
// 	return s.dbSession.Clone()
// }

// // Copy copy session
// func (s MongoSession) Copy() ISession {
// 	return s.dbSession.Copy()
// }

// // IDataLayer is an interface to access to the database struct.
// type IDataLayer interface {
// 	C(name string) ICollection
// }

// // MongoCollection wraps a mgo.Collection to embed methods in models.
// type MongoCollection struct {
// 	*mgo.Collection
// }

// // MongoQuery wraps a mgo.Query to embed methods in models.
// type MongoQuery struct {
// 	*mgo.Query
// }

// // ICollection is an interface to access to the collection struct.
// type ICollection interface {
// 	Find(query interface{}) Query
// 	Count() (n int, err error)
// 	Insert(docs ...interface{}) error
// 	Remove(selector interface{}) error
// 	Update(selector interface{}, update interface{}) error
// 	GetMyDocuments() ([]interface{}, error)
// }

// // MongoDatabase wraps a mgo.Database to embed methods in models.
// type MongoDatabase struct {
// 	*mgo.Database
// }

// // C shadows *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
// func (d MongoDatabase) C(name string) ICollection {
// 	return &MongoCollection{Collection: d.Database.C(name)}
// }

// // GetMyDocuments returns all scores.
// func (c *MongoCollection) GetMyDocuments() ([]interface{}, error) {
// 	var documents []interface{}
// 	err := c.Find(bson.M{}).All(&documents)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return documents, nil
// }

// // Find shadows *mgo.Collection to returns a Query interface instead of *mgo.Query.
// func (c MongoCollection) Find(query interface{}) Query {
// 	return MongoQuery{Query: c.Collection.Find(query)}
// }

// // Query is an interface to access to the database struct
// type Query interface {
// 	All(result interface{}) error
// 	One(result interface{}) (err error)
// 	Distinct(key string, result interface{}) error
// }
