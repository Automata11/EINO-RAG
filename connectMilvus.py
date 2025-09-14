from pymilvus import MilvusClient

client = MilvusClient(host='localhost', port='19530')
print("Connected to Milvus server:", client)