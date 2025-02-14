package makefile

import (
	"context"
	"github.com/murphysecurity/murphysec/model"
	"github.com/murphysecurity/murphysec/utils"
)

type Inspector struct{}

func (i Inspector) String() string {
	return "MakefileInspector"
}

func (i Inspector) CheckDir(dir string) bool {
	return utils.IsFile("Makefile")
}

func (i Inspector) InspectProject(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (i Inspector) SupportFeature(feature model.InspectorFeature) bool {
	return model.InspectorFeatureAllowNested&feature > 0
}
