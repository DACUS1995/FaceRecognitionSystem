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

import service_face_detector_pb2
import service_face_detector_pb2_grpc


class Detector(service_face_detector_pb2_grpc.FaceDetector):
	def DetectFaces(self, request, context):

		# for data in request_iterator:
		# 	image.append(data)

		image = request.image
		image = np.array(Image.open(io.BytesIO(image)))

		image_shape = [349, 620, 3]
		image_shape = request.image_shape
		image = np.reshape(image, image_shape)

		detected_faces_boxes, detected_faces_embeddings = self._detect_face(image, False)
		return service_face_detector_pb2.Reply(
			detected_faces_boxes = detected_faces_boxes,
			detected_faces_embeddings = detected_faces_embeddings
		)

	def _detect_face(self, image, display_image = True):
		detector = dlib.get_frontal_face_detector()
		sp = dlib.shape_predictor("./trained_models/shape_predictor.dat")
		facerec = dlib.face_recognition_model_v1("./trained_models/face_recognition.dat")

		image_window = None
		if display_image == True:
			image_window = dlib.image_window()
			image_window.clear_overlay()
			image_window.set_image(image)

		# Ask the detector to find the bounding boxes of each face. The 1 in the
		# second argument indicates that we should upsample the image 1 time. This
		# will make everything bigger and allow us to detect more faces.
		dets = detector(image, 1)
		print("Number of faces detected: {}".format(len(dets)))

		detected_faces_boxes = []
		detected_faces_embeddings = []

		# Now process each face we found.
		for k, d in enumerate(dets):
			detected_faces_boxes.append(d.left())
			detected_faces_boxes.append(d.top())
			detected_faces_boxes.append(d.right())
			detected_faces_boxes.append(d.bottom())

			# Get the landmarks/parts for the face in box d.
			shape = sp(image, d)

			if display_image == True:
				# Draw the face landmarks on the screen so we can see what face is currently being processed.
				image_window.clear_overlay()
				image_window.add_overlay(d)
				image_window.add_overlay(shape)

			face_descriptor = facerec.compute_face_descriptor(image, shape) #128 size vector
			face_chip = dlib.get_face_chip(image, shape)
			face_descriptor_from_prealigned_image = facerec.compute_face_descriptor(face_chip)
			face_descriptor_from_prealigned_image = np.asarray(face_descriptor_from_prealigned_image)

			detected_faces_embeddings += face_descriptor_from_prealigned_image.tolist()
			# dlib.hit_enter_to_continue()

		return detected_faces_boxes, detected_faces_embeddings
