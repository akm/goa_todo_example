/**
 * Usage
 *   $ npx ts-node src/todos/show.ts <title> open/closed
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

  const { response } = await client.create({
    title: process.argv[2],
    state: process.argv[3],
  });

  // const { id, title, state, createdAt, updatedAt } = response;
  // console.log({ id, title, state, createdAt, updatedAt });
  console.log(response);
};

main().then(() => {
  console.log("done");
});
