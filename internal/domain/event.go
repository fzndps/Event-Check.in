package domain

import "time"

// Entity event
type Event struct {
	ID               string    `json:"id"` // UUID format
	OrganizerID      int64     `json:"organizer_id"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"` // URL-friendly name
	Date             time.Time `json:"date"`
	Venue            string    `json:"venue"`
	ParticipantCount int       `json:"participant_count"`
	TotalPrice       int       `json:"total_price"`
	PaymentStatus    string    `json:"payment_status"`
	PaymentProofURL  string    `json:"payment_proof_url"`
	ScannerPIN       string    `json:"scanner_pin"`
	CreatedAt        time.Time `json:"created_at"`

	// Relationships (untuk response, tidak disimpan di DB)
	Organizer    *Organizer     `json:"organizer,omitempty"`
	Participants []*Participant `json:"participants,omitempty"`
}

const (
	PaymentStatusPending  = "pending"
	PaymentStatusVerified = "verified"
	PaymentStatusActive   = "active"
)

// DTO create event
type CreateEventRequest struct {
	Name             string    `json:"name" binding:"required,min=3,max=255"`
	Date             time.Time `json:"date" binding:"required"`
	Venue            string    `json:"venue" binding:"required,min=5,max=500"`
	ParticipantCount int       `json:"participant_count" binding:"required,min=1"`
}

// DTO update event
type UpdateEventRequest struct {
	Name  string    `json:"name" binding:"omitempty,required,min=3,max=255"`
	Date  time.Time `json:"date" binding:"required"`
	Venue string    `json:"venue" binding:"omitempty,required,min=5,max=500"`
}

// Response list event
type EventListResponse struct {
	Events    []*Event `json:"events"`
	Total     int      `json:"total"`
	Page      int      `json:"page"`
	Limit     int      `json:"Limit"`
	TotalPage int      `json:"total_page"`
}

// Response detail event dengan partisipan
type EventDetailResponse struct {
	Event                 *Event         `json:"event"`
	Participants          []*Participant `json:"participants"`
	ParticipantRegistered int            `json:"participant_registered"`
	ParticipantCheckedIn  int            `json:"participant_checked_in"`
}

// Melakukan validasi untuk event
func (r *CreateEventRequest) Validate() error {
	// Event date tidak boleh di masalalu
	if r.Date.Before(time.Now()) {
		return ErrInvalidEventDate
	}

	return nil
}

// kalkulasi harga per partisipan
func CalculatePrice(participantCount int) int {
	var pricePerParticipant int

	switch {
	case participantCount <= 50:
		pricePerParticipant = 5000
	case participantCount <= 100:
		pricePerParticipant = 4500
	case participantCount <= 500:
		pricePerParticipant = 4000
	default:
		pricePerParticipant = 3500
	}

	return participantCount * pricePerParticipant
}
