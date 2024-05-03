APP_BIN = app/build/bin
PROGNAME= v-mailer
PROGNAME_WIN= $(PROGNAME).exe
DESTDIR=
VERSION=0.0.1
BUILD_OPTIONS = -ldflags "-X main.Version=$(VERSION)"
BINDIR= /usr/local/bin
MANPAGE= docs/$(PROGNAME).1
MANDIR= /usr/local/share/man/man1
CP= /bin/cp -v
OS= $(shell go env GOOS)

all:
	@echo "- Compiling ${PROGNAME} ..."
	go build

install: install-bin

install-bin:
	$(CP) $(APP_BIN) $(DESTDIR)$(BINDIR)
	$(CP) $(MANPAGE) $(DESTDIR)$(MANDIR)

example:
	@./scripts/mkexamples.sh

$(PROGNAME) : native

# to strip out debug info: go build -ldflags="-s -w"

linux:
	@echo "- Building $(PROGNAME)_linux"
	@CGO_ENABLED=0 GOOS=linux go build -o $(PROGNAME)_linux
	/bin/ls -lt $(PROGNAME)_linux
	@echo ""

native:
	@echo "- Building $(PROGNAME)"
	@go build -o $(PROGNAME)
	/bin/ls -lt $(PROGNAME)
	@echo ""

windows:
	@echo "- Building $(PROGNAME_WIN) for Windows amd64"
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o $(PROGNAME_WIN)
	/bin/ls -lt $(PROGNAME_WIN)
	@echo ""

lint:
	golangci-lint run

build: clean $(APP_BIN)

$(APP_BIN)
	go build -o $(APP_BIN) ./app/cmd/v-mailer.go

clean:
#	/bin/rm -rf $(APP_BIN)$(PROGNAME) $(APP_BIN)$(PROGNAME_WIN) $(APP_BIN)$(PROGNAME)_$(OS) $(APP_BIN)$(PROGNAME)_* *.bak $(APP_BIN)$(PROGNAME)_linux
	rm -rf ./app/build || true

# generate files/examples.txt from docs/examples.md
# generate examples.go from examples.txt for -ex flag
gen: example doc

dev: gen all doc

doc:
	@./scripts/mkdocs.sh
	@echo " - Generate docs/v-mailsend.1"
	@pandoc --standalone --to man README.md -o docs/v-mailsend.1

swagger:
	swag init -g ./app/cmd/v-mailer.go -o ./app/docs

help:
	@echo "============================================================"
	@echo " make gen   - assemble document, create usage.txt and examples.


	go"
	@echo " make       - build native client"
	@echo " make linux - build linux client"
	@echo " make win   - build windows client"
	@echo " make clean"
	@echo "============================================================"