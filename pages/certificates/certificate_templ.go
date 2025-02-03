// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package certificates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/henrriusdev/hangrite/models"
	"strings"
)

func getSpanishMonth(month string) string {
	months := map[string]string{
		"January":   "ENERO",
		"February":  "FEBRERO",
		"March":     "MARZO",
		"April":     "ABRIL",
		"May":       "MAYO",
		"June":      "JUNIO",
		"July":      "JULIO",
		"August":    "AGOSTO",
		"September": "SEPTIEMBRE",
		"October":   "OCTUBRE",
		"November":  "NOVIEMBRE",
		"December":  "DICIEMBRE",
	}
	return months[month]
}

func Certificate(player models.Player, certificate models.Certificate) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html><head><meta content=\"text/html; charset=UTF-8\" http-equiv=\"content-type\"><title>Constancia - ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(player.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/certificates/certificate.templ`, Line: 31, Col: 36}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</title><style type=\"text/css\">\n                @import url(https://themes.googleusercontent.com/fonts/css?kit=fpjTOVmNbO4Lz34iLyptLU5sOsqza4scPUlExctaWh-jUS0hn0umzC0XmKC5aH3C);\n                @page {\n                    size: letter;\n                    margin: 0;\n                }\n                body {\n                    background-color: #ffffff;\n                    max-width: 447.1pt;\n                    padding: 72pt 85pt 70.7pt 80pt;\n                    font-family: \"Times New Roman\";\n                    margin: 0 auto;\n                    position: relative;\n                }\n                .background-image {\n                    position: fixed;\n                    top: 50%;\n                    left: 50%;\n                    transform: translate(-50%, -50%);\n                    width: 406.47px;\n                    height: 332.53px;\n                    z-index: -1;\n                }\n                .header-logo {\n                    text-align: center;\n                    margin-bottom: 20px;\n                }\n                .header-logo img {\n                    width: 179.43px;\n                    height: 144.87px;\n                }\n                .title {\n                    text-align: center;\n                    font-size: 18pt;\n                    margin-bottom: 20px;\n                }\n                .content {\n                    text-align: justify;\n                    font-size: 13.5pt;\n                    margin: 20px 5pt;\n                    line-height: 1.5;\n                }\n                .footer {\n                    text-align: center;\n                    margin-top: 40px;\n                }\n                .footer-text {\n                    font-family: Arial;\n                    font-weight: bold;\n                    margin: 5px 0;\n                }\n                .footer-image {\n                    margin-top: 20px;\n                    text-align: center;\n                }\n                .footer-image img {\n                    width: 496.34px;\n                    height: 151.85px;\n                }\n                .bold {\n                    font-weight: bold;\n                }\n                .date {\n                    margin: 20px 5pt;\n                    font-size: 14pt;\n                }\n                .center-image {\n                    text-align: center;\n                    margin: 20px 0;\n                }\n                .center-image img {\n                    width: 406.47px;\n                    height: 332.53px;\n                }\n                @media print {\n                    body {\n                        -webkit-print-color-adjust: exact;\n                        print-color-adjust: exact;\n                    }\n                    #print-button {\n                        display: none;\n                    }\n                }\n            </style></head><body><img class=\"background-image\" alt=\"\" src=\"/assets/images/certificates/image3.jpg\"><div class=\"header-logo\"><img alt=\"\" src=\"/assets/images/certificates/image2.jpg\"></div><div class=\"title\">Constancia</div><div class=\"content\"><span>Quien Suscribe: </span> <span class=\"bold\">HENRRY JOSE BOURGEOT LISSIR</span> <span>, Titular de la Cédula de Identidad N° </span> <span class=\"bold\">V-8.599.738</span> <span>. Representante legal de la </span> <span class=\"bold\">ACADEMIA DE BEISBOL HANG RITE C.A</span> <span>, Ubicado en el ESTADIO DE SOROCAIMA II, CALLE MARIA CASTRO, Por medio de la presente hago constar que el Prospecto </span> <span class=\"bold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(strings.ToUpper(player.Name))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/certificates/certificate.templ`, Line: 131, Col: 53}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</span> <span>, portador de la Cédula de Identidad <span class=\"bold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(player.IDNumber)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/certificates/certificate.templ`, Line: 132, Col: 84}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</span>. Se desempeña en esta Academia como (")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(strings.ToUpper(player.Position))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/certificates/certificate.templ`, Line: 132, Col: 166}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "), En el Horario Comprendido entre ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(certificate.TrainingHours)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/certificates/certificate.templ`, Line: 132, Col: 230}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, ". </span></div><div class=\"date\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if spanishMonth := getSpanishMonth(certificate.IssueDate.Format("January")); spanishMonth != "" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "Constancia que se expide a petición de la parte interesada en Maracay, en fecha ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(certificate.IssueDate.Format("2"))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/certificates/certificate.templ`, Line: 136, Col: 121}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, " de ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(spanishMonth)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/certificates/certificate.templ`, Line: 136, Col: 141}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, " del ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(certificate.IssueDate.Format("2006"))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/certificates/certificate.templ`, Line: 136, Col: 186}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, ".")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "</div><div class=\"footer\"><div class=\"footer-text\">HENRRY JOSE BOURGEOT LISSIR</div><div class=\"footer-text\">C.I: N° V-8.599.738</div><div class=\"footer-text\">Celular: 0412-8908663</div></div><div class=\"footer-image\"><img alt=\"\" src=\"/assets/images/certificates/image1.png\"></div><button id=\"print-button\" style=\"position: fixed; top: 20px; right: 20px; padding: 10px 20px; background: #007bff; color: white; border: none; border-radius: 5px; cursor: pointer;\" onclick=\"window.print()\">Imprimir / Guardar PDF</button><script>\n\t\t\t\t\t// Auto-focus the print button\n\t\t\t\t\tdocument.getElementById('print-button').focus();\n\t\t\t</script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
