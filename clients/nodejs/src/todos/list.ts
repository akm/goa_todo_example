/**
 * Usage
 *   $ npx ts-node src/todos/list.ts
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

  const { response } = await client.list({});

  const { total, offset } = response;
  console.log({ total, offset });
  response.items?.field.forEach((item) => {
    console.log(item);
  });
};

main().then(() => {
  console.log("done");
});
