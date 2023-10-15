/**
 * Usage
 *   $ npx ts-node src/todos/list.ts
 */

import { credentials } from "@grpc/grpc-js";

import {
  TodosClient,
  ITodosClient,
} from "../protos/goagen_goa_todo_example_todos_grpc_pb";
import {
  ListRequest,
  ListResponse,
} from "../protos/goagen_goa_todo_example_todos_pb";

const client: ITodosClient = new TodosClient(
  `localhost:8080`,
  credentials.createInsecure()
);

const request: ListRequest = new ListRequest();

client.list(request, (err: any, response: ListResponse) => {
  response
    ?.getItems()
    ?.getFieldList()
    .forEach((todo) => {
      console.log(todo.toObject());
    });
});
