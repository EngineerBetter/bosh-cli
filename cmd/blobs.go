package cmd

import (
	boshreldir "github.com/cloudfoundry/bosh-init/releasedir"
	boshui "github.com/cloudfoundry/bosh-init/ui"
	boshtbl "github.com/cloudfoundry/bosh-init/ui/table"
)

type BlobsCmd struct {
	blobsDir boshreldir.BlobsDir
	ui       boshui.UI
}

func NewBlobsCmd(blobsDir boshreldir.BlobsDir, ui boshui.UI) BlobsCmd {
	return BlobsCmd{blobsDir: blobsDir, ui: ui}
}

func (c BlobsCmd) Run() error {
	blobs, err := c.blobsDir.Blobs()
	if err != nil {
		return err
	}

	table := boshtbl.Table{
		Content: "blobs",

		Header: []string{"Path", "Size", "Blobstore ID", "SHA1"},

		SortBy: []boshtbl.ColumnSort{
			{Column: 0, Asc: true},
		},
	}

	for _, blob := range blobs {
		blobID := blob.BlobstoreID

		if len(blobID) == 0 {
			blobID = "(local)"
		}

		table.Rows = append(table.Rows, []boshtbl.Value{
			boshtbl.ValueString{blob.Path},
			boshtbl.ValueBytes{uint64(blob.Size)},
			boshtbl.ValueString{blobID},
			boshtbl.ValueString{blob.SHA1},
		})
	}

	c.ui.PrintTable(table)

	return nil
}
