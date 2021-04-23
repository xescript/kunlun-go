RM=npx rimraf --glob

.PHONY: all test clean

all: bin/kunlun-migrator

test:
	go test -v \
		./cmd/... \
		./internal/... \
		./ent/schema/...

clean:
	$(RM) 'ent/!(generate).go'
	$(RM) 'ent/!(schema)/**'
	$(RM) 'bin'
	$(RM) './**/*.sqlite3'

bin/kunlun-migrator:
	go build -o $@ ./cmd/migrator
