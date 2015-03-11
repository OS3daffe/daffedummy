FROM golang
 
ADD . /go/src/unhosted/daffedummy
RUN go install /go/src/unhosted/daffedummy
ENTRYPOINT /go/bin/daffedummy
 
EXPOSE 8080