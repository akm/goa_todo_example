package design

import (
	"time"

	. "goa.design/goa/v3/dsl"
)

func todoFields(dt string, action string) []string {
	r := []string{}

	if dt == "RT" || action == "update" {
		r = append(r, field(1, "id", UInt64, "ID"))
	}

	r = append(r, field(2, "title", String, "Title"))
	r = append(r, field(3, "State", String, "State", func() {
		Enum("open", "closed")
	}))

	if dt == "RT" {
		r = append(r, field(4, "created_at", String, "CreatedAt", func() { Format(FormatDateTime); Example(time.RFC3339) }))
		r = append(r, field(5, "updated_at", String, "UpdatedAt", func() { Format(FormatDateTime); Example(time.RFC3339) }))
	}

	return r
}

var TodoRT = ResultType("application/vnd.todo", func() {
	Attributes(func() {
		Required(todoFields("RT", "show")...)
	})
})
var TodoListItemRT = ResultType("application/vnd.todo-list-item", func() {
	Attributes(func() {
		Required(todoFields("RT", "show")...)
	})
})
var TodoListRT = ResultType("application/vnd.todo-list", func() {
	Attributes(func() {
		Required(
			field(1, "items", CollectionOf(TodoListItemRT), "Items"),
			field(2, "total", UInt64, "Total number of items", func() { Example(160) }),
			field(3, "offset", UInt64, "Offset", func() { Example(0) }),
		)
	})
})

var TodoCreatePayload = Type("TodoCreatePayload", func() {
	Required(todoFields("Payload", "create")...)
})
var TodoUpdatePayload = Type("TodoUpdatePayload", func() {
	Required(todoFields("Payload", "update")...)
})

var _ = Service("todos", func() {
	HTTP(func() {
		Path("/api/todos")
	})

	Method("list", func() {
		Result(TodoListRT)
		HTTP(func() {
			GET("")
			Response(StatusOK)
		})
		GRPC(func() {
		})
	})

	Method("show", func() {
		Result(TodoRT)
		Payload(func() { Required(field(1, "id", UInt64, "ID")) })

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
		GRPC(func() {
		})
	})

	Method("create", func() {
		Result(TodoRT)
		Payload(TodoCreatePayload)

		HTTP(func() {
			POST("")
			Response(StatusCreated)
		})
		GRPC(func() {
		})
	})

	Method("update", func() {
		Result(TodoRT)
		Payload(TodoUpdatePayload)

		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
		GRPC(func() {
		})
	})

	Method("delete", func() {
		Result(TodoRT)
		Payload(func() { Required(field(1, "id", UInt64, "ID")) })

		HTTP(func() {
			DELETE("/{id}")
			Response(StatusOK)
		})
		GRPC(func() {
		})
	})
})
