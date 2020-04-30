from concurrent import futures
import logging
import sys
import os
import dlib
import glob
import grpc
import io
import numpy as np
import cv2

import service_pb2
import service_pb2_grpc


class Detector(service_pb2_grpc.CameraSampler):
	def __init__(self):
		super().__init__()
		self.camera = cv2.VideoCapture(0)
		self.width = cap.get(3)
		self.height = cap.get(4)

	def SampleImage(self, request, context):
		ret, frame = cam.read()


		return service_pb2.Reply(
			image = frame.tobytes(),
			image_shape = [self.width, self.height]
		)


	def __del__(self):
		self.camera.release()