// Code generated by goa v3.13.2, DO NOT EDIT.
//
// apisvr HTTP client CLI support package
//
// Command:
// $ goa gen github.com/akm/goa_todo_example/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	todosc "github.com/akm/goa_todo_example/gen/http/todos/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
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
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		todosFlags = flag.NewFlagSet("todos", flag.ContinueOnError)

		todosListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		todosShowFlags  = flag.NewFlagSet("show", flag.ExitOnError)
		todosShowIDFlag = todosShowFlags.String("id", "REQUIRED", "ID")

		todosCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		todosCreateBodyFlag = todosCreateFlags.String("body", "REQUIRED", "")

		todosUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		todosUpdateBodyFlag = todosUpdateFlags.String("body", "REQUIRED", "")
		todosUpdateIDFlag   = todosUpdateFlags.String("id", "REQUIRED", "ID")

		todosDeleteFlags  = flag.NewFlagSet("delete", flag.ExitOnError)
		todosDeleteIDFlag = todosDeleteFlags.String("id", "REQUIRED", "ID")
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
			c := todosc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = todosc.BuildShowPayload(*todosShowIDFlag)
			case "create":
				endpoint = c.Create()
				data, err = todosc.BuildCreatePayload(*todosCreateBodyFlag)
			case "update":
				endpoint = c.Update()
				data, err = todosc.BuildUpdatePayload(*todosUpdateBodyFlag, *todosUpdateIDFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = todosc.BuildDeletePayload(*todosDeleteIDFlag)
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
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todos show -id UINT64

Show implements show.
    -id UINT64: ID

Example:
    %[1]s todos show --id 17973800868787948738
`, os.Args[0])
}

func todosCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todos create -body JSON

Create implements create.
    -body JSON: 

Example:
    %[1]s todos create --body '{
      "State": "open",
      "title": "Nesciunt itaque ratione quis id et blanditiis."
   }'
`, os.Args[0])
}

func todosUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todos update -body JSON -id UINT64

Update implements update.
    -body JSON: 
    -id UINT64: ID

Example:
    %[1]s todos update --body '{
      "body": {
         "State": "closed",
         "title": "Accusamus tempora dolores facere nulla."
      }
   }' --id 10738902383531522955
`, os.Args[0])
}

func todosDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todos delete -id UINT64

Delete implements delete.
    -id UINT64: ID

Example:
    %[1]s todos delete --id 13138609169822583145
`, os.Args[0])
}