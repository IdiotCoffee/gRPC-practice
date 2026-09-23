from concurrent import futures
import grpc
import hello_pb2_grpc
import hello_pb2

class HelloSayer(hello_pb2_grpc.HelloSayerServicer):
    def SpeakPersonDetails(self, request, context):
        return hello_pb2.OutputStatement(
            introduction=f"Hello, {request.name}, are you {request.age} years old?",
            other_info = f"I've hear you already earn {request.salary}! Are you married? Let me check.... married={request.is_married}"
        )

server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

hello_pb2_grpc.add_HelloSayerServicer_to_server(
    HelloSayer(),
    server
)
server.add_insecure_port("[::]:50051")
server.start()
print("Server running on port 50051")
server.wait_for_termination()
