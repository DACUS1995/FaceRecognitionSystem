protoc -I ../../services/FaceDetector/ ../../services/FaceDetector/service.proto --go_out=plugins=grpc:../../services/FaceDetector/
REM go generate ../../services/FaceDetector/ ...
REM go generate google.golang.org/grpc/examples/helloworld/...
