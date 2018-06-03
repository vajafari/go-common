package gocommon

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PagingData this struct hold information abount sorting
type PagingData struct {
	SortInfo   SortingData
	PageNumber int
	PageCount  int
}

// SortingData this struct hold information abount sorting data
type SortingData struct {
	FieldName string
	sortOrder bool
}

// MongoSession is currently a Mongo session.
type MongoSession struct {
	*mgo.Session
}

// DB shadows *mgo.DB to returns a IDataLayer interface instead of *mgo.Database.
func (s *MongoSession) DB(name string) IDatabase {
	return &MongoDatabase{Database: s.Session.DB(name)}
}

// Close shadows closeing mongo db session
func (s *MongoSession) Close() {
	s.Session.Close()
}

// Clone clone session
func (s *MongoSession) Clone() ISession {
	//return session.clone
	return &MongoSession{Session: s.Session.Clone()}
	//return s.Session.Clone()
}

// Copy copy session
func (s *MongoSession) Copy() ISession {
	return &MongoSession{Session: s.Session.Copy()}
}

// MongoDatabase wraps a mgo.Database to embed methods in models.
type MongoDatabase struct {
	*mgo.Database
}

// C shadows *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
func (d *MongoDatabase) C(name string) ICollection {
	return &MongoCollection{Collection: d.Database.C(name)}
}

// MongoCollection wraps a mgo.Collection to embed methods in models.
type MongoCollection struct {
	*mgo.Collection
}

// GetMyDocuments returns all scores.
func (c *MongoCollection) GetMyDocuments() ([]interface{}, error) {
	var documents []interface{}
	err := c.Find(bson.M{}).All(&documents)
	if err != nil {
		return nil, err
	}
	return documents, nil
}

// UpsertID is a convenience helper equivalent to:
func (c *MongoCollection) UpsertID(id interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	return c.Collection.UpsertId(id, update)
}

// Find shadows *mgo.Collection to returns a Query interface instead of *mgo.Query.
func (c *MongoCollection) Find(query interface{}) IQuery {
	return &MongoQuery{Query: c.Collection.Find(query)}
}

// FindID shadows *mgo.Collection to returns a Query interface instead of *mgo.Query.
func (c *MongoCollection) FindID(id interface{}) IQuery {
	return &MongoQuery{Query: c.Collection.FindId(id)}
}

// Count Find shadows *mgo.Collection to returns a Query interface instead of *mgo.Query.
func (c *MongoCollection) Count() (n int, err error) {
	return c.Collection.Count()
}

// Insert Find shadows *mgo.Collection to returns a Query interface instead of *mgo.Query.
func (c *MongoCollection) Insert(docs ...interface{}) error {
	return c.Collection.Insert(docs)
}

// Remove Find shadows *mgo.Collection to returns a Query interface instead of *mgo.Query.
func (c *MongoCollection) Remove(selector interface{}) error {
	return c.Collection.Remove(selector)
}

// Update Find shadows *mgo.Collection to returns a Query interface instead of *mgo.Query.
func (c *MongoCollection) Update(selector interface{}, update interface{}) error {
	return c.Collection.Update(selector, update)
}

// MongoQuery Find shadows *mgo.Collection to returns a Query interface instead of *mgo.Query.
type MongoQuery struct {
	*mgo.Query
}

// Sort shadows *mgo.Collection to returns a Query interface instead of *mgo.Query.
func (c *MongoQuery) Sort(fields ...string) IQuery {
	return &MongoQuery{Query: c.Sort(fields)}
}

// All Find shadows *mgo.Collection to returns a Query interface instead of *mgo.Query.
func (c *MongoQuery) All(result interface{}) error {
	return c.Query.All(result)
}

// One Find shadows *mgo.Collection to returns a Query interface instead of *mgo.Query.
func (c *MongoQuery) One(result interface{}) error {
	return c.Query.One(result)
}

// Distinct Find shadows *mgo.Collection to returns a Query interface instead of *mgo.Query.
func (c *MongoQuery) Distinct(key string, result interface{}) error {
	return c.Query.Distinct(key, result)
}
