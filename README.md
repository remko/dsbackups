# `dsbackups` [![Go Reference](https://pkg.go.dev/badge/github.com/remko/dsbackups.svg)](https://pkg.go.dev/github.com/remko/dsbackups)

This package contains code to work with [Google Cloud Datastore exports](https://cloud.google.com/datastore/docs/export-import-entities).
(e.g. to restore a subset of entities from a backup, which is not possible using 'official' interfaces)

**DISCLAIMERS**: 
- The format of Datastore exports is undocumented. This package is based on reverse engineering and the little
information known about the datastore export format. None of this is officially supported by Google.
  Since all this is unofficial and undocumented, this may break at some point (although I doubt it will any time soon, as it hasn't for years).
- This package exists mostly as proof-of-concept and documentation, based on my own usage of datastore. 
  Don't assume this code works and doesn't delete all your data.  
  The safest method of using this in production is probably to use this code to import data into a fresh local datastore emulator, verify the imported data
  (e.g. using [DSAdmin](https://github.com/remko/dsadmin)), [export the data from your emulator](https://cloud.google.com/datastore/docs/tools/emulator-export-import),
  and import your export [using the official methods](https://cloud.google.com/datastore/docs/export-import-entities).


## Documentation

- [Go reference](https://pkg.go.dev/github.com/remko/dsbackups).
- See the commands in `cmd/dsbackuputil/dsbackuputil.go` for examples using this package.
- [Google Cloud Datastore Export Format description](https://github.com/remko/dsbackups/blob/master/doc/DatastoreExportFormat.md)

## `dsbackuputil`

The `dsbackuputil` command contains utility commands to inspect or import backups.
You can use it as is for quick backup inspections, or as a basis for more advanced commands.

See `dsbackuputil help` for a full list of commands.


Examples: 

- List all entities in an export (on GCS):

      export GOOGLE_APPLICATION_CREDENTIALS=credentials.json 
      dsbackuputil list \
        gs://my-bucket/mybackup/mybackup.overall_export_metadata

- List all sub-exports in an export (on GCS):

      export GOOGLE_APPLICATION_CREDENTIALS=credentials.json 
      dsbackuputil describe-exports \
        gs://my-bucket/mybackup/mybackup.overall_export_metadata

  Output:

      MyKind: default_namespace/kind_MyKind/default_namespace_kind_MyKind.export_metadata (3308 items, 808448b)
      MyOtherKind: default_namespace/kind_MyOtherKind/default_namespace_kind_MyOtherKind.export_metadata (360 items, 5944254b)

- List all output files of a single sub-export in an export:

      export GOOGLE_APPLICATION_CREDENTIALS=credentials.json 
      dsbackuputil describe-export \
        gs://my-bucket/mybackup/default_namespace/kind_MyKind/default_namespace_kind_MyKind.export_metadata 

  Output: 
  
      Start: 2021-07-29 15:00:03.000612651 +0200 CEST
      End: 2021-07-29 15:02:55.000135493 +0200 CEST
      Outputs:
        output-0
        output-1
        output-2
        output-3

- Import a kind from an export on GCS into a local datastore emulator:

      $ export GOOGLE_APPLICATION_CREDENTIALS=credentials.json 
      $ dsbackuputil import --project=my-project --kind=MyKind \
        gs://my-bucket/mybackup/mybackup.overall_export_metadata




