#********************with golang docker image & direct run command*************
FROM golang:1.14.1-alpine3.11

WORKDIR /go/src/github.com/Emon331046/library-management-api

ADD . /go/src/github.com/Emon331046/library-management-api

RUN go mod tidy
RUN go mod vendor
RUN go build -o main .
# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]


#*********************copy the build in binary store in my working directory**************

#
#FROM ubuntu
#
#WORKDIR /go/src/github.com/Emon331046/library-management-api
#
#COPY main .
#
##when u are in directory only main didn't work
#CMD ["./main"]
#
#EXPOSE 8000
#
##
###*********************copy the build in binary store in bin**************
#FROM ubuntu
#
## in my local directory i have created a bin folder
## then created my binary named main in that folder
## here i am copying my local main in the /bin folder of the system
#COPY /bin/main /bin
#
#
#EXPOSE 8000
#
##when u are in bin folder $ ./main didn't work
#CMD ["main"]
#

#------------------------Using Ubuntu image---------------------------------------------------
#FROM ubuntu:18.10
#
#COPY main /bin/go-api-server
#
#COPY bin/main /bin/go-api-server
#CMD ["start","-b"]
#ENTRYPOINT ["/bin/go-api-server"]