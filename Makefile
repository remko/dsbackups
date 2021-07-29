PB_DIR = internal/pb
PB_FILES = $(PB_DIR)/dsbackups.pb.go
AEPB_DIR = internal/aepb
AEPB_FILES = $(AEPB_DIR)/datastore_v3.pb.go

PROTOC_FLAGS = --proto_path=proto --go_opt=paths=source_relative

.PHONY: all
all: bin

.PHONY: bin
bin: $(PB_FILES) $(AEPB_FILES)
	go build -o bin/ ./cmd/...

$(PB_FILES): proto/dsbackups.proto
	mkdir -p $(PB_DIR)
	protoc $(PROTOC_FLAGS) --go_out=$(PB_DIR) $<

$(AEPB_FILES): proto/datastore_v3.proto
	mkdir -p $(AEPB_DIR)
	protoc --go_opt=Mdatastore_v3.proto=github.com/remko/dsbackups/internal/aepb $(PROTOC_FLAGS) --go_out=$(AEPB_DIR) $<

run-datastore-emulator:
	gcloud beta emulators datastore start --data-dir=tmp/datastore

run-dsadmin:
	 npx dsadmin --project=dsbackups-dev