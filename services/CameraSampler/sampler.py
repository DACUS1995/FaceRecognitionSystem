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


class Detector(service_pb2_grpc.CameraSampler):
	def SampleImage(self, request, context):
		pass