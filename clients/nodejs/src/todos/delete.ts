/**
 * Usage
 *   $ npx ts-node src/todos/show.ts <id>
 */

import { credentials } from "@grpc/grpc-js";

import {
  TodosClient,
  ITodosClient,
} from "../protos/goagen_goa_todo_example_todos_grpc_pb";
import {
  DeleteRequest,
  DeleteResponse,
} from "../protos/goagen_goa_todo_example_todos_pb";

const client: ITodosClient = new TodosClient(
  `localhost:8080`,
  credentials.createInsecure()
);

const request: DeleteRequest = new DeleteRequest();
request.setId(parseInt(process.argv[2], 10));

client.delete(request, (err: any, response: DeleteResponse) => {
  console.log(response.toObject());
});
