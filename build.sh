repository=`cat go.mod | grep -E "^module\s[0-9a-zA-Z\./_\-]+" | awk '{print $2}'`
docker build --build-arg repository=$repository -v /go/pkg/mod:/go/pkg/mod . -t image --file Dockerfile
