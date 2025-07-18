all:
	go build ./cmd/csv-benchmark
	go run ./cmd/csv-benchmark

bench:
	go test ./... -bench=. -benchmem -count 10 -run=^#
