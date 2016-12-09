PKG=github.com/mmcloughlin/cite
CMD=${PKG}/cmd/cite
GITSHA=`git rev-parse HEAD`
LDFLAGS="-X ${CMD}/cmd.gitSHA=${GITSHA}"

install:
	go install -ldflags ${LDFLAGS} ${CMD}
