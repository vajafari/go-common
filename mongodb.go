package gocommon

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// OperatorEnum define operators for filter database
type OperatorEnum int

// PageInfo this struct hold information abount sorting
type PageInfo struct {
	PageNumber    int
	RecordPerPage int
}

// SortInfo this struct hold information abount sorting data
type SortInfo struct {
	FieldName string
	SortOrder int
}

// PagingData select specific page of data
type PagingData struct {
	Page   *PageInfo
	Sort   []*SortInfo
	Filter interface{}
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

// SelectPage shadows *mgo.Collection to returns a Query interface instead of *mgo.Query.
func (c *MongoCollection) SelectPage(pagingData PagingData) IQuery {

	sortSlc := make([]string, 0)
	if len(pagingData.Sort) > 0 {
		for _, item := range pagingData.Sort {
			if item != nil && len(item.FieldName) > 0 {
				if item.SortOrder > 0 {
					// Asc
					sortSlc = append(sortSlc, item.FieldName)

				} else {
					// Desc
					sortSlc = append(sortSlc, "-"+item.FieldName)
				}
			}
		}
	}

	if pagingData.Page != nil && len(sortSlc) > 0 {
		return &MongoQuery{Query: c.Collection.Find(pagingData.Filter).Skip((pagingData.Page.PageNumber - 1) * pagingData.Page.RecordPerPage).Limit(pagingData.Page.RecordPerPage).Sort(sortSlc...)}
	} else if pagingData.Page != nil && len(sortSlc) == 0 {
		return &MongoQuery{Query: c.Collection.Find(pagingData.Filter).Skip((pagingData.Page.PageNumber - 1) * pagingData.Page.RecordPerPage).Limit(pagingData.Page.RecordPerPage)}
	} else if pagingData.Page == nil && len(sortSlc) > 0 {
		return &MongoQuery{Query: c.Collection.Find(pagingData.Filter).Sort(sortSlc...)}
	}

	return &MongoQuery{Query: c.Collection.Find(pagingData.Filter)}

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
	return &MongoQuery{Query: c.Query.Sort(fields...)}
}

// Skip skips over the n initial documents from the query results.  Note that
// this only makes sense with capped collections where documents are naturally
// ordered by insertion time, or with sorted results.
func (c *MongoQuery) Skip(n int) IQuery {
	return &MongoQuery{Query: c.Query.Skip(n)}
}

// Select select subset of result
func (c *MongoQuery) Select(selector interface{}) IQuery {
	return &MongoQuery{Query: c.Query.Select(selector)}
}

// Limit restricts the maximum number of documents retrieved to n, and also
// changes the batch size to the same value.  Once n documents have been
// returned by Next, the following call will return ErrNotFound.
func (c *MongoQuery) Limit(n int) IQuery {
	return &MongoQuery{Query: c.Query.Limit(n)}
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
