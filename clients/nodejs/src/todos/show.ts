/**
 * Usage
 *   $ npx ts-node src/todos/show.ts <todo_id>
 */

import { credentials } from "@grpc/grpc-js";

import {
  TodosClient,
  ITodosClient,
} from "../protos/goagen_goa_todo_example_todos_grpc_pb";
import {
  ShowRequest,
  ShowResponse,
} from "../protos/goagen_goa_todo_example_todos_pb";

const client: ITodosClient = new TodosClient(
  `localhost:8080`,
  credentials.createInsecure()
);

const request: ShowRequest = new ShowRequest();
request.setId(parseInt(process.argv[2], 10));

client.show(request, (err: any, response: ShowResponse) => {
  console.log(response.toObject());
});
