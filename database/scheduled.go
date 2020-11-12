package database

const (
	chimeChannel       = "766469052549169155"
	cockroachChannel   = "765742926244347935"
	crowdstrikeChannel = "765741926879330355"
	fullstoryChannel   = "765742955302223882"
	googleChannel      = "765742596340580354"
	lightstepChannel   = "765743315462127618"
	microsoftChannel   = "765742716357050388"
	splitChannel       = "765946801936400456"
	synopsysChannel    = "765946888976728115"
)

func Scheduled() error {

	sm := &ScheduledMessage{
		TargetChannel: chimeChannel,
		SponsorID:     1,
		Message:       "Hi, there! 👋  Thanks for popping into our channel. If you haven't talked with us yet - we're Chime(https://chime.com), and we'd love to get to know you! We're away from the chat right now, and will be back around 4:00pm EST, so be sure to come back and Chime in with us then! Want to read more about us before we really get to know each other? Check out our blog, Life at Chime(https://medium.com/life-at-chime)!😎",
		Cron:          "0 20 9-13 * *",
	}
	result := db.Create(&sm)
	if result.Error != nil {
		return result.Error
	}
	sm2 := &ScheduledMessage{
		TargetChannel: chimeChannel,
		SponsorID:     1,
		Message:       "Hi, there! 👋  Thanks for popping into our channel. If you haven't talked with us yet - we're Chime(https://chime.com), and we'd love to get to know you! We're away from the chat right now, and will be back around 12:00pm EST, so be sure to come back and Chime in with us then! Want to read more about us before we really get to know each other? Check out our blog, Life at Chime!(https://medium.com/life-at-chime)😎",
		Cron:          "0 2 10-14 * *",
	}
	result = db.Create(&sm2)
	if result.Error != nil {
		return result.Error
	}

	sm3 := &ScheduledMessage{
		TargetChannel: chimeChannel,
		SponsorID:     1,
		Message:       "Hi, there! 👋  Thanks for popping into our channel. If you haven't talked with us yet - we're Chime(https://chime.com), and we'd love to get to know you! We're away from the chat right now, and will be back around 12:00pm EST, so be sure to come back and Chime in with us then! Want to read more about us before we really get to know each other? Check out our blog, Life at Chime!(https://medium.com/life-at-chime)😎",
		Cron:          "0 7 10-14 * *",
	}
	result = db.Create(&sm3)
	if result.Error != nil {
		return result.Error
	}

	sm4 := &ScheduledMessage{
		TargetChannel: chimeChannel,
		SponsorID:     1,
		Message:       "Hi, there! 👋  Thanks for popping into our channel. If you haven't talked with us yet - we're Chime(https://chime.com), and we'd love to get to know you! We're away from the chat right now, and will be back around 12:00pm EST, so be sure to come back and Chime in with us then! Want to read more about us before we really get to know each other? Check out our blog, Life at Chime!(https://medium.com/life-at-chime)😎",
		Cron:          "0 12 10-14 * *",
	}
	result = db.Create(&sm4)
	if result.Error != nil {
		return result.Error
	}

	sm5 := &ScheduledMessage{
		TargetChannel: cockroachChannel,
		SponsorID:     9,
		Message:       "Ah, bummer! Sorry, we aren't online. Feel free to message in the chat and we will get to it as soon as we can. Enjoy GoVirCon! ",
		Cron:          "0 21,4,21 10-14 * *",
	}
	result = db.Create(&sm5)
	if result.Error != nil {
		return result.Error
	}
	sm6 := &ScheduledMessage{
		TargetChannel: crowdstrikeChannel,
		SponsorID:     2,
		Message:       "Ah, bummer! Sorry, we aren't online. Feel free to message in the chat and we will get to it as soon as we can. Enjoy GoVirCon! ",
		Cron:          "0 15,21 10-14 * *",
	}
	result = db.Create(&sm6)
	if result.Error != nil {
		return result.Error
	}

	sm7 := &ScheduledMessage{
		TargetChannel: crowdstrikeChannel,
		SponsorID:     2,
		Message:       "Hi Gophers! Don’t miss our session on Untangling the Monorepo. Catch it at our VMS today at 2:55pm PST.",
		Cron:          "0 15,19 10-14 * *",
	}
	result = db.Create(&sm7)
	if result.Error != nil {
		return result.Error
	}

	sm8 := &ScheduledMessage{
		TargetChannel: crowdstrikeChannel,
		SponsorID:     2,
		Message:       "Sorry we missed you! Leave us a message and we’ll get back to you. In the meantime, you can learn more about our team and culture on our portal: https://www.gophercon.com/page/1623781/engage-crowdstrike",
		Cron:          "0 23,6 10-14 * *",
	}
	result = db.Create(&sm8)
	if result.Error != nil {
		return result.Error
	}

	sm9 := &ScheduledMessage{
		TargetChannel: fullstoryChannel,
		SponsorID:     10,
		Message:       "FullStorians have signed off for the day but we can't wait to reply to you when we're back online tomorrow morning",
		Cron:          "0 21,6 10-14 * *",
	}
	result = db.Create(&sm9)
	if result.Error != nil {
		return result.Error
	}

	sm10 := &ScheduledMessage{
		TargetChannel: googleChannel,
		SponsorID:     3,
		Message:       "Sorry we missed you. Please leave us a message and we will respond when we return. Check out our Q&A schedule too - we'd love to chat with you in our virtual meeting space!",
		Cron:          "0 15,18,22 10-14 * *",
	}
	result = db.Create(&sm10)
	if result.Error != nil {
		return result.Error
	}

	sm11 := &ScheduledMessage{
		TargetChannel: googleChannel,
		SponsorID:     3,
		Message:       "Sorry we missed you. Please leave us a message and we will respond when we return. Check out our Q&A schedule too - we'd love to chat with you in our virtual meeting space!g",
		Cron:          "0 14,19,20,2 10-14 * *",
	}
	result = db.Create(&sm11)
	if result.Error != nil {
		return result.Error
	}

	sm12 := &ScheduledMessage{
		TargetChannel: googleChannel,
		SponsorID:     3,
		Message:       "Sorry we missed you. Please leave us a message and we will respond when we return. Check out our Q&A schedule too - we'd love to chat with you in our virtual meeting space!g",
		Cron:          "45 21 12 * *",
	}
	result = db.Create(&sm12)
	if result.Error != nil {
		return result.Error
	}

	sm13 := &ScheduledMessage{
		TargetChannel: googleChannel,
		SponsorID:     3,
		Message:       "Sorry we missed you. Please leave us a message and we will respond when we return. Check out our Q&A schedule too - we'd love to chat with you in our virtual meeting space!g",
		Cron:          "50 20 13 * *",
	}
	result = db.Create(&sm13)
	if result.Error != nil {
		return result.Error
	}

	sm14 := &ScheduledMessage{
		TargetChannel: googleChannel,
		SponsorID:     3,
		Message:       "Sorry we missed you. Please leave us a message and we will respond when we return. Check out our Q&A schedule too - we'd love to chat with you in our virtual meeting space!g",
		Cron:          "0 18 13 * *",
	}
	result = db.Create(&sm14)
	if result.Error != nil {
		return result.Error
	}

	sm15 := &ScheduledMessage{
		TargetChannel: lightstepChannel,
		SponsorID:     5,
		Message: `We are sorry we missed you. If you'd like, you can leave a message and we'll respond as soon as our team is back online. In the meantime, if you'd like to learn more about Lightstep or OpenTelemetry, we recommend checking out one of these resources: 
OpenTelemetry Docs: https://opentelemetry.lightstep.com/ 
How Lightstep works: https://lightstep.com/how-it-works 
Lyft Case Study: https://lightstep.com/case-studies/lyft 
Plans and pricing: https://lightstep.com/pricing `,
		Cron: "0 23,3,7 10-14 * *",
	}
	result = db.Create(&sm15)
	if result.Error != nil {
		return result.Error
	}

	sm16 := &ScheduledMessage{
		TargetChannel: microsoftChannel,
		SponsorID:     6,
		Message:       "thank you for stopping by! our team is away from discord at the moment but leave us a message and we will get back to you when we return.",
		Cron:          "0 0,4,8,12 10-14 * *",
	}
	result = db.Create(&sm16)
	if result.Error != nil {
		return result.Error
	}

	sm17 := &ScheduledMessage{
		TargetChannel: splitChannel,
		SponsorID:     13,
		Message:       "Thank you for stopping by. The Split team is away at the moment, but leave us a message and we will reply when we return.",
		Cron:          "0 0,5,10 10-14 * *",
	}
	result = db.Create(&sm17)
	if result.Error != nil {
		return result.Error
	}

	sm18 := &ScheduledMessage{
		TargetChannel: synopsysChannel,
		SponsorID:     14,
		Message:       "Thanks for visiting our channel! We will be available during the main programming days, Wednesday – Friday. Drop us a comment and we’ll be in touch soon! In the meantime, if you’d like to learn more about Synopsys, please visit https://www.synopsys.com/software-integrity.html",
		Cron:          "0 14 10,11 * *",
	}
	result = db.Create(&sm18)
	if result.Error != nil {
		return result.Error
	}

	sm19 := &ScheduledMessage{
		TargetChannel: synopsysChannel,
		SponsorID:     14,
		Message:       "We are sorry we missed you. Please leave us a message and we will respond when we return",
		Cron:          "30 21,1,4,10 11-13 * *",
	}
	result = db.Create(&sm19)
	if result.Error != nil {
		return result.Error
	}

	return nil

}
