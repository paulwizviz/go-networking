package client

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func Example_soapCall() {
	handler := func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			fmt.Println("Method is POST")
		}
		fmt.Println(req.Header)
		body, _ := io.ReadAll(req.Body)
		fmt.Printf("Request Body: %s", body)
		rw.WriteHeader(http.StatusOK)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	u := fmt.Sprintf("%s/%s", server.URL, "hello")

	sBody := SoapEnv{
		XMLNSSoapEnv: wsseXMLNS,
	}

	xmlBody, _ := xml.Marshal(sBody)

	client := &SOAPClient{}
	client.SoapCall(u, xmlBody)

	// Output:
	// Method is POST
	// map[Accept:[text/xml] Accept-Encoding:[gzip] Content-Length:[175] Content-Type:[text/xml; charset=utf-8] Soapaction:[http://schemas.xmlsoap.org/soap/envelope/] User-Agent:[Go-http-client/1.1]]
	// Request Body: <soap:Envelope xmlns:soapenv="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"><soapenv:Header></soapenv:Header><soap></soap></soap:Envelope>

}
