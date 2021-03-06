// +build slow

package photoprism

import (
	"testing"

	"github.com/photoprism/photoprism/internal/test"
)

func TestCreateThumbnailsFromOriginals(t *testing.T) {
	conf := test.NewConfig()

	conf.CreateDirectories()

	conf.InitializeTestData(t)

	tensorFlow := NewTensorFlow(conf.TensorFlowModelPath())

	indexer := NewIndexer(conf.OriginalsPath(), tensorFlow, conf.Db())

	converter := NewConverter(conf.DarktableCli())

	importer := NewImporter(conf.OriginalsPath(), indexer, converter)

	importer.ImportPhotosFromDirectory(conf.ImportPath())

	CreateThumbnailsFromOriginals(conf.OriginalsPath(), conf.ThumbnailsPath(), 600, false)

	CreateThumbnailsFromOriginals(conf.OriginalsPath(), conf.ThumbnailsPath(), 300, true)
}
