package todo

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

/*
This is an example of a simple data store that uses a JSON file to persist.
Don't use this in production! Use database instead.
*/

const storeFilePath = "tmp/todos.json"

func Load() (Todos, error) {
	_, err := os.Stat(storeFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return Todos{}, nil
		}
		return nil, errors.Wrapf(err, "failed to stat %s", storeFilePath)
	}

	f, err := os.Open(storeFilePath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open %s", storeFilePath)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s", storeFilePath)
	}

	var todos Todos
	if err := json.Unmarshal(b, &todos); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal %s", storeFilePath)
	}

	return todos, nil
}

func Save(todos Todos) error {
	b, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return errors.Wrapf(err, "failed to marshal %s", storeFilePath)
	}

	if err := os.MkdirAll(filepath.Dir(storeFilePath), 0755); err != nil {
		return errors.Wrapf(err, "failed to mkdir %s", storeFilePath)
	}

	f, err := os.Create(storeFilePath)
	if err != nil {
		return errors.Wrapf(err, "failed to create %s", storeFilePath)
	}
	defer f.Close()

	if _, err := f.Write(b); err != nil {
		return errors.Wrapf(err, "failed to write %s", storeFilePath)
	}

	return nil
}
