package main

// This file contains configuration for all the entities in the game
// at some point it would be nice to have this imported through
// some form of configuration file so the game can be modded.

type EntityType interface {
	GetName() string
	GetDescription() string
	InitiateBrain() *Brain
	InitiateFighter() *Fighter
}

// @todo finish and implement
