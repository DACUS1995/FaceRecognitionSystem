protoc -I ../../services/CameraSampler/ ../../services/CameraSampler/service_camera_sampler.proto --go_out=plugins=grpc:../../services/CameraSampler/
mv -f ../../services/CameraSampler/service_camera_sampler.pb.go ./service_camera_sampler.pb.go
