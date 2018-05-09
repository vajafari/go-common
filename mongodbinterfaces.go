package gocommon

import mgo "gopkg.in/mgo.v2"

// ISession is an interface to access to the Session struct.
type ISession interface {
	DB(name string) IDatabase
	Close()
	Clone() ISession
	Copy() ISession
}

// IDatabase is an interface to access to the database struct.
type IDatabase interface {
	C(name string) ICollection
}

// ICollection is an interface to access to the collection struct.
type ICollection interface {
	Find(query interface{}) IQuery
	Count() (n int, err error)
	Insert(docs ...interface{}) error
	Remove(selector interface{}) error
	Update(selector interface{}, update interface{}) error
	GetMyDocuments() ([]interface{}, error)
	FindID(id interface{}) IQuery
	UpsertID(interface{}, interface{}) (*mgo.ChangeInfo, error)
}

// IQuery is an interface to access to the database struct
type IQuery interface {
	All(result interface{}) error
	One(result interface{}) (err error)
	Distinct(key string, result interface{}) error
}
