package environment

import (
	"github.com/fiscusproject/fiscus/internal/core/commons"
)

var (
	HREnabled        = false
	HRFiscalEndpoint = ""
)

func Load() {
	commons.LoadBoolEnvVariable("FISCUS_HR_ENABLED", &HREnabled)
	commons.LoadStrEnvVariable("FISCUS_HR_FISCAL_ENDPOINT", &HRFiscalEndpoint)
}
