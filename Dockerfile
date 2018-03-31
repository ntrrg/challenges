FROM golang:1.10-alpine3.7
COPY . /hackerrank
WORKDIR /hackerrank
ENV PATH="/hackerrank/bin:${PATH}"
CMD ["hrrun.sh"]
