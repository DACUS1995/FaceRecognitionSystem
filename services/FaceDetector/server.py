from concurrent import futures
import logging
import sys
import os
import dlib
import glob
import grpc
import io
import numpy as np
from PIL import Image

import service_pb2
import service_pb2_grpc
from detector import Detector

class Server():
	def __init__(self, port = 50051):
		self._port = port
		self._server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

		service_pb2_grpc.add_FaceDetectorServicer_to_server(Detector(), self._server)
		self._server.add_insecure_port(f"[::]:{self._port}")

	def run(self):
		self._server.start()
		print(f"FaceDetector service::Started server on port {self._port}")
		self._server.wait_for_termination()
		print("FaceDetector service::Closed server")

if __name__ == '__main__':
	logging.basicConfig()
	server = Server()
	server.run()
	# detect_face()