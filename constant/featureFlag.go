package constant

import (
	"encoding/json"
	"fmt"
	"strings"

	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/ffcontext"
)

const (
	SnapCoreBankTransferRouting = "snap-core-bank-transfer-routing"
)

type VirtualAccountCustomLogic struct {
	LogicName string `json:"logicName"`
	Data      any    `json:"data"`
}

func GetVirtualAccountCustomLogicFromFF(vaNumber string) (*VirtualAccountCustomLogic, bool) {
	userFlag := ffcontext.NewEvaluationContext(vaNumber)
	dataJson, err := ffclient.JSONVariation("snap-core-virtual-account-custom-logic", userFlag, nil)
	if err != nil {
		return nil, false
	}

	if dataJson == nil {
		return nil, false
	}

	dataBytes, err := json.Marshal(dataJson)
	if err != nil {
		return nil, false
	}

	var result VirtualAccountCustomLogic
	err = json.Unmarshal(dataBytes, &result)
	if err != nil {
		return nil, false
	}

	return &result, true
}

type TransferActions struct {
	Intrabank string `json:"intrabank"`
	Interbank string `json:"interbank"`
	BIFast    string `json:"bifast"`
}

func GetCustomRemarkRules(bank string) *TransferActions {
	flagKey := "snap-core-custom-remark-rules"
	bank = strings.ToLower(bank)
	bankFlag := ffcontext.NewEvaluationContext(bank)
	var result TransferActions

	dataJson, err := ffclient.JSONVariation(flagKey, bankFlag, nil)
	if err != nil || dataJson == nil {
		return &result
	}

	if intrabank, ok := dataJson["intrabank"]; ok {
		result.Intrabank = fmt.Sprintf("%v", intrabank)
	}
	if interbank, ok := dataJson["interbank"]; ok {
		result.Interbank = fmt.Sprintf("%v", interbank)
	}
	if bifast, ok := dataJson["bifast"]; ok {
		result.BIFast = fmt.Sprintf("%v", bifast)
	}
	return &result
}

func EnableSlackLoggerFlag() bool {
	key := "snap-core-slack-logger-notifier"
	queryContext := ffcontext.NewEvaluationContext("slack-logger-status")
	isActive, err := ffclient.BoolVariation(key, queryContext, false)
	if err != nil {
		return false
	}

	return isActive
}

func IsUseCB() bool {
	userFlag := ffcontext.NewEvaluationContext("transfer-use-cb")
	if isActive, err := ffclient.BoolVariation(SnapCoreBankTransferRouting, userFlag, false); err != nil {
		return false
	} else if isActive {
		return true
	}
	return false
}
