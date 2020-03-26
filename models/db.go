package models

import "gopkg.in/couchbase/gocb.v1"

var globalCluster *gocb.Cluster = nil
var globalBucket *gocb.Bucket = nil

func InitDB(COUCHBASE_HOST string, COUCHBASE_USERNAME string, COUCHBASE_PASSWORD string) {
	var err error

	globalCluster, err = gocb.Connect(COUCHBASE_HOST)
	if err != nil {
		panic(err)
	}

	globalCluster.Authenticate(gocb.PasswordAuthenticator{
		Username: COUCHBASE_USERNAME,
		Password: COUCHBASE_PASSWORD,
	})

	globalBucket, err = globalCluster.OpenBucket("tires", "")
	if err != nil {
		panic(err)
	}
}
