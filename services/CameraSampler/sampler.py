from concurrent import futures
import logging
import sys
import os
import glob
import grpc
import io
import numpy as np
import cv2

import service_camera_sampler_pb2
import service_camera_sampler_pb2_grpc


class Sampler(service_camera_sampler_pb2_grpc.CameraSampler):
	def __init__(self):
		super().__init__()
		self.camera = cv2.VideoCapture(0)
		self.width = self.camera.get(3)
		self.height = self.camera.get(4)

	def SampleImage(self, request, context):
		ret, frame = self.camera.read()


		return service_camera_sampler_pb2.Reply(
			image = frame.tobytes(),
			image_shape = [int(self.width), int(self.height)]
		)


	def __del__(self):
		self.camera.release()