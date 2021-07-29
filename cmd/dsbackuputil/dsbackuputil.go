package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/remko/dsbackups"
	"github.com/remko/dsbackups/internal"
	"github.com/spf13/cobra"
	pb "google.golang.org/genproto/googleapis/datastore/v1"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dsbackuputil",
		Short: "A tool for reading Google Cloud Datastore Backups",
	}

	describeExportsCmd = &cobra.Command{
		Use:   "describe-exports [file].overall_export_metadata",
		Short: "List the exported kinds",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			f, err := openBlob(args[0])
			if err != nil {
				panic(err)
			}
			defer f.Close()
			om, err := dsbackups.ReadOverallExportMetadata(f)
			if err != nil {
				panic(err)
			}
			for _, export := range om.Exports {
				fmt.Printf("%s: %s (%d items, %db)\n", export.Kind, export.Path, export.Count, export.Size)
			}
		},
	}

	describeExportCmd = &cobra.Command{
		Use:   "describe-export [file].export_metadata",
		Short: "List the files of an export",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			f, err := openBlob(args[0])
			if err != nil {
				panic(err)
			}
			defer f.Close()
			em, err := dsbackups.ReadExportMetadata(f)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Start: %s\n", em.StartTime)
			fmt.Printf("End: %s\n", em.EndTime)
			fmt.Printf("Outputs:\n")
			for _, output := range em.Outputs {
				fmt.Printf("  %s\n", output)
			}
		},
	}

	describeOutputCmd = &cobra.Command{
		Use:   "describe-output output-[index]",
		Short: "List the entities in an output file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			f, err := openBlob(args[0])
			if err != nil {
				panic(err)
			}
			defer f.Close()
			r := dsbackups.NewOutputReader(f)
			for {
				e, err := r.Next()
				if err == io.EOF {
					break
				} else if err != nil {
					panic(err)
				}
				log.Printf("%s\n", internal.KeyToString(e.Key))
			}
		},
	}

	listCmd = &cobra.Command{
		Use:   "list [file].overall_export_metadata",
		Short: "List all entities",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			oemFile := args[0]
			var err error

			f, err := openBlob(oemFile)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			om, err := dsbackups.ReadOverallExportMetadata(f)
			if err != nil {
				panic(err)
			}
			for _, export := range om.Exports {
				emFile := relativePath(oemFile, export.Path)
				emf, err := openBlob(emFile)
				if err != nil {
					panic(err)
				}
				defer emf.Close()
				em, err := dsbackups.ReadExportMetadata(emf)
				if err != nil {
					panic(err)
				}
				for _, output := range em.Outputs {
					oFile := relativePath(emFile, output)
					of, err := openBlob(oFile)
					if err != nil {
						panic(err)
					}
					defer of.Close()

					ofr := dsbackups.NewOutputReader(of)
					for {
						e, err := ofr.Next()
						if err == io.EOF {
							break
						} else if err != nil {
							panic(err)
						}
						fmt.Printf("%s\n", internal.KeyToString(e.Key))
					}
				}
			}
		},
	}

	importCmd = &cobra.Command{
		Use:   "import [file].overall_export_metadata",
		Short: "Imports entities into datastore",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			oemFile := args[0]
			var importer Importer
			var err error
			if dryRun {
				importer = &DryRunImporter{}
			} else {
				importer, err = dsbackups.NewImporter(dsServer, dsProject)
				if err != nil {
					panic(err)
				}
			}
			defer importer.Close()

			f, err := openBlob(oemFile)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			om, err := dsbackups.ReadOverallExportMetadata(f)
			if err != nil {
				panic(err)
			}
			for _, export := range om.Exports {
				if export.Kind != "" && export.Kind != kind {
					continue
				}
				emFile := relativePath(oemFile, export.Path)
				log.Printf("opening export metadata file %s", emFile)
				emf, err := openBlob(emFile)
				if err != nil {
					panic(err)
				}
				defer emf.Close()
				em, err := dsbackups.ReadExportMetadata(emf)
				if err != nil {
					panic(err)
				}
				for _, output := range em.Outputs {
					oFile := relativePath(emFile, output)
					log.Printf("opening export output file %s", oFile)
					of, err := openBlob(oFile)
					if err != nil {
						panic(err)
					}
					defer of.Close()
					if err := importer.Import(context.Background(), of, func(e *pb.Entity) bool { return e.Key.Path[len(e.Key.Path)-1].Kind == kind }); err != nil {
						panic(err)
					}
				}
			}
		},
	}

	dsServer  string
	dsProject string
	kind      string
	dryRun    bool
)

func init() {
	importCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Only report number of entities that would be imported")
	importCmd.Flags().StringVar(&dsServer, "server", "localhost:8081", "Datastore server to import to")
	importCmd.Flags().StringVar(&dsProject, "project", "", "Datastore project to import to")
	importCmd.Flags().StringVar(&kind, "kind", "", "Entity kind to import")
	importCmd.MarkFlagRequired("project")
	importCmd.MarkFlagRequired("kind")

	rootCmd.AddCommand(describeExportsCmd)
	rootCmd.AddCommand(describeExportCmd)
	rootCmd.AddCommand(describeOutputCmd)
	rootCmd.AddCommand(importCmd)
	rootCmd.AddCommand(listCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
