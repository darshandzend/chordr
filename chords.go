package main

type chord struct {
	id        int
	name      string
	soundfile string
}

var chords = []chord{
	chord{1, "C", "sounds/C.mp3"},
	chord{2, "D", "sounds/D.mp3"},
	chord{3, "Dm", "sounds/Dm.mp3"},
	chord{4, "E", "sounds/E.mp3"}}

func Get(id int) chord {
	return chords[id]
}
