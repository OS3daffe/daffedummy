FROM golang #base image with golang installed on /go
 
ADD . /go/src/github.com/OS3daffe/daffedummy #add our repo to the $GOPATH of the image
RUN go get github.com/OS3daffe/daffedummy/... #download all the dependencies inside the container
RUN go install github.com/OS3daffe/daffedummy #compile and install the service
ENTRYPOINT /go/bin/daffedummy #export the container
 
EXPOSE 8080 #container listens to 8080 internal
