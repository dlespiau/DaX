bindir=bin

all:
	@mkdir -p $(bindir)
	@go build -o $(bindir)/dax-examples ./examples

check:
	@gometalinter --tests --vendor --disable-all \
		--enable=gofmt \
		--enable=vet \
		./...
