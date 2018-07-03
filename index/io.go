package index

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"

	"os"

	"github.com/pkg/errors"
)

// Save ... save index to disk as json format.
// NOTE: we do not include `items` and `nodes` (not exported) fields in Index struct in its json output,
// which means the loaded index cannot be rebuilt.
func (idx *Index) Save(path string) error {
	idxJSON, err := json.Marshal(idx)
	if err != nil {
		return errors.Wrap(err, "failed to json.marshal.")
	}

	z, err := makeGzip(idxJSON)
	if err != nil {
		return errors.Wrap(err, "failed to makeGzip.")
	}

	err = ioutil.WriteFile(path, z, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "failed to ioutil.WriteFile")
	}
	return nil
}

// Load ... load index from disk
func (idx *Index) Load(path string) error {
	fi, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "failed to os.Open.")
	}

	defer fi.Close()

	fz, err := gzip.NewReader(fi)
	if err != nil {
		return errors.Wrap(err, "failed to gzip.NewReader")
	}

	defer fz.Close()

	raw, err := ioutil.ReadAll(fz)
	if err != nil {
		return errors.Wrap(err, "failed to ioutil.ReadAll")
	}

	err = json.Unmarshal(raw, idx)
	if err != nil {
		return errors.Wrap(err, "failed to json.Unmarshal.")
	}

	// in order to prevent execution of build method on loaded index
	idx.isLoadedIndex = true
	return nil
}

func makeGzip(body []byte) ([]byte, error) {
	var b bytes.Buffer
	err := func() error {
		gw := gzip.NewWriter(&b)
		defer gw.Close()

		if _, err := gw.Write(body); err != nil {
			return err
		}
		return nil
	}()
	return b.Bytes(), err
}
