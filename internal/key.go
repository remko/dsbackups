package internal

import (
	"fmt"
	"strings"

	pb "google.golang.org/genproto/googleapis/datastore/v1"
)

func KeyToString(key *pb.Key) string {
	ks := []string{}
	if key.GetPartitionId().GetProjectId() != "" {
		ks = append(ks, fmt.Sprintf("PROJECT('%s')", key.GetPartitionId().GetProjectId()))
	}
	if key.GetPartitionId().GetNamespaceId() != "" {
		ks = append(ks, fmt.Sprintf("NAMESPACE('%s')", key.GetPartitionId().GetNamespaceId()))
	}
	for _, p := range key.GetPath() {
		ks = append(ks, p.Kind)
		if name := p.GetName(); name != "" {
			ks = append(ks, fmt.Sprintf("\"%s\"", name))
		} else {
			ks = append(ks, fmt.Sprintf("%d", p.GetId()))
		}
	}

	return fmt.Sprintf("key(%s)", strings.Join(ks, ", "))
}
