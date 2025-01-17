package serializers

import (
	"testing"

	"github.com/CycloneDX/cyclonedx-go"
	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/bom-squad/protobom/pkg/sbom"
	"github.com/stretchr/testify/require"
)

func TestComponentType(t *testing.T) {
	sut := CDX{}
	node := &sbom.Node{}

	for s, tc := range map[string]struct {
		prepare  func(*sbom.Node)
		compType cyclonedx.ComponentType
	}{
		"node file":              {func(n *sbom.Node) { n.PrimaryPurpose = "file"; n.Type = sbom.Node_FILE }, cyclonedx.ComponentTypeFile},
		"node file, ne purpose":  {func(n *sbom.Node) { n.PrimaryPurpose = "library"; n.Type = sbom.Node_FILE }, cyclonedx.ComponentTypeFile},
		"application":            {func(n *sbom.Node) { n.PrimaryPurpose = "application"; n.Type = sbom.Node_PACKAGE }, cdx.ComponentTypeApplication},
		"container":              {func(n *sbom.Node) { n.PrimaryPurpose = "container"; n.Type = sbom.Node_PACKAGE }, cdx.ComponentTypeContainer},
		"data":                   {func(n *sbom.Node) { n.PrimaryPurpose = "data"; n.Type = sbom.Node_PACKAGE }, cdx.ComponentTypeData},
		"device":                 {func(n *sbom.Node) { n.PrimaryPurpose = "device"; n.Type = sbom.Node_PACKAGE }, cdx.ComponentTypeDevice},
		"device-driver":          {func(n *sbom.Node) { n.PrimaryPurpose = "device-driver"; n.Type = sbom.Node_PACKAGE }, cdx.ComponentTypeDeviceDriver},
		"node package, pp file":  {func(n *sbom.Node) { n.PrimaryPurpose = "file"; n.Type = sbom.Node_PACKAGE }, cdx.ComponentTypeFile},
		"firmware":               {func(n *sbom.Node) { n.PrimaryPurpose = "firmware"; n.Type = sbom.Node_PACKAGE }, cdx.ComponentTypeFirmware},
		"framework":              {func(n *sbom.Node) { n.PrimaryPurpose = "framework"; n.Type = sbom.Node_PACKAGE }, cdx.ComponentTypeFramework},
		"library":                {func(n *sbom.Node) { n.PrimaryPurpose = "library"; n.Type = sbom.Node_PACKAGE }, cyclonedx.ComponentTypeLibrary},
		"machine-learning-model": {func(n *sbom.Node) { n.PrimaryPurpose = "machine-learning-model"; n.Type = sbom.Node_PACKAGE }, cdx.ComponentTypeMachineLearningModel},
		"operating-system":       {func(n *sbom.Node) { n.PrimaryPurpose = "operating-system"; n.Type = sbom.Node_PACKAGE }, cdx.ComponentTypeOS},
		"platform":               {func(n *sbom.Node) { n.PrimaryPurpose = "platform"; n.Type = sbom.Node_PACKAGE }, cdx.ComponentTypePlatform},
	} {
		tc.prepare(node)
		comp := sut.nodeToComponent(node)
		require.Equal(t, comp.Type, tc.compType, s)
	}
}
