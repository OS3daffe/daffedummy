FROM golang
 
ADD . ~/Grid/go-workspace/src/unhosted/daffedummy
RUN go install ~/Grid/go-workspace/src/unhosted/daffedummy
ENTRYPOINT ~/Grid/go-workspace/bin/daffedummy
 
EXPOSE 8080