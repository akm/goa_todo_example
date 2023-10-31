/**
 * Usage
 *   $ npx ts-node src/todos/show.ts <id> <title> open/closed
 */

import { ChannelCredentials } from "@grpc/grpc-js";
import { GrpcTransport } from "@protobuf-ts/grpc-transport";

import {
  TodosClient,
  ITodosClient,
} from "../protos/goagen_goa_todo_example_todos.client";

const main = async () => {
  const transport = new GrpcTransport({
    host: "localhost:8080",
    channelCredentials: ChannelCredentials.createInsecure(),
  });

  const client: ITodosClient = new TodosClient(transport);

  const { response } = await client.update({
    id: BigInt(process.argv[2]),
    title: process.argv[3],
    state: process.argv[4],
  });

  console.log(response);
};

main().then(() => {
  console.log("done");
});
