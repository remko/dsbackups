package main

import (
	"context"
	"io"
	"net/url"
	"os"
	"path"
	"strings"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/gcsblob"
)

type blobReader struct {
	Bucket *blob.Bucket
	Reader *blob.Reader
}

// Small wrapper around Cloud CDK for easier reading from buckets
func openBlob(path string) (io.ReadCloser, error) {
	ctx := context.Background()
	if strings.HasPrefix(path, "gs://") {
		uri, err := url.Parse(path)
		if err != nil {
			return nil, err
		}
		object := uri.Path[1:]
		uri.Path = ""
		bucket, err := blob.OpenBucket(ctx, uri.String())
		if err != nil {
			return nil, err
		}
		r, err := bucket.NewReader(ctx, object, nil)
		if err != nil {
			bucket.Close()
			return nil, err
		}
		return &blobReader{bucket, r}, nil
	}
	return os.Open(path)
}

func (br *blobReader) Read(p []byte) (int, error) {
	return br.Reader.Read(p)
}

func (br *blobReader) Close() error {
	br.Reader.Close()
	return br.Bucket.Close()
}

func relativePath(base string, target string) string {
	if strings.HasPrefix(base, "gs://") {
		uri, err := url.Parse(base)
		if err != nil {
			panic(err)
		}
		uri.Path = path.Join(path.Dir(uri.Path), target)
		return uri.String()
	} else {
		return path.Join(path.Dir(base), target)
	}
}
