package util

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomHeroName() string {
	adjectives := []string{
		"Amazing", "Spectacular", "Incredible", "Mighty", "Invincible",
		"Fantastic", "Brave", "Fearless", "Savage", "Dynamic", "Heroic",
		"Swift", "Stealthy", "Thunderous", "Radiant", "Blazing", "Vigilant",
	}

	creatures := []string{
		"Spider", "Panther", "Falcon", "Wolf", "Lion", "Tiger", "Eagle",
		"Phoenix", "Dragon", "Bear", "Cheetah", "Cobra", "Shark", "Rhino",
		"Hawk", "Jaguar", "Serpent", "Stallion", "Crane", "Fox", "Leopard",
	}

	rand.Seed(time.Now().UnixNano())

	adjective := adjectives[rand.Intn(len(adjectives))]
	creature := creatures[rand.Intn(len(creatures))]

	// Get last 4 digits of time in milliseconds
	timestamp := time.Now().UnixMilli() % 10000

	return fmt.Sprintf("%s-%s-%03d", adjective, creature, timestamp)
}
