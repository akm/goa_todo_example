// Code generated by goa v3.13.2, DO NOT EDIT.
//
// apisvr gRPC client CLI support package
//
// Command:
// $ goa gen github.com/akm/goa_todo_example/design

package cli

import (
	"flag"
	"fmt"
	"os"

	todosc "github.com/akm/goa_todo_example/gen/grpc/todos/client"
	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `todos (list|show|create|update|delete)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` todos list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, any, error) {
	var (
		todosFlags = flag.NewFlagSet("todos", flag.ContinueOnError)

		todosListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		todosShowFlags       = flag.NewFlagSet("show", flag.ExitOnError)
		todosShowMessageFlag = todosShowFlags.String("message", "", "")

		todosCreateFlags       = flag.NewFlagSet("create", flag.ExitOnError)
		todosCreateMessageFlag = todosCreateFlags.String("message", "", "")

		todosUpdateFlags       = flag.NewFlagSet("update", flag.ExitOnError)
		todosUpdateMessageFlag = todosUpdateFlags.String("message", "", "")

		todosDeleteFlags       = flag.NewFlagSet("delete", flag.ExitOnError)
		todosDeleteMessageFlag = todosDeleteFlags.String("message", "", "")
	)
	todosFlags.Usage = todosUsage
	todosListFlags.Usage = todosListUsage
	todosShowFlags.Usage = todosShowUsage
	todosCreateFlags.Usage = todosCreateUsage
	todosUpdateFlags.Usage = todosUpdateUsage
	todosDeleteFlags.Usage = todosDeleteUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "todos":
			svcf = todosFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "todos":
			switch epn {
			case "list":
				epf = todosListFlags

			case "show":
				epf = todosShowFlags

			case "create":
				epf = todosCreateFlags

			case "update":
				epf = todosUpdateFlags

			case "delete":
				epf = todosDeleteFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "todos":
			c := todosc.NewClient(cc, opts...)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = todosc.BuildShowPayload(*todosShowMessageFlag)
			case "create":
				endpoint = c.Create()
				data, err = todosc.BuildCreatePayload(*todosCreateMessageFlag)
			case "update":
				endpoint = c.Update()
				data, err = todosc.BuildUpdatePayload(*todosUpdateMessageFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = todosc.BuildDeletePayload(*todosDeleteMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// todosUsage displays the usage of the todos command and its subcommands.
func todosUsage() {
	fmt.Fprintf(os.Stderr, `Service is the todos service interface.
Usage:
    %[1]s [globalflags] todos COMMAND [flags]

COMMAND:
    list: List implements list.
    show: Show implements show.
    create: Create implements create.
    update: Update implements update.
    delete: Delete implements delete.

Additional help:
    %[1]s todos COMMAND --help
`, os.Args[0])
}
func todosListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todos list

List implements list.

Example:
    %[1]s todos list
`, os.Args[0])
}

func todosShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todos show -message JSON

Show implements show.
    -message JSON: 

Example:
    %[1]s todos show --message '{
      "id": 13511584583745042765
   }'
`, os.Args[0])
}

func todosCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todos create -message JSON

Create implements create.
    -message JSON: 

Example:
    %[1]s todos create --message '{
      "State": "open",
      "title": "Et laudantium aut culpa vel magnam vitae."
   }'
`, os.Args[0])
}

func todosUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todos update -message JSON

Update implements update.
    -message JSON: 

Example:
    %[1]s todos update --message '{
      "State": "open",
      "id": 12782481505029911788,
      "title": "Vero aut perspiciatis nesciunt accusamus."
   }'
`, os.Args[0])
}

func todosDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todos delete -message JSON

Delete implements delete.
    -message JSON: 

Example:
    %[1]s todos delete --message '{
      "id": 11495435370355398098
   }'
`, os.Args[0])
}
