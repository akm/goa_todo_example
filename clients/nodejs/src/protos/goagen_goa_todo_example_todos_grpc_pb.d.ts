// package: todos
// file: goagen_goa_todo_example_todos.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as goagen_goa_todo_example_todos_pb from "./goagen_goa_todo_example_todos_pb";

interface ITodosService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    list: ITodosService_IList;
    show: ITodosService_IShow;
    create: ITodosService_ICreate;
    update: ITodosService_IUpdate;
    delete: ITodosService_IDelete;
}

interface ITodosService_IList extends grpc.MethodDefinition<goagen_goa_todo_example_todos_pb.ListRequest, goagen_goa_todo_example_todos_pb.ListResponse> {
    path: "/todos.Todos/List";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<goagen_goa_todo_example_todos_pb.ListRequest>;
    requestDeserialize: grpc.deserialize<goagen_goa_todo_example_todos_pb.ListRequest>;
    responseSerialize: grpc.serialize<goagen_goa_todo_example_todos_pb.ListResponse>;
    responseDeserialize: grpc.deserialize<goagen_goa_todo_example_todos_pb.ListResponse>;
}
interface ITodosService_IShow extends grpc.MethodDefinition<goagen_goa_todo_example_todos_pb.ShowRequest, goagen_goa_todo_example_todos_pb.ShowResponse> {
    path: "/todos.Todos/Show";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<goagen_goa_todo_example_todos_pb.ShowRequest>;
    requestDeserialize: grpc.deserialize<goagen_goa_todo_example_todos_pb.ShowRequest>;
    responseSerialize: grpc.serialize<goagen_goa_todo_example_todos_pb.ShowResponse>;
    responseDeserialize: grpc.deserialize<goagen_goa_todo_example_todos_pb.ShowResponse>;
}
interface ITodosService_ICreate extends grpc.MethodDefinition<goagen_goa_todo_example_todos_pb.CreateRequest, goagen_goa_todo_example_todos_pb.CreateResponse> {
    path: "/todos.Todos/Create";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<goagen_goa_todo_example_todos_pb.CreateRequest>;
    requestDeserialize: grpc.deserialize<goagen_goa_todo_example_todos_pb.CreateRequest>;
    responseSerialize: grpc.serialize<goagen_goa_todo_example_todos_pb.CreateResponse>;
    responseDeserialize: grpc.deserialize<goagen_goa_todo_example_todos_pb.CreateResponse>;
}
interface ITodosService_IUpdate extends grpc.MethodDefinition<goagen_goa_todo_example_todos_pb.UpdateRequest, goagen_goa_todo_example_todos_pb.UpdateResponse> {
    path: "/todos.Todos/Update";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<goagen_goa_todo_example_todos_pb.UpdateRequest>;
    requestDeserialize: grpc.deserialize<goagen_goa_todo_example_todos_pb.UpdateRequest>;
    responseSerialize: grpc.serialize<goagen_goa_todo_example_todos_pb.UpdateResponse>;
    responseDeserialize: grpc.deserialize<goagen_goa_todo_example_todos_pb.UpdateResponse>;
}
interface ITodosService_IDelete extends grpc.MethodDefinition<goagen_goa_todo_example_todos_pb.DeleteRequest, goagen_goa_todo_example_todos_pb.DeleteResponse> {
    path: "/todos.Todos/Delete";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<goagen_goa_todo_example_todos_pb.DeleteRequest>;
    requestDeserialize: grpc.deserialize<goagen_goa_todo_example_todos_pb.DeleteRequest>;
    responseSerialize: grpc.serialize<goagen_goa_todo_example_todos_pb.DeleteResponse>;
    responseDeserialize: grpc.deserialize<goagen_goa_todo_example_todos_pb.DeleteResponse>;
}

export const TodosService: ITodosService;

export interface ITodosServer extends grpc.UntypedServiceImplementation {
    list: grpc.handleUnaryCall<goagen_goa_todo_example_todos_pb.ListRequest, goagen_goa_todo_example_todos_pb.ListResponse>;
    show: grpc.handleUnaryCall<goagen_goa_todo_example_todos_pb.ShowRequest, goagen_goa_todo_example_todos_pb.ShowResponse>;
    create: grpc.handleUnaryCall<goagen_goa_todo_example_todos_pb.CreateRequest, goagen_goa_todo_example_todos_pb.CreateResponse>;
    update: grpc.handleUnaryCall<goagen_goa_todo_example_todos_pb.UpdateRequest, goagen_goa_todo_example_todos_pb.UpdateResponse>;
    delete: grpc.handleUnaryCall<goagen_goa_todo_example_todos_pb.DeleteRequest, goagen_goa_todo_example_todos_pb.DeleteResponse>;
}

export interface ITodosClient {
    list(request: goagen_goa_todo_example_todos_pb.ListRequest, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.ListResponse) => void): grpc.ClientUnaryCall;
    list(request: goagen_goa_todo_example_todos_pb.ListRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.ListResponse) => void): grpc.ClientUnaryCall;
    list(request: goagen_goa_todo_example_todos_pb.ListRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.ListResponse) => void): grpc.ClientUnaryCall;
    show(request: goagen_goa_todo_example_todos_pb.ShowRequest, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.ShowResponse) => void): grpc.ClientUnaryCall;
    show(request: goagen_goa_todo_example_todos_pb.ShowRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.ShowResponse) => void): grpc.ClientUnaryCall;
    show(request: goagen_goa_todo_example_todos_pb.ShowRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.ShowResponse) => void): grpc.ClientUnaryCall;
    create(request: goagen_goa_todo_example_todos_pb.CreateRequest, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    create(request: goagen_goa_todo_example_todos_pb.CreateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    create(request: goagen_goa_todo_example_todos_pb.CreateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    update(request: goagen_goa_todo_example_todos_pb.UpdateRequest, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    update(request: goagen_goa_todo_example_todos_pb.UpdateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    update(request: goagen_goa_todo_example_todos_pb.UpdateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    delete(request: goagen_goa_todo_example_todos_pb.DeleteRequest, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    delete(request: goagen_goa_todo_example_todos_pb.DeleteRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    delete(request: goagen_goa_todo_example_todos_pb.DeleteRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
}

export class TodosClient extends grpc.Client implements ITodosClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public list(request: goagen_goa_todo_example_todos_pb.ListRequest, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.ListResponse) => void): grpc.ClientUnaryCall;
    public list(request: goagen_goa_todo_example_todos_pb.ListRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.ListResponse) => void): grpc.ClientUnaryCall;
    public list(request: goagen_goa_todo_example_todos_pb.ListRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.ListResponse) => void): grpc.ClientUnaryCall;
    public show(request: goagen_goa_todo_example_todos_pb.ShowRequest, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.ShowResponse) => void): grpc.ClientUnaryCall;
    public show(request: goagen_goa_todo_example_todos_pb.ShowRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.ShowResponse) => void): grpc.ClientUnaryCall;
    public show(request: goagen_goa_todo_example_todos_pb.ShowRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.ShowResponse) => void): grpc.ClientUnaryCall;
    public create(request: goagen_goa_todo_example_todos_pb.CreateRequest, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    public create(request: goagen_goa_todo_example_todos_pb.CreateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    public create(request: goagen_goa_todo_example_todos_pb.CreateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.CreateResponse) => void): grpc.ClientUnaryCall;
    public update(request: goagen_goa_todo_example_todos_pb.UpdateRequest, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    public update(request: goagen_goa_todo_example_todos_pb.UpdateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    public update(request: goagen_goa_todo_example_todos_pb.UpdateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.UpdateResponse) => void): grpc.ClientUnaryCall;
    public delete(request: goagen_goa_todo_example_todos_pb.DeleteRequest, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    public delete(request: goagen_goa_todo_example_todos_pb.DeleteRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
    public delete(request: goagen_goa_todo_example_todos_pb.DeleteRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: goagen_goa_todo_example_todos_pb.DeleteResponse) => void): grpc.ClientUnaryCall;
}
