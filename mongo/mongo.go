package mongo

import (
	"context"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database interface {
	Collection(string) Collection
	Client() Client
}

type Collection interface {
	InsertOne(context.Context, interface{}) (interface{}, error)
	InsertMany(context.Context, []interface{}) ([]interface{}, error)
	DeleteOne(context.Context, interface{}) (int64, error)
	CountDocuments(context.Context, interface{}, ...*options.CountOptions) (int64, error)
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

type Client interface {
	Database(string) Database
	Connect(context.Context) error
	Disconnect(context.Context) error
	StartSession() (mongo.Session, error)
	UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error
	Ping(context.Context) error
}

type mongoClient struct {
	cl *mongo.Client
}
type mongoDatabase struct {
	db *mongo.Database
}
type mongoCollection struct {
	coll *mongo.Collection
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

type mongoCursor struct {
	mc *mongo.Cursor
}

type mongoSession struct {
	mongo.Session
}

type nullawareDecoder struct {
	defDecoder bsoncodec.ValueDecoder
	zeroValue  reflect.Value
}

func NewClient(connection string) (Client, error) {
	time.Local = time.UTC
	c, err := mongo.NewClient(options.Client().ApplyURI(connection))
	return &mongoClient{cl: c}, err
}

func (mc *mongoClient) Ping(ctx context.Context) error {
	return mc.cl.Ping(ctx, readpref.Primary())
}

func (mc *mongoClient) Database(dbName string) Database {
	db := mc.cl.Database(dbName)
	return &mongoDatabase{db: db}
}

func (mc *mongoClient) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return mc.cl.UseSession(ctx, fn)
}

func (mc *mongoClient) StartSession() (mongo.Session, error) {
	session, err := mc.cl.StartSession()
	return &mongoSession{session}, err
}

func (mc *mongoClient) Connect(ctx context.Context) error {
	return mc.cl.Connect(ctx)
}

func (mc *mongoClient) Disconnect(ctx context.Context) error {
	return mc.cl.Disconnect(ctx)
}

func (md *mongoDatabase) Collection(colName string) Collection {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (md *mongoDatabase) Client() Client {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

func (mc *mongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return mc.coll.UpdateOne(ctx, filter, update, opts[:]...)
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	id, err := mc.coll.InsertOne(ctx, document)
	return id.InsertedID, err
}

func (mc *mongoCollection) InsertMany(ctx context.Context, document []interface{}) ([]interface{}, error) {
	res, err := mc.coll.InsertMany(ctx, document)
	return res.InsertedIDs, err
}

func (mc *mongoCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	count, err := mc.coll.DeleteOne(ctx, filter)
	return count.DeletedCount, err
}

func (mc *mongoCollection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return mc.coll.UpdateMany(ctx, filter, update, opts[:]...)
}

func (mc *mongoCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return mc.coll.CountDocuments(ctx, filter, opts...)
}
