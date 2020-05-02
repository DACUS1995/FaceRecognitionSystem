protoc -I ../../services/FaceDetector/ ../../services/FaceDetector/service_face_detector.proto --go_out=plugins=grpc:../../services/FaceDetector/
mv -f ../../services/FaceDetector/service_face_detector.pb.go ./service_face_detector.pb.go
