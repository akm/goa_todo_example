// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// Code generated with goa v3.13.2, DO NOT EDIT.
//
// todos protocol buffer definition
//
// Command:
// $ goa gen github.com/akm/goa_todo_example/design
//
'use strict';
var grpc = require('@grpc/grpc-js');
var goagen_goa_todo_example_todos_pb = require('./goagen_goa_todo_example_todos_pb.js');

function serialize_todos_CreateRequest(arg) {
  if (!(arg instanceof goagen_goa_todo_example_todos_pb.CreateRequest)) {
    throw new Error('Expected argument of type todos.CreateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_todos_CreateRequest(buffer_arg) {
  return goagen_goa_todo_example_todos_pb.CreateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todos_CreateResponse(arg) {
  if (!(arg instanceof goagen_goa_todo_example_todos_pb.CreateResponse)) {
    throw new Error('Expected argument of type todos.CreateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_todos_CreateResponse(buffer_arg) {
  return goagen_goa_todo_example_todos_pb.CreateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todos_DeleteRequest(arg) {
  if (!(arg instanceof goagen_goa_todo_example_todos_pb.DeleteRequest)) {
    throw new Error('Expected argument of type todos.DeleteRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_todos_DeleteRequest(buffer_arg) {
  return goagen_goa_todo_example_todos_pb.DeleteRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todos_DeleteResponse(arg) {
  if (!(arg instanceof goagen_goa_todo_example_todos_pb.DeleteResponse)) {
    throw new Error('Expected argument of type todos.DeleteResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_todos_DeleteResponse(buffer_arg) {
  return goagen_goa_todo_example_todos_pb.DeleteResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todos_ListRequest(arg) {
  if (!(arg instanceof goagen_goa_todo_example_todos_pb.ListRequest)) {
    throw new Error('Expected argument of type todos.ListRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_todos_ListRequest(buffer_arg) {
  return goagen_goa_todo_example_todos_pb.ListRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todos_ListResponse(arg) {
  if (!(arg instanceof goagen_goa_todo_example_todos_pb.ListResponse)) {
    throw new Error('Expected argument of type todos.ListResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_todos_ListResponse(buffer_arg) {
  return goagen_goa_todo_example_todos_pb.ListResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todos_ShowRequest(arg) {
  if (!(arg instanceof goagen_goa_todo_example_todos_pb.ShowRequest)) {
    throw new Error('Expected argument of type todos.ShowRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_todos_ShowRequest(buffer_arg) {
  return goagen_goa_todo_example_todos_pb.ShowRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todos_ShowResponse(arg) {
  if (!(arg instanceof goagen_goa_todo_example_todos_pb.ShowResponse)) {
    throw new Error('Expected argument of type todos.ShowResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_todos_ShowResponse(buffer_arg) {
  return goagen_goa_todo_example_todos_pb.ShowResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todos_UpdateRequest(arg) {
  if (!(arg instanceof goagen_goa_todo_example_todos_pb.UpdateRequest)) {
    throw new Error('Expected argument of type todos.UpdateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_todos_UpdateRequest(buffer_arg) {
  return goagen_goa_todo_example_todos_pb.UpdateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todos_UpdateResponse(arg) {
  if (!(arg instanceof goagen_goa_todo_example_todos_pb.UpdateResponse)) {
    throw new Error('Expected argument of type todos.UpdateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_todos_UpdateResponse(buffer_arg) {
  return goagen_goa_todo_example_todos_pb.UpdateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Service is the todos service interface.
var TodosService = exports.TodosService = {
  // List implements list.
list: {
    path: '/todos.Todos/List',
    requestStream: false,
    responseStream: false,
    requestType: goagen_goa_todo_example_todos_pb.ListRequest,
    responseType: goagen_goa_todo_example_todos_pb.ListResponse,
    requestSerialize: serialize_todos_ListRequest,
    requestDeserialize: deserialize_todos_ListRequest,
    responseSerialize: serialize_todos_ListResponse,
    responseDeserialize: deserialize_todos_ListResponse,
  },
  // Show implements show.
show: {
    path: '/todos.Todos/Show',
    requestStream: false,
    responseStream: false,
    requestType: goagen_goa_todo_example_todos_pb.ShowRequest,
    responseType: goagen_goa_todo_example_todos_pb.ShowResponse,
    requestSerialize: serialize_todos_ShowRequest,
    requestDeserialize: deserialize_todos_ShowRequest,
    responseSerialize: serialize_todos_ShowResponse,
    responseDeserialize: deserialize_todos_ShowResponse,
  },
  // Create implements create.
create: {
    path: '/todos.Todos/Create',
    requestStream: false,
    responseStream: false,
    requestType: goagen_goa_todo_example_todos_pb.CreateRequest,
    responseType: goagen_goa_todo_example_todos_pb.CreateResponse,
    requestSerialize: serialize_todos_CreateRequest,
    requestDeserialize: deserialize_todos_CreateRequest,
    responseSerialize: serialize_todos_CreateResponse,
    responseDeserialize: deserialize_todos_CreateResponse,
  },
  // Update implements update.
update: {
    path: '/todos.Todos/Update',
    requestStream: false,
    responseStream: false,
    requestType: goagen_goa_todo_example_todos_pb.UpdateRequest,
    responseType: goagen_goa_todo_example_todos_pb.UpdateResponse,
    requestSerialize: serialize_todos_UpdateRequest,
    requestDeserialize: deserialize_todos_UpdateRequest,
    responseSerialize: serialize_todos_UpdateResponse,
    responseDeserialize: deserialize_todos_UpdateResponse,
  },
  // Delete implements delete.
delete: {
    path: '/todos.Todos/Delete',
    requestStream: false,
    responseStream: false,
    requestType: goagen_goa_todo_example_todos_pb.DeleteRequest,
    responseType: goagen_goa_todo_example_todos_pb.DeleteResponse,
    requestSerialize: serialize_todos_DeleteRequest,
    requestDeserialize: deserialize_todos_DeleteRequest,
    responseSerialize: serialize_todos_DeleteResponse,
    responseDeserialize: deserialize_todos_DeleteResponse,
  },
};

exports.TodosClient = grpc.makeGenericClientConstructor(TodosService);
