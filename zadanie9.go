package main 

import (
	"fmt"
)

const (
	Single = "single"
	Double = "double"
	Suit = "suit"
	
)

const (
	Free = "free"
	Booked = "booked"
	Maintenance = "maintenance"
)

type HotelRoom struct {
	Type string 
	Status string
	Cost float64 
}

func Book_room (hotel_rooms map[string] * HotelRoom, number_room string) {
	if room, ok := hotel_rooms[number_room];ok && room.Status == Free {
		room.Status = Booked
		fmt.Println("Комната ", number_room, " забронирована")
	} else {
		fmt.Println("Комната ", number_room, " недоступна")
	}
}

func main () {
	var hotel_rooms = make(map[string]*HotelRoom)
	hotel_rooms["123"] = &HotelRoom{Type: Single, Status: Free, Cost: 500}
	hotel_rooms["144"] = &HotelRoom{Type: Double, Status: Maintenance, Cost: 501}
	
	Book_room(hotel_rooms, "123")
}