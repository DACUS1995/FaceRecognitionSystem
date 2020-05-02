protoc -I ../../services/FaceDetector/ ../../services/FaceDetector/service.proto --go_out=plugins=grpc:../../services/FaceDetector/
mv -f ../../services/FaceDetector/service.pb.go ./service_face_detector.pb.go
