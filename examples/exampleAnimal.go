package main

type animal struct {
	TypeAnimal string `json:"typeAnimal,omitempty"`
	Name       string `json:"name,omitempty"`
	Age        int    `json:"age,omitempty"`
	Color      string `json:"color,omitempty"`
	NbrLegs    int    `json:"nbrLegs,omitempty"`
	NbrEyes    int    `json:"nbrEyes,omitempty"`
}

var Animal animal
