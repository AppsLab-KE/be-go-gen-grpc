# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from auth import auth_pb2 as auth_dot_auth__pb2


class AuthStub(object):
    """Service
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Login = channel.unary_unary(
                '/Auth/Login',
                request_serializer=auth_dot_auth__pb2.LoginRequest.SerializeToString,
                response_deserializer=auth_dot_auth__pb2.LoginResponse.FromString,
                )
        self.Register = channel.unary_unary(
                '/Auth/Register',
                request_serializer=auth_dot_auth__pb2.RegisterRequest.SerializeToString,
                response_deserializer=auth_dot_auth__pb2.RegisterResponse.FromString,
                )
        self.Logout = channel.unary_unary(
                '/Auth/Logout',
                request_serializer=auth_dot_auth__pb2.LogoutRequest.SerializeToString,
                response_deserializer=auth_dot_auth__pb2.LogoutResponse.FromString,
                )


class AuthServicer(object):
    """Service
    """

    def Login(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Register(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Logout(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AuthServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Login': grpc.unary_unary_rpc_method_handler(
                    servicer.Login,
                    request_deserializer=auth_dot_auth__pb2.LoginRequest.FromString,
                    response_serializer=auth_dot_auth__pb2.LoginResponse.SerializeToString,
            ),
            'Register': grpc.unary_unary_rpc_method_handler(
                    servicer.Register,
                    request_deserializer=auth_dot_auth__pb2.RegisterRequest.FromString,
                    response_serializer=auth_dot_auth__pb2.RegisterResponse.SerializeToString,
            ),
            'Logout': grpc.unary_unary_rpc_method_handler(
                    servicer.Logout,
                    request_deserializer=auth_dot_auth__pb2.LogoutRequest.FromString,
                    response_serializer=auth_dot_auth__pb2.LogoutResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Auth', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Auth(object):
    """Service
    """

    @staticmethod
    def Login(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Auth/Login',
            auth_dot_auth__pb2.LoginRequest.SerializeToString,
            auth_dot_auth__pb2.LoginResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Register(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Auth/Register',
            auth_dot_auth__pb2.RegisterRequest.SerializeToString,
            auth_dot_auth__pb2.RegisterResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Logout(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Auth/Logout',
            auth_dot_auth__pb2.LogoutRequest.SerializeToString,
            auth_dot_auth__pb2.LogoutResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
