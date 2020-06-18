import protobufjs from "protobufjs";
import { IdentityPayload } from "./generated/ivms101_pb";

export * as ivms from "./generated/ivms101_pb";
export const normalizedJsonOutput = (payload: IdentityPayload): Promise<any> =>
  new Promise((resolve, reject) => {
    const bin = payload.serializeBinary();

    protobufjs.load("../ivms101.proto", (err, root) => {
      if (err) reject(err);
      const payloadType = root!.lookupType("ivms101.IdentityPayload");
      const decodedMessage = payloadType.decode(bin);

      const toObjectWithString = payloadType.toObject(decodedMessage, {
        longs: String,
        enums: String,
        bytes: String,
      });

      resolve(toObjectWithString);
    });
  });
