package common

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-master-data/constanta"
	"go-master-data/model"
)

func GenerateI18NErrorMessage(err model.ErrorModel, language string) (output string) {
	defer func() {
		if r := recover(); r != nil {
			output = err.Error.Error()
		}
	}()

	if language == "" {
		language = constanta.DefaultApplicationsLanguage
	}

	localize := i18n.NewLocalizer(ErrorBundle, language)
	if err.ErrorParameter == nil {
		output = localize.MustLocalize(&i18n.LocalizeConfig{
			MessageID: err.Error.Error(),
		})
	} else {
		param := make(map[string]interface{})
		for i := 0; i < len(err.ErrorParameter); i++ {
			var parameterValue = err.ErrorParameter[i].ErrorParameterValue
			if err.ErrorParameter[i].ErrorParameterKey == "FieldName" {
				parameterValue = GenerateConstantaI18n(err.ErrorParameter[i].ErrorParameterValue, language, nil)
			}
			param[err.ErrorParameter[i].ErrorParameterKey] = parameterValue
			if param["RuleName"] != nil {
				param["RuleName"] = GenerateConstantaI18n(param["RuleName"].(string), language, nil)
			} else if param["Other"] != nil {
				param["Other"] = GenerateConstantaI18n(param["Other"].(string), language, nil)
			}

		}

		output = localize.MustLocalize(&i18n.LocalizeConfig{
			MessageID:    err.Error.Error(),
			TemplateData: param,
		})
	}

	return
}

func GenerateI18NServiceMessage(bundle *i18n.Bundle, messageID string, language string, param map[string]interface{}) (output string) {
	defer func() {
		if r := recover(); r != nil {
			output = messageID
		}
	}()

	if language == "" {
		language = constanta.DefaultApplicationsLanguage
	}

	localize := i18n.NewLocalizer(bundle, language)

	if param == nil {
		output = localize.MustLocalize(&i18n.LocalizeConfig{
			MessageID: messageID,
		})
	} else {
		output = localize.MustLocalize(&i18n.LocalizeConfig{
			MessageID:    messageID,
			TemplateData: param,
		})
	}
	return
}

func GenerateConstantaI18n(messageID string, language string, param map[string]interface{}) string {
	return GenerateI18NServiceMessage(ConstantaBundle, messageID, language, param)
}

func GenerateParams(params map[string]string) (result string) {
	i := 0
	for key, value := range params {
		i++
		result += key + "=" + value
		if i != len(params) {
			result += "&"
		}
	}
	return
}
