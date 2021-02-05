Project := git-syncer
BuildDist = dist/build
ReleaseDist = dist/release

Platforms = linux darwin windows
GOOS = $(shell go env GOOS)

Clis := $(foreach n,$(shell go list ./cli/*),$(notdir $(n)))

Version := $(shell git describe --tags --dirty --match="v*" 2> /dev/null || echo v0.0.0-dev)
Date := $(shell date -u '+%Y-%m-%d-%H%M UTC')

CliReleaseDistribution = $(foreach p,$(Platforms),$(foreach c,$(Clis),$(ReleaseDist)/$(Version)/$(c)-$(p)-amd64))
CliBuildDistribution = $(foreach c,$(Clis),$(BuildDist)/$(c))

go-clean:
	go clean ./cmd/...
.PHONY: go-clean

go-get:
	go get
	go mod download
.PHONY: go-get

clean: go-clean
	rm -rf $(ReleaseDist)/$(Version)
.PHONY: clean

$(CliReleaseDistribution):
	$(call go_build,$@)
.PHONY: $(CliReleaseDistribution)

$(CliBuildDistribution):
	$(call go_build,$@)
.PHONY: $(CliBuildDistribution)

build: $(CliBuildDistribution)
.PHONY: build

release: $(CliReleaseDistribution)
.PHONY: release

define go_build
	@-rm $1
	$(eval buildPlatform = $(shell $(foreach p,$(Platforms),echo $1 | grep -owh $(p);)))
	$(eval buildCmd := $(shell $(foreach c,$(Clis),echo $1 | grep -owh $(c);)))
	$(eval buildGOOS := $(if $(buildPlatform),$(buildPlatform),$(GOOS)))
	CGO_ENABLED=0 GOOS=$(buildGOOS) GOARCH=amd64 \
		go build \
			-gcflags="all=-N -l" \
			-ldflags="-X 'github.com/funnyecho/git-syncer/command/command.BuildPlatform=$(buildGOOS)-amd64' -X 'github.com/funnyecho/git-syncer/command/command.Version=$(Version)' -X 'github.com/funnyecho/git-syncer/command/command.BuildTime=$(Date)'" \
			-o ./$1 \
			./cli/$(buildCmd)/main.go;
endef

test:
	go test github.com/funnyecho/git-syncer/...

.PHONY:
	test