# Google Cloud Datastore Export Structure & Format

This file contains a description of the structure and formats used in 
[Google Cloud Datastore Exports](https://cloud.google.com/datastore/docs/export-import-entities).

⚠️ Apart from the format of the actual stored entities, all of this is found through guessing
and reverse engineering. None of this information is officially documented or supported by
Google, so proceed with caution.

## Directory Structure 

A datastore export into a `myexport` folder consists of the following directory structure:

    myexport.overall_export_meta
    default_namespace/
      kind_MyKind/
        default_namespace_kind_MyKind.export_metadata
        output-0
        output-1
        output-4
        ...
      kind_MyOtherKind/
        output-2
        output-6
        ...
      ...

If you exported all your kinds in one operation, the structure can be similar but simpler:

    myexport.overall_export_meta
    all_namespaces/
        all_kinds/
            all_namespaces_all_kinds.export_metadata
            output-0
            output-1
            output-4
            ...

An export contains the following files:
- The `.overall_export_meta` file lists all the sub-exports (by kind) that were exported in this export, the paths at which their `.export_metadata` can be found, and some metadata about each sub-export (count and total size). If you did not export specific kinds, there will only be one sub-export.
- The `*.export_metadata` files contain the list of files created for each sub-export, together with some metadata (start time / end time of the sub-export)
- The `output-*` files contain the actual exported data.

## `.overall_metadata` format

This file contains a list of all sub-exports, and pointers to the `.export_metadata` files containing metadata of each sub-export.

This file is in [LevelDB log format](https://github.com/google/leveldb/blob/master/doc/log_format.md), containing 2 records.

- Record 1 contains fixed version information (?)

- Record 2 contains a [`OverallExportMetadata` protocol buffer message](https://github.com/remko/dsbackups/blob/master/proto/dsbackups.proto)

## `.export_metadata` format

Every sub-export has its list of output files stored in the `.export_metadata` file.

This file is a raw protocol buffer message, described by the [`ExportMetadata` protocol buffer message](https://github.com/remko/dsbackups/blob/master/proto/dsbackups.proto).


## `output-` format

These `output-*` files contain the actual exported entities. 

The toplevel structure of this file is in [LevelDB log format](https://github.com/google/leveldb/blob/master/doc/log_format.md).<sup>[ref](https://stackoverflow.com/a/44831736/675011)</sup>

Each LevelDB record contains an entity, serialized as the `EntityProto` structure from the  [AppEngine Datastore V3 Protocol Buffer definition](https://github.com/golang/appengine/blob/master/internal/datastore/datastore_v3.proto#L111).
