/**
 * Usage
 *   $ npx ts-node src/todos/show.ts <title> open/closed
 */

import { credentials } from "@grpc/grpc-js";

import {
  TodosClient,
  ITodosClient,
} from "../protos/goagen_goa_todo_example_todos_grpc_pb";
import {
  CreateRequest,
  CreateResponse,
} from "../protos/goagen_goa_todo_example_todos_pb";

const client: ITodosClient = new TodosClient(
  `localhost:8080`,
  credentials.createInsecure()
);

const request: CreateRequest = new CreateRequest();
request.setTitle(process.argv[2]);
request.setState(process.argv[3]);

client.create(request, (err: any, response: CreateResponse) => {
  if (err) {
    console.error(err);
    return;
  }
  console.log(response.toObject());
});
