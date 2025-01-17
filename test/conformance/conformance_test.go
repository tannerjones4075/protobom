package conformance

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bom-squad/protobom/pkg/formats"
	"github.com/bom-squad/protobom/pkg/reader"
	"github.com/bom-squad/protobom/pkg/sbom"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestUnserializeFormats(t *testing.T) {
	for _, format := range formats.List {
		files := findFiles(t, format)
		r := reader.New()
		for _, fname := range files {
			sut, err := r.ParseFile(fname)
			require.NoError(t, err)
			golden := readProtobom(t, fname+".proto")
			t.Logf("sut: %s golden: %s", fname, fname+".proto")
			t.Run("testNodes", func(t *testing.T) { testNodes(t, golden, sut) })
		}
	}
}

func findFiles(t *testing.T, f formats.Format) []string {
	ret := []string{}
	dirName := filepath.Join("testdata", f.Type(), f.Version(), f.Encoding())
	files, err := os.ReadDir(dirName)
	if errors.Is(err, os.ErrNotExist) {
		return ret
	}
	require.NoError(t, err)
	for _, fsentry := range files {
		if strings.HasSuffix(fsentry.Name(), ".proto") {
			continue
		}
		ret = append(ret, filepath.Join(dirName, fsentry.Name()))
	}
	return ret
}

func testNodes(t *testing.T, golden, sut *sbom.Document) {
	require.Equal(t, len(golden.NodeList.Nodes), len(sut.NodeList.Nodes), "number of nodes")
	require.True(t, golden.NodeList.Equal(sut.NodeList), "equal nodelist", golden, sut)
}

func readProtobom(t *testing.T, path string) *sbom.Document {
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	bom := &sbom.Document{}
	if err := proto.Unmarshal(data, bom); err != nil {
		logrus.Fatal(fmt.Errorf("unmarshaling protobuf: %v", err))
	}
	return bom
}
