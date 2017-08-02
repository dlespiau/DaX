
check:
	@gometalinter --tests --vendor --disable-all \
		--enable=gofmt \
		--enable=vet \
		./...
