import generated.ivms101_pb2 as ivms
from google.protobuf.json_format import MessageToDict
import json

def normalized_json_output(payload):
  patload_dict = MessageToDict(payload)
  return json.dumps(patload_dict, ensure_ascii=False, indent=2)