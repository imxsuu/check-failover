FROM golang:1.17.5
LABEL maintainer="imxsuu@gmail.com"

RUN apt-get update -y 
RUN apt-get install jq -y
COPY . .

EXPOSE 9010
CMD ["./server"]
