# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import service_pb2 as service__pb2


class FaceDetectorStub(object):
    """Missing associated documentation comment in .proto file"""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.DetectFaces = channel.unary_unary(
                '/face_detector.FaceDetector/DetectFaces',
                request_serializer=service__pb2.Request.SerializeToString,
                response_deserializer=service__pb2.Reply.FromString,
                )


class FaceDetectorServicer(object):
    """Missing associated documentation comment in .proto file"""

    def DetectFaces(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_FaceDetectorServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'DetectFaces': grpc.unary_unary_rpc_method_handler(
                    servicer.DetectFaces,
                    request_deserializer=service__pb2.Request.FromString,
                    response_serializer=service__pb2.Reply.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'face_detector.FaceDetector', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class FaceDetector(object):
    """Missing associated documentation comment in .proto file"""

    @staticmethod
    def DetectFaces(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/face_detector.FaceDetector/DetectFaces',
            service__pb2.Request.SerializeToString,
            service__pb2.Reply.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)