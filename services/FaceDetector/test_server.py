
import grpc

import service_pb2
import service_pb2_grpc


def run():
	with grpc.insecure_channel('localhost:50051') as channel:
		stub = service_pb2_grpc.FaceDetectorStub(channel)
		response = stub.SayHello(service_pb2.HelloRequest(name='you'))
	print("Greeter client received: " + response.message)


if __name__ == '__main__':
	run()