package main

import (
	"context"
	"io"
	"log"

	"github.com/remko/dsbackups"
	"github.com/remko/dsbackups/internal"
	pb "google.golang.org/genproto/googleapis/datastore/v1"
)

type Importer interface {
	Close() error
	Import(ctx context.Context, output io.Reader, predicate func(entity *pb.Entity) bool) error
}

type DryRunImporter struct {
	Count int
}

func (i *DryRunImporter) Close() error {
	log.Printf("would have imported %d entities", i.Count)
	return nil
}

func (i *DryRunImporter) Import(ctx context.Context, output io.Reader, predicate func(entity *pb.Entity) bool) error {
	records := dsbackups.NewOutputReader(output)
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
		log.Printf("[dry-run] importing %s", internal.KeyToString(e.Key))
		i.Count += 1
	}
	return nil
}
