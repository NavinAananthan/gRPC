import grpc
import service_pb2
import service_pb2_grpc


def run():
    with grpc.insecure_channel('localhost:50052') as channel:
        stub = service_pb2_grpc.CommunicationServiceStub(channel)
        response = stub.SendMessage(service_pb2.MessageRequest(message='Hello from Python!'))
        print("Response from Go server:", response.response)

if __name__ == '__main__':
    run()
