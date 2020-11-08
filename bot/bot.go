package bot

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/bketelsen/gcbot/database"
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
)

const (
	LAB_GHA            = "Complimentary Conference Lab: Get Go-ing with GitHub Actions"
	TWENTY_PERPET      = "One-Time 20%, plus Perpetual 10% Promo"
	COMMUNITY          = "Community Ticket"
	WORKSHOP_NLG       = "Pre-Conference Workshop: Next Level Go"
	WORKSHOP_DEBUG     = "Pre-Conference Workshop: Debugging Techniques for Go"
	WORKSHOP_ML        = "Pre-Conference Workshop: Infrastructure for ML Applications"
	WORKSHOP_CONCUR    = "Pre-Conference Workshop: Introduction to Concurrency in Go"
	WORKSHOP_SLS       = "Pre-Conference Workshop: Serverless Go"
	WORKSHOP_DGRAPH    = "Pre-Conference Workshop: Ultimate Dgraph with GraphQL and Go"
	ONETIME_TWENTY     = "One-Time 20% Promo"
	G2A_GROUP          = "G2A Grp Ticket"
	FRIENDS            = "Perpetual Promo: Friends of GopherCon"
	GEN_EARLY          = "General Admission - Early Gopher"
	GEN_SELF           = "General Admission - Self-Paid"
	WORKSHOP_PM_MICRO  = "P.M.Workshop: How to Structure Your Microservices"
	WORKSHOP_AM_CONCUR = "A.M. Workshop: Introduction to Concurrency in Go"
	WORKSHOP_AM_ML     = "A.M. Workshop: Infrastructure for ML Applications"
	WORKSHOP_PM_CONCUR = "P.M. Workshop: Introduction to Concurrency in Go"
	WORKSHOP_PM_SLS    = "P.M. Workshop: Serverless Go"
	WORKSHOP_PM_ML     = "P.M. Workshop: Infrastructure for ML Applications"
	WORKSHOP_AM_MICRO  = "A.M.Workshop: How to Structure Your Microservices"
	WORKSHOP_AM_DEBUG  = "A.M. Workshop: Debugging Techniques for Go"
	GEN_CORP           = "General Admission - Corporate"
	WORKSHOP_PM_K8S    = "P.M. Workshop: Go & Kubernetes in the Real World"
	WORKSHOP_AM_SLS    = "A.M. Workshop: Serverless Go"
	WORKSHOP_PM_DEBUG  = "P.M. Workshop: Debugging Techniques for Go"
	WORKSHOP_AM_K8S    = "A.M. Workshop: Go & Kubernetes in the Real World"
	WORKSHOP_PM_PACMAN = "P.M. Workshop: Pac-Man from Scratch"
	WORKSHOP_AM_PACMAN = "A.M. Workshop: Pac-Man from Scratch"
)

var TicketRole = map[string]string{
	LAB_GHA:            "WorkshopGithub",
	WORKSHOP_NLG:       "WorkshopNLG",
	WORKSHOP_AM_MICRO:  "WorkshopMicro",
	WORKSHOP_PM_MICRO:  "WorkshopMicro",
	WORKSHOP_DEBUG:     "WorkshopDebug",
	WORKSHOP_AM_DEBUG:  "WorkshopDebug",
	WORKSHOP_PM_DEBUG:  "WorkshopDebug",
	WORKSHOP_ML:        "WorkshopML",
	WORKSHOP_AM_ML:     "WorkshopML",
	WORKSHOP_PM_ML:     "WorkshopML",
	WORKSHOP_CONCUR:    "WorkshopConcur",
	WORKSHOP_AM_CONCUR: "WorkshopConcur",
	WORKSHOP_PM_CONCUR: "WorkshopConcur",
	WORKSHOP_SLS:       "WorkshopServerless",
	WORKSHOP_AM_SLS:    "WorkshopServerless",
	WORKSHOP_PM_SLS:    "WorkshopServerless",
	WORKSHOP_AM_K8S:    "WorkshopKubernetes",
	WORKSHOP_PM_K8S:    "WorkshopKubernetes",
	WORKSHOP_AM_PACMAN: "WorkshopPacman",
	WORKSHOP_PM_PACMAN: "WorkshopPacman",
	WORKSHOP_DGRAPH:    "WorkshopDgraph",
	GEN_EARLY:          "GC20 Ticket",
	GEN_SELF:           "GC20 Ticket",
	GEN_CORP:           "GC20 Ticket",
	COMMUNITY:          "GC20 Ticket",
	TWENTY_PERPET:      "GC20 Ticket",
	ONETIME_TWENTY:     "GC20 Ticket",
	FRIENDS:            "GC20 Ticket",
	G2A_GROUP:          "GC20 Ticket",
}

const introductions = "765700981639217153"

type Config struct {
	Token       string `json:"token"`
	LogChannel  string `json:"logChannel"`
	GopherGuild string `json:"gopherGuild"`
}

func NewConfig(token, logChannel, gopherGuild string) *Config {
	return &Config{
		Token:       token,
		LogChannel:  logChannel,
		GopherGuild: gopherGuild,
	}
}

func (c *Config) GetToken() string {
	if strings.HasPrefix(c.Token, "$") {
		return os.Getenv(strings.TrimPrefix(c.Token, "$"))
	}
	return c.Token
}

func (c *Config) GetGopherGuild() string {
	if strings.HasPrefix(c.GopherGuild, "$") {
		return os.Getenv(strings.TrimPrefix(c.GopherGuild, "$"))
	}
	return c.GopherGuild
}

func (c *Config) GetLogChannel() string {
	if strings.HasPrefix(c.LogChannel, "$") {
		return os.Getenv(strings.TrimPrefix(c.LogChannel, "$"))
	}
	return c.LogChannel
}

type Discord struct {
	config     *Config
	session    *discordgo.Session
	GuildRoles []*discordgo.Role
	role       *discordgo.Role
	commands   []Command
}

func NewDiscord(c *Config) *Discord {

	d := &Discord{}

	a1 := Arg{
		Name:              "Email",
		Required:          true,
		ValidationMessage: "email is required",
	}

	a2 := Arg{
		Name:              "TicketNumber",
		Required:          true,
		ValidationMessage: "ticket number is required",
	}
	command := Command{
		Trigger:     "validate",
		Args:        []Arg{a1, a2},
		Bot:         d,
		Description: "Automatically assign appropriate discord roles based on the ticket you have purchased.\n `validate myemail@some.com 152345`",
		Action: func(c Command, s *discordgo.Session, user *discordgo.User, channel string) error {
			d.logAction("user", c.Args[0].Value, "sending search notice")
			s.ChannelMessageSend(channel, user.Mention()+" searching for tickets assigned to "+c.Args[0].Value)
			// process stuff here'

			d.logAction("user", c.Args[0].Value, "email", c.Args[1].Value, "searching for tickets")
			tix, err := database.ValidateByEmailTicket(c.Args[0].Value, c.Args[1].Value)
			if err != nil {
				d.logAction("user", c.Args[0].Value, "email", c.Args[1].Value, "database error", err.Error())
				s.ChannelMessageSend(channel, user.Mention()+" Sorry, we broke things... {validatebyemail}"+err.Error())
				return err
			}
			if len(tix) > 0 {

				d.logAction("user", c.Args[0].Value, "email", c.Args[1].Value, "found tix", strconv.Itoa(len(tix)))
				pretty := fmt.Sprintf(" Great news, I found %d ticket.", len(tix))
				s.ChannelMessageSend(channel, user.Mention()+pretty)
				alltix, err := database.TicketsByEmail(c.Args[0].Value)

				if err != nil {

					d.logAction("user", c.Args[0].Value, "email", c.Args[1].Value, "tix by email error", err.Error())
					s.ChannelMessageSend(channel, user.Mention()+" Sorry, we broke things... {al tickets}"+err.Error())
					return err
				}
				var roles []string
				for _, ticket := range alltix {
					if strings.Contains(ticket.TicketType, "orkshop") {
						roles = append(roles, "GC20 Workshop")
					}
					role := TicketRole[ticket.TicketType]
					roles = append(roles, role)

					if ticket.PromoCode != "" {
						sponsorRoles := database.GetPromoRoles(ticket.PromoCode)
						if len(sponsorRoles) > 0 {
							roles = append(roles, sponsorRoles...)
						}
					}
				}
				pretty = fmt.Sprintf(" There are %d total ticket(s) associated. Adding appropriate roles to your account.  Beep Boop.", len(alltix))
				s.ChannelMessageSend(channel, user.Mention()+pretty)
				v := &database.ValidatedUser{
					UserID:       user.ID,
					EmailAddress: c.Args[0].Value,
				}
				v.MakeRoleList(roles)

				//now apply the roles
				for _, r := range roles {
					var roleFound bool
					for _, gr := range d.GuildRoles {
						if (gr.Name == r) || (gr.ID == r) {
							roleFound = true
							fmt.Println(gr.Name)
							s.ChannelMessageSend(channel, user.Mention()+gr.Name)
							s.GuildMemberRoleAdd(d.config.GetGopherGuild(), v.UserID, gr.ID)
						}
					}
					if !roleFound {
						fmt.Println("Couldn't find a guild role for ", r)
					}
				}
				//s.GuildMemberRoleAdd()

				err = database.SaveValidatedUser(v)
				if err != nil {

					d.logAction("user", c.Args[0].Value, "email", c.Args[1].Value, "save validated user error", err.Error())
					return err
				}
				s.ChannelMessageSend(channel, user.Mention()+"You are a certified Gopher. Thanks!")

				d.logAction("user", c.Args[0].Value, "email", c.Args[1].Value, "validated successfully")
				return nil
			}

			d.logAction("user", c.Args[0].Value, "email", c.Args[1].Value, "no tickets found")
			s.ChannelMessageSend(channel, user.Mention()+" No tickets found... try a different email address?")
			return nil
		},
	}

	sponsors := Command{
		Trigger:     "sponsors",
		Args:        []Arg{},
		Bot:         d,
		Description: "List GopherCon sponsors",
		Action: func(c Command, s *discordgo.Session, user *discordgo.User, channel string) error {
			s.ChannelMessageSend(channel, user.Mention()+" finding sponsors...")
			sp, err := database.ListSponsors()
			if err != nil {
				s.ChannelMessageSend(channel, user.Mention()+" error listing sponsors: "+err.Error())
			}
			for _, sponsor := range sp {
				ch, err := s.Channel(sponsor.Channel)
				if err != nil {
					fmt.Println(err)

					s.ChannelMessageSend(channel, sponsor.Name)
					continue
				}
				mention := ch.Mention()
				s.ChannelMessageSend(channel, sponsor.Level+" "+mention+" "+sponsor.Name)
			}
			return nil
		},
	}
	d.config = c
	d.commands = []Command{command, sponsors}
	return d
}

func (d *Discord) PresenceUpdate(s *discordgo.Session, m *discordgo.PresenceUpdate) {
	// ignore the stupid role bot
	if m.User.ID == "480240313525600267" {
		return
	}
	go d.ReconcileWorkshops(s, m)
	greeted, err := database.Greeted(m.User.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	if greeted {
		fmt.Println("already greeted", m.User.ID)
		return
	}
	// don't trigger for the double status notifications
	// as people connect, then come online.
	if m.Status == discordgo.StatusOnline {
		fmt.Println("Sending a greeting to ", m.User.ID)
		d.Greet(m.User)
	}

}

// ReconcileWorkshops adds the GC20 Workshop role to people who should have had it
func (d *Discord) ReconcileWorkshops(s *discordgo.Session, m *discordgo.PresenceUpdate) {

	d.logAction("reconciling roles for ", m.User.ID)
	member, err := d.session.GuildMember(d.config.GetGopherGuild(), m.User.ID)
	if err != nil {
		d.logAction("error getting guild member", err.Error())
		return
	}

	var hasWorkshop bool
	var hasGC20Workshop bool
	for _, r := range member.Roles {
		for _, gr := range d.GuildRoles {
			if gr.ID == r {
				if gr.Name[0:2] == "Wo" {
					hasWorkshop = true
					d.logAction("workshop found for", m.User.ID)
				}
				if gr.ID == "763773702415450173" {
					d.logAction("GC20 Workshop role found for", m.User.ID)
					hasGC20Workshop = true
					return
				}

			}
		}
	}
	if hasWorkshop && !hasGC20Workshop {
		err = s.GuildMemberRoleAdd(d.config.GetGopherGuild(), m.User.ID, "763773702415450173")
		if err != nil {
			d.logAction("error adding gc20 guild role to user", m.User.ID, err.Error())
			return
		}
		d.logAction("GC20 Workshop added to", m.User.ID)
	}

	return

}

func (d *Discord) Greet(u *discordgo.User) {
	ch, err := d.session.UserChannelCreate(u.ID)
	if err != nil {
		d.logAction("error creating dm channel", err.Error())
		fmt.Println("error creating dm channel", err)
		d.session.ChannelMessageSend(introductions, u.Mention()+"please enable Direct Messages to validate your tickets.")
		return
	}
	_, err = d.session.ChannelMessageSend(ch.ID, "Welcome! I'm the GopherCon bot.  Here are some things I can do:")
	if err != nil {
		d.logAction("error sending greeting to user", err.Error())
		return
	}
	for _, cmd := range d.commands {
		d.session.ChannelMessageSend(ch.ID, cmd.Help())
	}

	d.session.ChannelMessageSend(ch.ID, "If this is your first visit, please use the `validate` command to automatically join the correct groups for your tickets.")
	err = database.NewGreeting(u.ID, ch.ID)
	if err != nil {
		fmt.Println("error saving greeting", err)
	}
}

// MessageCreate is the function called on new message events
func (d *Discord) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	var dm bool
	if m.GuildID == "" {
		dm = true
	}

	c, e := m.ContentWithMoreMentionsReplaced(d.session)
	if e != nil {
		fmt.Println(e)
	}

	if dm {
		d.logAction("processing dm command ", c)
		d.ProcessCommand(dm, m.Author, m.ChannelID, c)
		return
	}
	for _, rr := range m.MentionRoles {
		if rr == d.role.ID {
			// get the raw content
			// now try to process this command
			d.logAction("processing @mention command: ", c)
			d.ProcessCommand(dm, m.Author, m.ChannelID, c)
		}
	}
}

func (d *Discord) ProcessCommand(dm bool, user *discordgo.User, channel, message string) {

	parts := strings.Fields(message)
	var trigger string
	var rest string
	if dm {
		trigger = parts[0]
		if len(parts) >= 1 {

			rest = strings.Join(parts[1:], " ")
		}
	} else {
		if len(parts) >= 2 {
			trigger = parts[1]
			rest = strings.Join(parts[2:], " ")
		}
	}

	var found bool
	for _, c := range d.commands {
		if c.Trigger == trigger {
			found = true
			// this is the command, parse it
			err := c.Parse(rest)
			if err != nil {
				// something didn't validate. Print error and return
				d.session.ChannelMessageSend(channel, user.Mention()+" "+err.Error())
				return
			}
			// then execute it
			c.Action(c, d.session, user, channel)
		}
	}
	if !found {
		d.session.ChannelMessageSend(channel, user.Mention()+" unknown command :( ")
	}
}

// ProcessCommand tries to make sense of the user input
/*
func (d *Discord) ProcessCommand(s *discordgo.Session, user *discordgo.User, channel, message string) {
	parts := strings.Fields(message)
	if len(parts) > 1 {
		command := parts[1]
		if command == "validate" {
			if len(parts) > 2 {
				email := parts[2]
				ticketnum := parts[3]
				s.ChannelMessageSend(channel, user.Mention()+" looking for tickets assigned to "+email)
				// process stuff here
				tix, err := database.ValidateByEmailTicket(email, ticketnum)
				if err != nil {
					s.ChannelMessageSend(channel, user.Mention()+" Sorry, we broke things... {validatebyemail}"+err.Error())
				}
				if len(tix) > 0 {
					pretty := fmt.Sprintf(" Great news, I found %d tickets.", len(tix))
					s.ChannelMessageSend(channel, user.Mention()+pretty)
					alltix, err := database.TicketsByEmail(email)

					if err != nil {
						s.ChannelMessageSend(channel, user.Mention()+" Sorry, we broke things... {al tickets}"+err.Error())
					}
					var roles []string
					for _, ticket := range alltix {
						role := TicketRole[ticket.TicketType]
						roles = append(roles, role)
					}
					v := &database.ValidatedUser{
						UserID:       user.ID,
						EmailAddress: email,
					}
					v.MakeRoleList(roles)
					err = database.SaveValidatedUser(v)

					if err != nil {
						s.ChannelMessageSend(channel, user.Mention()+" Sorry, we broke things...{savevalidated}"+err.Error())
					}

					//now apply the roles
					//s.GuildMemberRoleAdd()
					return
				}
				s.ChannelMessageSend(channel, user.Mention()+" None found")
			} else {
				s.ChannelMessageSend(channel, user.Mention()+" I need your email address and your ticket number. Try `validate you@yourdomain.com` 12345670.")
				return
			}
			return
		}

	}
	s.ChannelMessageSend(channel, user.Mention()+" I don't know that trick yet.  Try `tickets`.")

}
*/

func (d *Discord) logAction(action ...string) {
	message := strings.Join(action, " : ")
	fmt.Println(message)
	d.session.ChannelMessageSend(d.config.GetLogChannel(), message)

}

func (d *Discord) RunSchedule(quit chan bool) {
	d.logAction("initializing scheduler")
	sm, err := database.GetScheduledMessages()
	if err != nil {
		panic(err) // panic??
	}
	c := cron.New()
	for _, m := range sm {
		d.logAction("adding", m.Cron, m.Message)
		c.AddFunc(m.Cron, func() {
			d.logAction("executing cron trigger", m.Cron, m.Message)
			d.session.ChannelMessageSend(m.TargetChannel, m.Message)
		})
	}
	d.logAction("Starting scheduler")
	c.Start()

	select {
	case _ = <-quit:
		fmt.Println("received quit - quitting scheduler")
		c.Stop()
		return
	}
	fmt.Println("outside of the select, shouldn't be here in scheduler")
}

// Run starts the Discord bot
func (d *Discord) Run(quit chan bool) error {

	var err error
	// Create a new Discord session using the provided bot token.
	d.session, err = discordgo.New("Bot " + d.config.GetToken())
	if err != nil {
		return fmt.Errorf("Error connecting to discord %v", err)
	}
	d.logAction("In ur serverz, doin bot thingz.")
	// Register the messageCreate func as a callback for MessageCreate events.
	d.session.AddHandler(d.MessageCreate)
	d.session.AddHandler(d.PresenceUpdate)

	// In this example, we only care about receiving message events.
	d.session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	// Open a websocket connection to Discord and begin listening.
	err = d.session.Open()
	if err != nil {
		return err
	}

	// Get all the roles in the guild
	d.GuildRoles, err = d.session.GuildRoles(d.config.GetGopherGuild())
	if err != nil {
		d.logAction("Error getting roles:", err.Error())
	}

	// find our role -- Probably a better way?
	for _, gr := range d.GuildRoles {
		fmt.Println(gr.Name)
		if gr.Name == "gcbot" {
			d.role = gr
		}
	}
	q := make(chan bool, 1)

	go d.RunSchedule(q)

	go func() {
		r := mux.NewRouter()
		r.HandleFunc("/order", hookHandler).Methods("POST")
		http.Handle("/", r)

		log.Printf("Listening on %s\n", "3000")
		if err := http.ListenAndServe(":3000", nil); err != nil {
			log.Fatal(err)
		}
	}()

	// Wait here until CTRL-C or other term signal is received.
	d.logAction("Bot is now running.  Press CTRL-C to exit.")
	// wait on the quit channel
	<-quit
	// kill our scheduler
	q <- true

	// Cleanly close down the Discord session.
	d.session.Close()

	return nil
}
