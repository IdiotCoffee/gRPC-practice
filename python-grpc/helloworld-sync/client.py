import grpc

import hello_pb2
import hello_pb2_grpc

channel = grpc.insecure_channel("localhost:50051")
stub = hello_pb2_grpc.HelloSayerStub(channel)
request = hello_pb2.InputInfo(name="John", age=23, salary=5000000,is_married=False)
response = stub.SpeakPersonDetails(request)
print(response.introduction, response.other_info)
