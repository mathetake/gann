package index

import (
	"encoding/json"
	"io/ioutil"

	"os"

	"github.com/pkg/errors"
)

// Save ... save index to disk as json format.
// NOTE: we do not include `items` and `nodes` (not exported) fields in Index struct,
// which means the loaded index cannot be rebuilt.
func (idx *Index) Save(path string) error {
	idxJSON, err := json.Marshal(idx)
	if err != nil {
		return errors.Wrap(err, "failed to json.marshal.")
	}

	err = ioutil.WriteFile(path, idxJSON, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "failed to ioutil.WriteFile")
	}
	return nil
}

// Load ... load index from disk
func (idx *Index) Load(path string) error {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "failed to ioutil.ReadFile.")
	}

	err = json.Unmarshal(raw, idx)
	if err != nil {
		return errors.Wrap(err, "failed to json.Unmarshal.")
	}

	// in order to prevent execution of build method on loaded index
	idx.isLoadedIndex = true
	return nil
}
