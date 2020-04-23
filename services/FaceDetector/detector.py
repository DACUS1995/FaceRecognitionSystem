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


class Detector(service_pb2_grpc.FaceDetector):
	def DetectFaces(self, request, context):
		# image = []

		# for data in request_iterator:
		# 	image.append(data)

		image_shape = [349, 620, 3]

		image = request.image
		# image = np.ndarray.tobytes(image)
		# image = np.frombuffer(image, dtype=np.uint8)

		image = np.array(Image.open(io.BytesIO(image)))
		image = np.reshape(image, image_shape)

		detected_faces = self._detect_face(image)
		return service_pb2.Reply(detected_faces=detected_faces)

	def _detect_face(self, image):
		detector = dlib.get_frontal_face_detector()
		sp = dlib.shape_predictor("./trained_models/shape_predictor.dat")
		facerec = dlib.face_recognition_model_v1("./trained_models/face_recognition.dat")
		win = dlib.image_window()
			
		win.clear_overlay()
		win.set_image(image)

		# Ask the detector to find the bounding boxes of each face. The 1 in the
		# second argument indicates that we should upsample the image 1 time. This
		# will make everything bigger and allow us to detect more faces.
		dets = detector(image, 1)
		print("Number of faces detected: {}".format(len(dets)))

		detected_faces_boxes = []

		# Now process each face we found.
		for k, d in enumerate(dets):
			detected_faces_boxes.append(d.left())
			detected_faces_boxes.append(d.top())
			detected_faces_boxes.append(d.right())
			detected_faces_boxes.append(d.bottom())
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

			# dlib.hit_enter_to_continue()

		return detected_faces_boxes
