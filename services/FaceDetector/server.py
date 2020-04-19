from concurrent import futures
import logging
import sys
import os
import dlib
import glob
import grpc

import service_pb2
import service_pb2_grpc


class Greeter(service_pb2_grpc.FaceDetector):

	def SayHello(self, request, context):
		return service_pb2.HelloReply(message='Hello, %s!' % request.name)



def detect_face():
	faces_folder_path = "./test_images"
	detector = dlib.get_frontal_face_detector()
	sp = dlib.shape_predictor("./trained_models/shape_predictor.dat")
	facerec = dlib.face_recognition_model_v1("./trained_models/face_recognition.dat")

	win = dlib.image_window()

	for f in glob.glob(os.path.join(faces_folder_path, "*.jpg")):
		print("Processing file: {}".format(f))
		img = dlib.load_rgb_image(f)

		win.clear_overlay()
		win.set_image(img)

		# Ask the detector to find the bounding boxes of each face. The 1 in the
		# second argument indicates that we should upsample the image 1 time. This
		# will make everything bigger and allow us to detect more faces.
		dets = detector(img, 1)
		print("Number of faces detected: {}".format(len(dets)))


		# Now process each face we found.
		for k, d in enumerate(dets):
			print("Detection {}: Left: {} Top: {} Right: {} Bottom: {}".format(k, d.left(), d.top(), d.right(), d.bottom()))
			
			# Get the landmarks/parts for the face in box d.
			shape = sp(img, d)
			
			# Draw the face landmarks on the screen so we can see what face is currently being processed.
			win.clear_overlay()
			win.add_overlay(d)
			win.add_overlay(shape)

			face_descriptor = facerec.compute_face_descriptor(img, shape) #128 size vector
			
			print("Computing descriptor on aligned image ..")
			face_chip = dlib.get_face_chip(img, shape)
			face_descriptor_from_prealigned_image = facerec.compute_face_descriptor(face_chip)                
			
			dlib.hit_enter_to_continue()

def serve():
	server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
	service_pb2_grpc.add_FaceDetectorServicer_to_server(Greeter(), server)
	server.add_insecure_port('[::]:50051')
	server.start()
	server.wait_for_termination()

class Server():
	def __init__(self):
		pass

	def run(self):
		pass


if __name__ == '__main__':
	# logging.basicConfig()
	# serve()
	detect_face()