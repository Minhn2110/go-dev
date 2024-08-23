package main

import (
	"fmt"
	"strings"
)

//func main() {
//	const scriptContent = `
//		<script>
//			window.ce_setting={{.CESetting}}
//			window.ce_token="{{.CEToken}}";
//			window.ce_cartId="{{.CECartID}}";
//			window.ce_customer_id="{{.CECustomerID}}";
//			window.ce_tenant_id="{{.CETenantID}}";
//		</script>
//`
//
//	type ScriptCE struct {
//		CESetting    string `json:"ce_setting"`
//		CEToken      string `json:"ce_token"`
//		CECartID     string `json:"ce_cart_id"`
//		CECustomerID string `json:"ce_customer_id"`
//		CETenantID   string `json:"ce_tenant_id"`
//	}
//
//	data := ScriptCE{
//		CESetting:    "setting",
//		CEToken:      "{{settings.storefront_api.token}}",
//		CECartID:     "{{cart_id}}",
//		CECustomerID: "{{customer.id}}",
//		CETenantID:   "4f22c1dd-1c7e-4cd8-b42f-1466a97244c1",
//	}
//
//	tmpl, err := template.New("ceScript").Parse(scriptContent)
//	if err != nil {
//		panic(err)
//	}
//	var buf bytes.Buffer
//	err = tmpl.Execute(&buf, data)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("buf", buf.String())
//}

func main() {
	emailSlice := strings.Split("minhnha@smartosc.com,ngatdt@smartosc.com", ",")
	fmt.Println("emailSlice", emailSlice)
}
