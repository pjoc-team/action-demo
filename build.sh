export REPOSITORY=`cat go.mod | grep -E "^module\s[0-9a-zA-Z\./_\-]+" | awk '{print $2}'`
export NAME=`basename $REPOSITORY`
docker build --build-arg repository=$REPOSITORY . -t image --file Dockerfile
