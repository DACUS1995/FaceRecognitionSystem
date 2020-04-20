from concurrent import futures
import logging
import sys
import os
import dlib
import glob
import grpc
import numpy as np

import service_pb2
import service_pb2_grpc


class Greeter(service_pb2_grpc.FaceDetector):

	# def SayHello(self, request, context):
	# 	return service_pb2.HelloReply(message='Hello, %s!' % request.name)

class Detector(service_pb2_grpc.FaceDetector):
	def DetectFaces(self, request, context):
		# image = []

		# for data in request_iterator:
		# 	image.append(data)

		image_shape = [349, 620, 3]

		image = request.image
		image = np.ndarray.tobytes(image)
		image = np.frombuffer(image, dtype=np.uint8)
		image = np.reshape(images, image_shape)

		detected_faces = self._detect_face(image)
		return service_pb2.Reply(detected_faces=detected_faces)

	def _detect_face(image):
		detector = dlib.get_frontal_face_detector()
		sp = dlib.shape_predictor("./trained_models/shape_predictor.dat")
		facerec = dlib.face_recognition_model_v1("./trained_models/face_recognition.dat")
			
		win.clear_overlay()
		win.set_image(image)

		# Ask the detector to find the bounding boxes of each face. The 1 in the
		# second argument indicates that we should upsample the image 1 time. This
		# will make everything bigger and allow us to detect more faces.
		dets = detector(image, 1)
		print("Number of faces detected: {}".format(len(dets)))


		# Now process each face we found.
		for k, d in enumerate(dets):
			print("Detection {}: Left: {} Top: {} Right: {} Bottom: {}".format(k, d.left(), d.top(), d.right(), d.bottom()))

			# Get the landmarks/parts for the face in box d.
			shape = sp(image, d)

			# Draw the face landmarks on the screen so we can see what face is currently being processed.
			win.clear_overlay()
			win.add_overlay(d)
			win.add_overlay(shape)

			face_descriptor = facerec.compute_face_descriptor(image, shape) #128 size vector

			print("Computing descriptor on aligned image ..")
			face_chip = dlib.get_face_chip(image, shape)
			face_descriptor_from_prealigned_image = facerec.compute_face_descriptor(face_chip)                

			dlib.hit_enter_to_continue()

		return dets


class Server():
	def __init__(self):
		self._server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
		service_pb2_grpc.add_FaceDetectorServicer_to_server(Detector, self._server)
		self._server.add_insecure_port('[::]:50051')

	def run(self):
		self.server.start()
		self.server.wait_for_termination()


if __name__ == '__main__':
	logging.basicConfig()
	server = Server()
	server.run()
	# detect_face()