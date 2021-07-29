package dsbackups

import (
	"io"
	"io/ioutil"
	"time"

	"github.com/golang/leveldb/record"
	pb "github.com/remko/dsbackups/internal/pb"
	"google.golang.org/protobuf/proto"
)

type OverallExportMetadata struct {
	Exports []ExportInfo
}

type ExportInfo struct {
	Kind  string
	Path  string
	Size  int
	Count int
}

// Reads a `.overall_export_metadata` file
func ReadOverallExportMetadata(r io.Reader) (OverallExportMetadata, error) {
	result := OverallExportMetadata{Exports: []ExportInfo{}}
	rr := record.NewReader(r)

	_, err := rr.Next()
	if err != nil {
		return result, err
	}
	// Fixed version number? Skip.

	rec, err := rr.Next()
	if err != nil {
		return result, err
	}
	s, err := ioutil.ReadAll(rec)
	if err != nil {
		return result, err
	}

	ompb := pb.OverallExportMetadata{}
	if err := proto.Unmarshal(s, &ompb); err != nil {
		return result, err
	}
	for _, export := range ompb.GetExports() {
		result.Exports = append(result.Exports, ExportInfo{
			Kind:  export.GetKind().GetKind(),
			Path:  export.GetPath(),
			Size:  int(export.GetSize()),
			Count: int(export.GetCount()),
		})
	}

	return result, nil
}

//////////////////////////////////////////////////////////////////////////////

type ExportMetadata struct {
	StartTime time.Time
	EndTime   time.Time
	Outputs   []string
}

// Reads a `.export_metadata` file
func ReadExportMetadata(r io.Reader) (ExportMetadata, error) {
	s, err := ioutil.ReadAll(r)
	if err != nil {
		return ExportMetadata{}, err
	}

	empb := pb.ExportMetadata{}
	if err := proto.Unmarshal(s, &empb); err != nil {
		return ExportMetadata{}, err
	}
	result := ExportMetadata{
		StartTime: unixMicroToTime(empb.GetOperation().GetStartTime()),
		EndTime:   unixMicroToTime(empb.GetOperation().GetEndTime()),
		Outputs:   make([]string, 0, len(empb.GetItems().GetOutputs())),
	}
	result.Outputs = append(result.Outputs, empb.GetItems().GetOutputs()...)

	return result, nil
}

func unixMicroToTime(t int64) time.Time {
	return time.Unix(t/(1000*1000), t%(1000*1000))
}
