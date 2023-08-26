FROM golang:1.21
LABEL authors="lubov"

ENTRYPOINT ["top", "-b"]