// package: todos
// file: goagen_goa_todo_example_todos.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class ListRequest extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListRequest): ListRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListRequest;
    static deserializeBinaryFromReader(message: ListRequest, reader: jspb.BinaryReader): ListRequest;
}

export namespace ListRequest {
    export type AsObject = {
    }
}

export class ListResponse extends jspb.Message { 

    hasItems(): boolean;
    clearItems(): void;
    getItems(): TodoListItemCollection | undefined;
    setItems(value?: TodoListItemCollection): ListResponse;
    getTotal(): number;
    setTotal(value: number): ListResponse;
    getOffset(): number;
    setOffset(value: number): ListResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListResponse): ListResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListResponse;
    static deserializeBinaryFromReader(message: ListResponse, reader: jspb.BinaryReader): ListResponse;
}

export namespace ListResponse {
    export type AsObject = {
        items?: TodoListItemCollection.AsObject,
        total: number,
        offset: number,
    }
}

export class TodoListItemCollection extends jspb.Message { 
    clearFieldList(): void;
    getFieldList(): Array<TodoListItem>;
    setFieldList(value: Array<TodoListItem>): TodoListItemCollection;
    addField(value?: TodoListItem, index?: number): TodoListItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TodoListItemCollection.AsObject;
    static toObject(includeInstance: boolean, msg: TodoListItemCollection): TodoListItemCollection.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TodoListItemCollection, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TodoListItemCollection;
    static deserializeBinaryFromReader(message: TodoListItemCollection, reader: jspb.BinaryReader): TodoListItemCollection;
}

export namespace TodoListItemCollection {
    export type AsObject = {
        fieldList: Array<TodoListItem.AsObject>,
    }
}

export class TodoListItem extends jspb.Message { 
    getId(): number;
    setId(value: number): TodoListItem;
    getTitle(): string;
    setTitle(value: string): TodoListItem;
    getState(): string;
    setState(value: string): TodoListItem;
    getCreatedAt(): string;
    setCreatedAt(value: string): TodoListItem;
    getUpdatedAt(): string;
    setUpdatedAt(value: string): TodoListItem;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TodoListItem.AsObject;
    static toObject(includeInstance: boolean, msg: TodoListItem): TodoListItem.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TodoListItem, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TodoListItem;
    static deserializeBinaryFromReader(message: TodoListItem, reader: jspb.BinaryReader): TodoListItem;
}

export namespace TodoListItem {
    export type AsObject = {
        id: number,
        title: string,
        state: string,
        createdAt: string,
        updatedAt: string,
    }
}

export class ShowRequest extends jspb.Message { 
    getId(): number;
    setId(value: number): ShowRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ShowRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ShowRequest): ShowRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ShowRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ShowRequest;
    static deserializeBinaryFromReader(message: ShowRequest, reader: jspb.BinaryReader): ShowRequest;
}

export namespace ShowRequest {
    export type AsObject = {
        id: number,
    }
}

export class ShowResponse extends jspb.Message { 
    getId(): number;
    setId(value: number): ShowResponse;
    getTitle(): string;
    setTitle(value: string): ShowResponse;
    getState(): string;
    setState(value: string): ShowResponse;
    getCreatedAt(): string;
    setCreatedAt(value: string): ShowResponse;
    getUpdatedAt(): string;
    setUpdatedAt(value: string): ShowResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ShowResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ShowResponse): ShowResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ShowResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ShowResponse;
    static deserializeBinaryFromReader(message: ShowResponse, reader: jspb.BinaryReader): ShowResponse;
}

export namespace ShowResponse {
    export type AsObject = {
        id: number,
        title: string,
        state: string,
        createdAt: string,
        updatedAt: string,
    }
}

export class CreateRequest extends jspb.Message { 
    getTitle(): string;
    setTitle(value: string): CreateRequest;
    getState(): string;
    setState(value: string): CreateRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateRequest): CreateRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateRequest;
    static deserializeBinaryFromReader(message: CreateRequest, reader: jspb.BinaryReader): CreateRequest;
}

export namespace CreateRequest {
    export type AsObject = {
        title: string,
        state: string,
    }
}

export class CreateResponse extends jspb.Message { 
    getId(): number;
    setId(value: number): CreateResponse;
    getTitle(): string;
    setTitle(value: string): CreateResponse;
    getState(): string;
    setState(value: string): CreateResponse;
    getCreatedAt(): string;
    setCreatedAt(value: string): CreateResponse;
    getUpdatedAt(): string;
    setUpdatedAt(value: string): CreateResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CreateResponse): CreateResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateResponse;
    static deserializeBinaryFromReader(message: CreateResponse, reader: jspb.BinaryReader): CreateResponse;
}

export namespace CreateResponse {
    export type AsObject = {
        id: number,
        title: string,
        state: string,
        createdAt: string,
        updatedAt: string,
    }
}

export class UpdateRequest extends jspb.Message { 
    getId(): number;
    setId(value: number): UpdateRequest;

    hasBody(): boolean;
    clearBody(): void;
    getBody(): TodoUpdatePayload | undefined;
    setBody(value?: TodoUpdatePayload): UpdateRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateRequest): UpdateRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateRequest;
    static deserializeBinaryFromReader(message: UpdateRequest, reader: jspb.BinaryReader): UpdateRequest;
}

export namespace UpdateRequest {
    export type AsObject = {
        id: number,
        body?: TodoUpdatePayload.AsObject,
    }
}

export class TodoUpdatePayload extends jspb.Message { 
    getTitle(): string;
    setTitle(value: string): TodoUpdatePayload;
    getState(): string;
    setState(value: string): TodoUpdatePayload;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TodoUpdatePayload.AsObject;
    static toObject(includeInstance: boolean, msg: TodoUpdatePayload): TodoUpdatePayload.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TodoUpdatePayload, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TodoUpdatePayload;
    static deserializeBinaryFromReader(message: TodoUpdatePayload, reader: jspb.BinaryReader): TodoUpdatePayload;
}

export namespace TodoUpdatePayload {
    export type AsObject = {
        title: string,
        state: string,
    }
}

export class UpdateResponse extends jspb.Message { 
    getId(): number;
    setId(value: number): UpdateResponse;
    getTitle(): string;
    setTitle(value: string): UpdateResponse;
    getState(): string;
    setState(value: string): UpdateResponse;
    getCreatedAt(): string;
    setCreatedAt(value: string): UpdateResponse;
    getUpdatedAt(): string;
    setUpdatedAt(value: string): UpdateResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateResponse): UpdateResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateResponse;
    static deserializeBinaryFromReader(message: UpdateResponse, reader: jspb.BinaryReader): UpdateResponse;
}

export namespace UpdateResponse {
    export type AsObject = {
        id: number,
        title: string,
        state: string,
        createdAt: string,
        updatedAt: string,
    }
}

export class DeleteRequest extends jspb.Message { 
    getId(): number;
    setId(value: number): DeleteRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteRequest): DeleteRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteRequest;
    static deserializeBinaryFromReader(message: DeleteRequest, reader: jspb.BinaryReader): DeleteRequest;
}

export namespace DeleteRequest {
    export type AsObject = {
        id: number,
    }
}

export class DeleteResponse extends jspb.Message { 
    getId(): number;
    setId(value: number): DeleteResponse;
    getTitle(): string;
    setTitle(value: string): DeleteResponse;
    getState(): string;
    setState(value: string): DeleteResponse;
    getCreatedAt(): string;
    setCreatedAt(value: string): DeleteResponse;
    getUpdatedAt(): string;
    setUpdatedAt(value: string): DeleteResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteResponse.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteResponse): DeleteResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteResponse;
    static deserializeBinaryFromReader(message: DeleteResponse, reader: jspb.BinaryReader): DeleteResponse;
}

export namespace DeleteResponse {
    export type AsObject = {
        id: number,
        title: string,
        state: string,
        createdAt: string,
        updatedAt: string,
    }
}
