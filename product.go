package main

type Product struct {
	URL    string
	Name   string
	Number string
	Track  bool
}

const (
	// plate urls
	steelPlateURL    = "https://www.roguefitness.com/rogue-calibrated-lb-steel-plates"
	machinedPlateURL = "https://www.roguefitness.com/rogue-machined-olympic-plates"
	olympicPlateURL  = "https://www.roguefitness.com/rogue-olympic-plates"
	echoPlateURL     = "https://www.roguefitness.com/rogue-echo-bumper-plates-with-white-text"
)

// plate products have multiple products per page so we will want to only perform a single request per url
var plateProducts = []Product{
	// Steel plates
	{
		URL:    steelPlateURL,
		Name:   "2.5LB Calibrated LB Steel Plate - Pair",
		Number: "38875",
		Track:  false,
	},
	{
		URL:    steelPlateURL,
		Name:   "5LB Calibrated LB Steel Plate - Pair",
		Number: "38877",
		Track:  true,
	},
	{
		URL:    steelPlateURL,
		Name:   "10LB Calibrated LB Steel Plate - Pair",
		Number: "38879",
		Track:  true,
	},
	{
		URL:    steelPlateURL,
		Name:   "25LB Calibrated LB Steel Plate - Pair",
		Number: "38881",
		Track:  true,
	},
	{
		URL:    steelPlateURL,
		Name:   "35LB Calibrated LB Steel Plate - Pair",
		Number: "38883",
		Track:  true,
	},
	{
		URL:    steelPlateURL,
		Name:   "45LB Calibrated LB Steel Plate - Pair",
		Number: "38885",
		Track:  true,
	},
	// Machined Olympic Plates
	{
		URL:    machinedPlateURL,
		Name:   "2.5LB Machined Olympic - Pair",
		Number: "47205",
		Track:  false,
	},
	{
		URL:    machinedPlateURL,
		Name:   "5LB Machined Olympic - Pair",
		Number: "47211",
		Track:  true,
	},
	{
		URL:    machinedPlateURL,
		Name:   "10LB Machined Olympic - Pair",
		Number: "47213",
		Track:  true,
	},
	{
		URL:    machinedPlateURL,
		Name:   "25LB Machined Olympic - Pair",
		Number: "47227",
		Track:  true,
	},
	{
		URL:    machinedPlateURL,
		Name:   "35LB Machined Olympic - Pair",
		Number: "47229",
		Track:  true,
	},
	{
		URL:    machinedPlateURL,
		Name:   "45LB Machined Olympic - Pair",
		Number: "47219",
		Track:  true,
	},
	// Machined Olympic Plates
	{
		URL:    olympicPlateURL,
		Name:   "2.5LB Olympic Plate - Pair",
		Number: "7185",
		Track:  false,
	},
	{
		URL:    olympicPlateURL,
		Name:   "5LB Olympic Plate - Pair",
		Number: "7183",
		Track:  true,
	},
	{
		URL:    olympicPlateURL,
		Name:   "10LB Olympic Plate - Pair",
		Number: "7181",
		Track:  true,
	},
	{
		URL:    olympicPlateURL,
		Name:   "25LB Olympic Plate - Pair",
		Number: "7179",
		Track:  true,
	},
	{
		URL:    olympicPlateURL,
		Name:   "35LB Olympic Plate - Pair",
		Number: "7177",
		Track:  true,
	},
	{
		URL:    olympicPlateURL,
		Name:   "45LB Olympic Plate - Pair",
		Number: "7175",
		Track:  true,
	},
	// Echo Bumper Plates
	{
		URL:    echoPlateURL,
		Name:   "10LB Echo Bumper Plate - Pair",
		Number: "65856",
		Track:  true,
	},
	{
		URL:    echoPlateURL,
		Name:   "15LB Echo Bumper Plate - Pair",
		Number: "65858",
		Track:  true,
	},
	{
		URL:    echoPlateURL,
		Name:   "25LB Echo Bumper Plate - Pair",
		Number: "65860",
		Track:  true,
	},
	{
		URL:    echoPlateURL,
		Name:   "35LB Echo Bumper Plate - Pair",
		Number: "65862",
		Track:  true,
	},
	{
		URL:    echoPlateURL,
		Name:   "45LB Echo Bumper Plate - Pair",
		Number: "65864",
		Track:  true,
	},
}

var barProducts = []Product{
	{
		URL:   "https://www.roguefitness.com/rogue-t-2-5kg-technique-bar",
		Name:  "Rogue T-2.5KG Technique Bar",
		Track: true,
	},
	{
		URL:   "https://www.roguefitness.com/rogue-bella-bar-2-0-black-zinc",
		Name:  "The Bella Bar 2.0 - Black Zinc",
		Track: true,
	},
	{
		URL:   "https://www.roguefitness.com/rogue-45lb-ohio-power-bar-black-zinc",
		Name:  "Rogue 45LB Ohio Power Bar - Black Zinc",
		Track: true,
	},
	{
		URL:   "https://www.roguefitness.com/rogue-45lb-ohio-power-bar-bare-steel",
		Name:  "Rogue 45LB Ohio Power Bar - Bare Steel",
		Track: true,
	},
	{
		URL:   "https://www.roguefitness.com/rogue-45lb-ohio-powerlift-bar-cerakote",
		Name:  "Rogue Ohio Power Bar - Cerakote",
		Track: true,
	},
	{
		URL:   "https://www.roguefitness.com/rogue-45lb-ohio-power-bar-stainless",
		Name:  "Rogue 45LB Ohio Power Bar - Stainless Steel",
		Track: true,
	},
	{
		URL:   "https://www.roguefitness.com/rogue-ohio-power-bar-e-coat",
		Name:  "Rogue 45LB Ohio Power Bar - E-Coat",
		Track: true,
	},
	{
		URL:   "https://www.roguefitness.com/rogue-echo-bar",
		Name:  "Rogue Echo Bar 2.0",
		Track: true,
	},
	{
		URL:   "https://www.roguefitness.com/rogue-ohio-bar-black-oxide",
		Name:  "The Ohio Bar - Black Oxide",
		Track: true,
	},
	{
		URL:   "https://www.roguefitness.com/the-ohio-bar-cerakote",
		Name:  "The Ohio Bar - Cerakote",
		Track: true,
	},
	{
		URL:   "https://www.roguefitness.com/the-ohio-bar-black-zinc",
		Name:  "The Ohio Bar - Black Zinc",
		Track: true,
	},
	{
		URL:   "https://www.roguefitness.com/rogue-ohio-bar-black-oxide",
		Name:  "The Ohio Bar - Black Oxide",
		Track: true,
	},
	{
		URL:   "https://www.roguefitness.com/rogue-ohio-power-bar-e-coat",
		Name:  "Rogue 45LB Ohio Power Bar - E-Coat",
		Track: true,
	},
	{
		URL:   "https://www.roguefitness.com/the-rogue-bar-2-0",
		Name:  "The Rogue Bar 2.0",
		Track: true,
	},
	{
		URL:   "https://www.roguefitness.com/rogue-25mm-womens-b-r-bar",
		Name:  "25MM Women’s B&R Bar 2.0",
		Track: false,
	},
	{
		URL:   "https://www.roguefitness.com/stainless-steel-ohio-bar",
		Name:  "The Ohio Bar - Stainless Steel",
		Track: false,
	},
}
