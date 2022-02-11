package client

import (
	"embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/module"
	"github.com/hashicorp/go-hclog"
)

//go:embed moduledata/*
var moduleData embed.FS

func ModuleInfo(logger hclog.Logger, module string, prefferedVersions []uint32) (resp module.InfoResponse, diags diag.Diagnostics) {
	for _, v := range prefferedVersions {
		fn := fmt.Sprintf("moduledata/%s/%v.hcl", module, v)
		data, err := moduleData.ReadFile(fn)
		if err != nil {
			continue
		}

		resp.Version = v
		resp.Info = map[string][]byte{
			"info": data,
		}
		break
	}
	if resp.Version == 0 {
		logger.Warn("received unsupported module info request", "module", module, "preferred_versions", prefferedVersions)
	}

	files, err := moduleData.ReadDir(fmt.Sprintf("moduledata/%s", module))
	if err != nil {
		return resp, diag.Diagnostics{diag.NewBaseError(err, diag.INTERNAL)}
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		vInt, err := strconv.ParseUint(strings.TrimSuffix(f.Name(), ".hcl"), 10, 32)
		if err != nil {
			continue
		}
		resp.OtherVersions = append(resp.OtherVersions, uint32(vInt))
	}

	return resp, nil
}
