/**
 * Usage
 *   $ npx ts-node src/todos/show.ts <id> <title> open/closed
 */

import { credentials } from "@grpc/grpc-js";

import {
  TodosClient,
  ITodosClient,
} from "../protos/goagen_goa_todo_example_todos_grpc_pb";
import {
  TodoUpdatePayload,
  UpdateRequest,
  UpdateResponse,
} from "../protos/goagen_goa_todo_example_todos_pb";

const client: ITodosClient = new TodosClient(
  `localhost:8080`,
  credentials.createInsecure()
);

const request: UpdateRequest = new UpdateRequest();
request.setId(parseInt(process.argv[2], 10));
const body = new TodoUpdatePayload();
request.setBody(body);
body.setTitle(process.argv[3]);
body.setState(process.argv[4]);

client.update(request, (err: any, response: UpdateResponse) => {
  console.log(response.toObject());
});
