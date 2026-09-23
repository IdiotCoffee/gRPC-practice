import asyncio
from grpc import aio

import hello_pb2
import hello_pb2_grpc

class HelloSayerServicer(hello_pb2_grpc.HelloSayerServicer):
    async def SpeakPersonDetails(self, request, context):
        return hello_pb2.OutputStatement(
            introduction=f"Hello, my name is {request.name}",
            other_info = f"I am {request.age} years old and earn {request.salary}. My married status is: {request.is_married}"
        )

async def server():
    server = aio.server()

    hello_pb2_grpc.add_HelloSayerServicer_to_server(
        HelloSayerServicer(),
        server
    )
    server.add_insecure_port("[::]:50051")
    await server.start()

    print("server started on port 50051")

    await server.wait_for_termination()



if __name__ == "__main__":
    asyncio.run(server())
