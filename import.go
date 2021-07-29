package dsbackups

import (
	"context"
	"io"
	"log"

	"github.com/remko/dsbackups/internal"
	pb "google.golang.org/genproto/googleapis/datastore/v1"
	"google.golang.org/grpc"
)

type Importer struct {
	projectID string
	conn      *grpc.ClientConn
	client    pb.DatastoreClient
}

func NewImporter(server string, projectID string) (*Importer, error) {
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pb.NewDatastoreClient(conn)
	return &Importer{projectID, conn, client}, nil
}

func (i *Importer) Close() error {
	return i.conn.Close()
}

// Imports kinds from a single export output file
func (i *Importer) Import(ctx context.Context, output io.Reader, predicate func(entity *pb.Entity) bool) error {
	records := NewOutputReader(output)
	for {
		e, err := records.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		if !predicate(e) {
			continue
		}
		log.Printf("importing %s", internal.KeyToString(e.Key))
		_, err = i.client.ReserveIds(ctx, &pb.ReserveIdsRequest{
			ProjectId: i.projectID,
			Keys:      []*pb.Key{e.Key},
		})
		if err != nil {
			return err
		}
		_, err = i.client.Commit(ctx, &pb.CommitRequest{
			ProjectId: i.projectID,
			Mode:      pb.CommitRequest_NON_TRANSACTIONAL,
			Mutations: []*pb.Mutation{
				{Operation: &pb.Mutation_Insert{Insert: e}},
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
