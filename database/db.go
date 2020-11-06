package database

import (
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	// github.com/mattn/go-sqlite3
	dsn := os.Getenv("DSN")

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//	fmt.Println(GetPromoRoles("BrnzSNBasdfI90320"))
}

func SaveTickets(tix []*GTicket) error {
	result := db.Create(tix)
	return result.Error
}

func Migrate() error {
	return db.AutoMigrate(&Sponsor{}, &GTicket{}, &ValidatedUser{}, &Sponsor{}, &TicketCode{}, &ScheduledMessage{}, &Greeting{})
}

type ValidatedUser struct {
	gorm.Model
	UserID       string
	EmailAddress string
	Roles        string
}

func (v *ValidatedUser) RoleList() []string {
	return strings.Split(v.Roles, ",")
}

func (v *ValidatedUser) MakeRoleList(roles []string) {
	v.Roles = strings.Join(roles, ",")
}

func SaveValidatedUser(v *ValidatedUser) error {
	result := db.Create(&v)
	return result.Error
}
func ValidatedUsers() ([]ValidatedUser, error) {
	var vu []ValidatedUser
	db.Find(&vu)
	return vu, db.Error
}

type GTicket struct {
	gorm.Model
	Ticket
}

type Ticket struct {
	FirstName              string  `xlsx:"0"`
	LastName               string  `xlsx:"1"`
	Company                string  `xlsx:"2"`
	JobTitle               string  `xlsx:"3"`
	OrderDate              string  `xlsx:"4"`
	TicketType             string  `xlsx:"5"`
	TicketPrice            float64 `xlsx:"6"`
	TicketPaid             float64 `xlsx:"7"`
	NetIncome              float64 `xlsx:"8"`
	TicketNumber           string  `xlsx:"9"`
	PromoCode              string  `xlsx:"10"`
	PromoName              string  `xlsx:"11"`
	RefundAmount           string  `xlsx:"12"`
	EventCheckedIn         string  `xlsx:"13"`
	EventCheckedInDate     string  `xlsx:"14"`
	SessionCheckIn         string  `xlsx:"15"`
	SessionCheckOut        string  `xlsx:"16"`
	Tax                    string  `xlsx:"17"`
	ProcessingFee          string  `xlsx:"18"`
	BizzaboFee             string  `xlsx:"19"`
	RegistrationStatus     string  `xlsx:"20"`
	OrderType              string  `xlsx:"21"`
	OrderPaymentStatus     string  `xlsx:"22"`
	PaymentMethod          string  `xlsx:"23"`
	OrderNumber            string  `xlsx:"24"`
	OrderPlacedByName      string  `xlsx:"25"`
	OrderPlacedByCompany   string  `xlsx:"26"`
	OrderPlacedByLastName  string  `xlsx:"27"`
	OrderPlacedByEmail     string  `xlsx:"28"`
	BilledContact          string  `xlsx:"29"`
	BilledCountry          string  `xlsx:"30"`
	BillingAddress         string  `xlsx:"31"`
	BillingCity            string  `xlsx:"32"`
	BillingState           string  `xlsx:"33"`
	BillingZipCode         string  `xlsx:"34"`
	AdditionalNotes        string  `xlsx:"35"`
	BillingTaxID           string  `xlsx:"36"`
	OrderNotes             string  `xlsx:"37"`
	Currency               string  `xlsx:"38"`
	ReferralCode           string  `xlsx:"39"`
	EventName              string  `xlsx:"40"`
	TicketDiscountAmount   string  `xlsx:"41"`
	EventVenueName         string  `xlsx:"42"`
	EventVenueCity         string  `xlsx:"43"`
	PayByInvoice           string  `xlsx:"44"`
	EventVenueCountry      string  `xlsx:"45"`
	EventVenueState        string  `xlsx:"46"`
	PaymentStatus          string  `xlsx:"47"`
	RegistrationDate       string  `xlsx:"-"`
	PaymentDate            string  `xlsx:"-"`
	LedgerCode             string  `xlsx:"50"`
	DateRefunded           string  `xlsx:"51"`
	InvoiceNumber          string  `xlsx:"52"`
	UTMSource              string  `xlsx:"53"`
	UTMMedium              string  `xlsx:"54"`
	UTMCampaign            string  `xlsx:"55"`
	UTMTerm                string  `xlsx:"56"`
	UTMContent             string  `xlsx:"57"`
	VirtualEventAttendance string  `xlsx:"58"`
	RefundPolicy           string  `xlsx:"59"`
	Country                string  `xlsx:"60"`
	Prefix                 string  `xlsx:"61"`
	TwitterHandle          string  `xlsx:"62"`
	SponsorContactOptIn    string  `xlsx:"63"`
	Industry               string  `xlsx:"64"`
	AttendeeAssociation    string  `xlsx:"65"`
	FirstTimeAttendees     string  `xlsx:"66"`
	CompanyWebsite         string  `xlsx:"67"`
	EmailAddress           string  `xlsx:"68"`
}

func ValidateByEmailTicket(email, ticketnum string) ([]GTicket, error) {
	var tix []GTicket
	db.Where("email_address = ? and ticket_number = ?", email, ticketnum).Find(&tix)
	return tix, db.Error
}

func TicketsByEmail(email string) ([]GTicket, error) {
	var tix []GTicket
	db.Where("email_address = ?", email).Find(&tix)
	return tix, db.Error
}
