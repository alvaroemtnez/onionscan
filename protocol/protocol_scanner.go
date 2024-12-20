package protocol

import (
	"github.com/alvaroemtnez/onionscan/config"
	"github.com/alvaroemtnez/onionscan/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
