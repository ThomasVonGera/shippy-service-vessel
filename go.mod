module github.com/ThomasVonGera/shippy-service-vessel

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/EwanValentine/shippy/shippy-service-vessel v0.0.0-20200612174527-3ad7cf4c07aa
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	go.mongodb.org/mongo-driver v1.4.0
	google.golang.org/protobuf v1.25.0
)
