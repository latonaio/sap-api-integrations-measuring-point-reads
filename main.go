package main

import (
	sap_api_caller "sap-api-integrations-measuring-point-reads/SAP_API_Caller"
	"sap-api-integrations-measuring-point-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Measuring_Point_Header_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata4/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "Equipment",
		}
	}

	caller.AsyncGetMeasuringPoint(
		inoutSDC.MeasuringPoint.MeasuringPoint,
		inoutSDC.MeasuringPoint.Equipment,
		accepter,
	)
}
