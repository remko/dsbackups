package dsbackups

import (
	"io"
	"io/ioutil"

	"github.com/golang/leveldb/record"
	pb "google.golang.org/genproto/googleapis/datastore/v1"

	"github.com/remko/dsbackups/internal"
)

// A reader for an exported `.output` file
// Converts entities from the internal (AppEngine) protobuf format to the
// Datastore V1 protobuf format (https://cloud.google.com/datastore/docs/reference/data/rpc/google.datastore.v1#google.datastore.v1.Entity)
type OutputReader struct {
	r *record.Reader
}

func NewOutputReader(r io.Reader) *OutputReader {
	return &OutputReader{
		r: record.NewReader(r),
	}
}

func (br *OutputReader) Next() (*pb.Entity, error) {
	rec, err := br.r.Next()
	if err != nil {
		return nil, err
	}
	s, err := ioutil.ReadAll(rec)
	if err != nil {
		return nil, err
	}
	return internal.UnmarshalAEEntity(s)
}
