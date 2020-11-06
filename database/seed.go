package database

import "strings"

func Seed() error {

	chime := makeSponsor(
		"Chime",
		Platinum,
		"766469052549169155",
		strings.Join([]string{SponsorRole, PlatinumRole}, ","),
	)
	result := db.Create(&chime)
	if result.Error != nil {
		return result.Error
	}

	crowdstrike := makeSponsor(
		"Crowdstrike",
		Platinum,
		"765741926879330355",
		strings.Join([]string{SponsorRole, PlatinumRole}, ","),
	)
	crowdstrike.TicketCodes = append(crowdstrike.TicketCodes, TicketCode{TicketCode: "PlatCRDSTK90320"})
	result = db.Create(&crowdstrike)
	if result.Error != nil {
		return result.Error
	}

	google := makeSponsor(
		"Google",
		Platinum,
		"765742596340580354",
		strings.Join([]string{SponsorRole, PlatinumRole}, ","),
	)

	google.TicketCodes = append(google.TicketCodes, TicketCode{TicketCode: "PlatGGLE90320"})
	result = db.Create(&google)
	if result.Error != nil {
		return result.Error
	}

	jetbrains := makeSponsor(
		"JetBrains",
		Platinum,
		"765742657284079619",
		strings.Join([]string{SponsorRole, PlatinumRole}, ","),
	)

	result = db.Create(&jetbrains)
	if result.Error != nil {
		return result.Error
	}

	lightstep := makeSponsor(
		"Lightstep",
		Platinum,
		"765743315462127618",
		strings.Join([]string{SponsorRole, PlatinumRole}, ","),
	)

	lightstep.TicketCodes = append(lightstep.TicketCodes, TicketCode{TicketCode: "PlatLGHTSP90320"})
	result = db.Create(&lightstep)
	if result.Error != nil {
		return result.Error
	}

	microsoft := makeSponsor(
		"Microsoft",
		Platinum,
		"765742716357050388",
		strings.Join([]string{SponsorRole, PlatinumRole}, ","),
	)
	microsoft.TicketCodes = append(microsoft.TicketCodes, TicketCode{TicketCode: "PlatMCSFT90320"})
	result = db.Create(&microsoft)
	if result.Error != nil {
		return result.Error
	}

	// GOLD

	onepasswd := makeSponsor(
		"1Password",
		Gold,
		"765742780672114698",
		strings.Join([]string{SponsorRole, GoldRole}, ","),
	)

	onepasswd.TicketCodes = append(onepasswd.TicketCodes, TicketCode{TicketCode: "GldPSWRD90320"})
	result = db.Create(&onepasswd)
	if result.Error != nil {
		return result.Error
	}

	capone := makeSponsor(
		"Capital One",
		Gold,
		"765742819557113896",
		strings.Join([]string{SponsorRole, GoldRole}, ","),
	)

	capone.TicketCodes = append(capone.TicketCodes, TicketCode{TicketCode: "GldCPON90320"})

	result = db.Create(&capone)
	if result.Error != nil {
		return result.Error
	}

	cockroach := makeSponsor(
		"Cockroach DB",
		Gold,
		"765742926244347935",
		strings.Join([]string{SponsorRole, GoldRole}, ","),
	)

	result = db.Create(&cockroach)
	if result.Error != nil {
		return result.Error
	}

	fullstory := makeSponsor(
		"FullStory",
		Gold,
		"765742955302223882",
		strings.Join([]string{SponsorRole, GoldRole}, ","),
	)

	fullstory.TicketCodes = append(fullstory.TicketCodes, TicketCode{TicketCode: "GldFLSTY90320"})
	result = db.Create(&fullstory)
	if result.Error != nil {
		return result.Error
	}

	mercari := makeSponsor(
		"Mercari",
		Gold,
		"765946614983426049",
		strings.Join([]string{SponsorRole, GoldRole}, ","),
	)

	mercari.TicketCodes = append(mercari.TicketCodes, TicketCode{TicketCode: "GldMRCR90320"})
	result = db.Create(&mercari)
	if result.Error != nil {
		return result.Error
	}

	salesforce := makeSponsor(
		"Salesforce",
		Gold,
		"765946722164801556",
		strings.Join([]string{SponsorRole, GoldRole}, ","),
	)
	salesforce.TicketCodes = append(salesforce.TicketCodes, TicketCode{TicketCode: "GldSLFRC90320"})
	result = db.Create(&salesforce)
	if result.Error != nil {
		return result.Error
	}

	split := makeSponsor(
		"Split",
		Gold,
		"765946801936400456",
		strings.Join([]string{SponsorRole, GoldRole}, ","),
	)
	result = db.Create(&split)
	if result.Error != nil {
		return result.Error
	}

	synopsys := makeSponsor(
		"Synopsys",
		Gold,
		"765946888976728115",
		strings.Join([]string{SponsorRole, GoldRole}, ","),
	)
	synopsys.TicketCodes = append(synopsys.TicketCodes, TicketCode{TicketCode: "GldSNPS90320"})
	result = db.Create(&synopsys)
	if result.Error != nil {
		return result.Error
	}

	// Silver

	orijtech := makeSponsor(
		"orijtech",
		Silver,
		"771077174421094441",
		strings.Join([]string{SponsorRole, SilverRole}, ","),
	)
	orijtech.TicketCodes = append(orijtech.TicketCodes, TicketCode{TicketCode: "SlvrOJTCH90320"})
	result = db.Create(&orijtech)
	if result.Error != nil {
		return result.Error
	}

	sourcegraph := makeSponsor(
		"sourcegraph",
		Silver,
		"765947250681577523",
		strings.Join([]string{SponsorRole, SilverRole}, ","),
	)

	sourcegraph.TicketCodes = append(sourcegraph.TicketCodes, TicketCode{TicketCode: "SlvrSRGRPH90320"})
	result = db.Create(&sourcegraph)
	if result.Error != nil {
		return result.Error
	}

	// bronze
	sonobi := makeSponsor(
		"sonobi",
		Bronze,
		"765947344096067645",
		strings.Join([]string{SponsorRole, BronzeRole}, ","),
	)

	sonobi.TicketCodes = append(sonobi.TicketCodes, TicketCode{TicketCode: "BrnzSNBI90320"})
	result = db.Create(&sonobi)
	if result.Error != nil {
		return result.Error
	}

	/*sm := &ScheduledMessage{
		TargetChannel: "765946614983426049",
		Message:       "Scheduled Message for Mercari",
		Cron:          "* * * * *",
	}
	result = db.Create(&sm)
	if result.Error != nil {
		return result.Error
	} */

	return nil

}

func makeSponsor(name, level, channel, roles string) Sponsor {
	s := Sponsor{
		Name:    name,
		Level:   level,
		Channel: channel,
		Roles:   roles,
	}
	return s
}
