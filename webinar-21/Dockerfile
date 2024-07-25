FROM golang

WORKDIR /api

COPY . .

WORKDIR /api/cmd/serve

RUN go build -o api .

CMD [ "./api" ]
