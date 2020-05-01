protoc -I ../../services/CameraSampler/ ../../services/CameraSampler/service.proto --go_out=plugins=grpc:../../services/CameraSampler/
mv -f ../../services/CameraSampler/service.pb.go ./service.pb.go
