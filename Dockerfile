FROM golang:1.10-alpine3.7
COPY . /hackerrank
WORKDIR /hackerrank
ENV PATH "${PWD}/bin:${PATH}"
