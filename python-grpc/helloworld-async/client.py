import asyncio
from grpc import aio

import hello_pb2
import hello_pb2_grpc

async def make_request(stub, request):
    return await stub.SpeakPersonDetails(request)
async def main():
    async with aio.insecure_channel("localhost:50051") as channel:

        stub = hello_pb2_grpc.HelloSayerStub(channel)

        request1 = hello_pb2.InputInfo(
            name = "John Doe",
            age = 50,
            salary = 10000000,
           is_married = False
        )

        request2 = hello_pb2.InputInfo(
            name = "Judy Alvarez",
            age = 23,
            salary = 20000,
           is_married = False
        )

        request3 = hello_pb2.InputInfo(
            name = "V",
            age = 25,
            salary = 800000,
           is_married = True
        )

        request4 = hello_pb2.InputInfo(
            name = "Johnny",
            age = 103,
            salary = 1000000,
           is_married = False
        )

        request5 = hello_pb2.InputInfo(
            name = "Thor",
            age = 1000,
            salary = 80,
           is_married = True
        )

        tasks = [
                asyncio.create_task(make_request(stub, request1)),
                asyncio.create_task(make_request(stub, request2)),
                asyncio.create_task(make_request(stub, request3)),
                asyncio.create_task(make_request(stub, request4)),
                asyncio.create_task(make_request(stub, request5)),
                ]

        for task in asyncio.as_completed(tasks):
            response = await task

            print(response.introduction)
            print(response.other_info)
            print("---")




if __name__ == "__main__":
    asyncio.run(main())
