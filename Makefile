bindir=bin

all:
	@mkdir -p $(bindir)
	@go build -o $(bindir)/dax-examples ./examples
	@go build -o $(bindir)/dax-mixer ./cmd/mixer

check:
	@gometalinter --tests --vendor --disable-all \
		--enable=gofmt \
		--enable=vet \
		./...
