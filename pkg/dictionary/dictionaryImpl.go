package dictionary

import (
	"context"
	"fmt"
	"strings"

	"github.com/paper-indonesia/pg-mcp-server/constant"

	errors "github.com/paper-indonesia/pg-mcp-server/pkg/error"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func (d *dictionary) get(lang, key string) (s string) {
	tag := language.Indonesian

	if strings.EqualFold(lang, EN) {
		tag = language.English
	}

	p := message.NewPrinter(tag)
	return p.Sprintf(key)
}

func (d *dictionary) getMessage(lang, key string) string {
	if strings.EqualFold(lang, EN) {
		return d.printerEn.Sprintf(key)
	}

	return d.printerID.Sprintf(key)
}

func (d *dictionary) getPrinter(lang string) (printer *message.Printer) {
	if strings.EqualFold(lang, EN) {
		return d.printerEn
	}

	return d.printerID
}

func (d *dictionary) getLanguage(ctx context.Context) string {
	if ctx.Value(constant.CtxAcceptLanguage) != nil {
		return ctx.Value(constant.CtxAcceptLanguage).(string)
	}
	return ""
}

func (d *dictionary) SetDictionaryMessage(
	ctx context.Context,
	dict string,
	args ...interface{}) string {
	lang := IN
	if d.getLanguage(ctx) != "" {
		lang = d.getLanguage(ctx)
	}

	if dict == "" {
		return ""
	}

	errCode, _ := errors.ExtractError(fmt.Errorf(dict))
	if errCode != "" {
		dict = errCode
	}

	printer := d.getPrinter(lang)
	return printer.Sprintf(GetTranslationCode(dict), args...)
}
