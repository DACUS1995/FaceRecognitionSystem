# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
import service_pb2 as service__pb2


class CameraSamplerStub(object):
    """Missing associated documentation comment in .proto file"""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SampleImage = channel.unary_unary(
                '/camera_sampler.CameraSampler/SampleImage',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=service__pb2.Reply.FromString,
                )


class CameraSamplerServicer(object):
    """Missing associated documentation comment in .proto file"""

    def SampleImage(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CameraSamplerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'SampleImage': grpc.unary_unary_rpc_method_handler(
                    servicer.SampleImage,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=service__pb2.Reply.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'camera_sampler.CameraSampler', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class CameraSampler(object):
    """Missing associated documentation comment in .proto file"""

    @staticmethod
    def SampleImage(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/camera_sampler.CameraSampler/SampleImage',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            service__pb2.Reply.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
