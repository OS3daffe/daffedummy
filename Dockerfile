#base image with golang installed on /go
FROM golang 

#add our repo to the $GOPATH of the image
ADD . /go/src/github.com/OS3daffe/daffedummy 

#download all the dependencies inside the container
RUN go get github.com/OS3daffe/daffedummy/... 

#compile and install the service
RUN go install github.com/OS3daffe/daffedummy

#export the container
ENTRYPOINT /go/bin/daffedummy 

#container listens to 8080 internal
EXPOSE 8080 