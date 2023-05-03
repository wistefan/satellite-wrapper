package main

import (
	"encoding/json"
	"fmt"

	"github.com/fiware/VCVerifier/logging"
	"github.com/fiware/dsba-pdp/model"
	"github.com/valyala/fasthttp"
)

type SatelliteConfig struct {
	PartiesPath string
	model.AuthorizationRegistry
}

type IShareClient struct {
	satelliteConfig SatelliteConfig
}

func (isc *IShareClient) GetParty(accessToken string, partyId string) (party Party, err error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(fmt.Sprintf("%s%s/%s", isc.satelliteConfig.Host, isc.satelliteConfig.PartiesPath, partyId))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	client.Do(req, resp)

	err = json.Unmarshal(resp.Body(), &party)
	if err != nil {
		logging.Log().Warnf("Was not able to unmarshal the response. Err: %v", err)
		return
	}
	return
}

func (isc *IShareClient) GetParties(accessToken string) (parties Parties, err error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(fmt.Sprintf("%s%s", isc.satelliteConfig.Host, isc.satelliteConfig.PartiesPath))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	client.Do(req, resp)

	err = json.Unmarshal(resp.Body(), &parties)
	if err != nil {
		logging.Log().Warnf("Was not able to unmarshal the response. Err: %v", err)
		return
	}
	return
}

type PartiesInfo struct {
	TotalCount int64     `json:"total_count"`
	PageCount  int64     `json:"pageCount"`
	Count      int64     `json:"count"`
	Data       []Parties `json:"data"`
}

type Adherence struct {
	Status    string `json:"status"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type AdditionalInfo struct {
	Description  string `json:"description"`
	Logo         string `json:"logo"`
	Website      string `json:"website"`
	CompanyPhone string `json:"company_phone"`
	CompanyEmail string `json:"company_email"`
}

type Certificate struct {
	SubjectName     string `json:"subject_name"`
	CertificateType string `json:"certificate_type"`
	EnabledFrom     string `json:"enabled_from"`
	X5C             string `json:"x5c"`
	Fingerprint     string `json:"x5t#S256"`
}

type Party struct {
	PartyId        string         `json:"party_id"`
	PartyName      string         `json:"party_name"`
	CapabitlityUrl string         `json:"capability_url"`
	RegistrarId    string         `json:"registrar_id"`
	Adherence      Adherence      `json:"adherence"`
	AdditionalInfo AdditionalInfo `json:"additional_info"`
	Certificates   []Certificate  `json:"certificates"`
}

type Parties struct {
	Parties PartiesInfo `json:"parties_info"`
}
