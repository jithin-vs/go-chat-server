package chatws

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	ID           primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Participants map[string]*Participant  `json:"participants"`
	CreatedAt    primitive.Timestamp `json:"createdAt" bson:"created_at"`
	UpdatedAt    primitive.Timestamp `json:"updatedAt" bson:"updated_at"`
	Type         string    `json:"type" bson:"type"`
}

type Hub struct {
	
	Rooms map[string]*Room
	Register  chan *Participant
	Broadcast chan *Message
}

func NewHub() *Hub {
	return &Hub{
        Rooms: make(map[string]*Room),
		Register: make(chan *Participant),
		Broadcast: make(chan *Message,5),
    }
}

func (h *Hub) Run(){
    for{
		select{
		    case cl:= <- h.Register:
			   if _,exists := h.Rooms[cl.RoomID]; exists {
				  r := h.Rooms[cl.RoomID]

				  if _,exists :=r.Participants[cl.ID]; !exists{
					r.Participants[cl.ID]=cl
				  }
			   }
		}
	}
}