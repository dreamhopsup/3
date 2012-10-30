all: githook 6g 

PREFIX=nimble-cube

PKGS=\
	$(PREFIX)/gpu/conv\
	$(PREFIX)/gpu\
	$(PREFIX)/mag\
	$(PREFIX)/dump\
	$(PREFIX)/core\

6g: ptx
	go install -v $(PKGS)
	go install -v 

GCCGO=gccgo -gccgoflags '-static-libgcc -O4 -Ofast -march=native'

gccgo: ptx
	go install -v -compiler $(GCCGO) $(PKGS)
	go install -v -compiler $(GCCGO)

ptx:
	make -C gpu/ptx

githook:
	ln -sf $(CURDIR)/pre-commit .git/hooks/pre-commit
	ln -sf $(CURDIR)/post-commit .git/hooks/post-commit

test: 6gtest gccgotest unittest

unittest:
	make -C test

6gtest: 6g
	go test -i $(PKGS) 

gccgotest: gccgo
	go test -i -compiler=$(gccgo) $(PKGS)

.PHONY: clean
clean:
	go clean -i -x $(PKGS)
	rm -rf $(GOPATH)/pkg/gccgo/$(PREFIX)/
