package gofakeit

import (
	"math/rand"
	"strconv"
)

// BeerName will return a random beer name
func BeerName() string {
	return beerName(globalFaker.Rand)
}

// BeerName will return a random beer name
func (f *Faker) BeerName() string {
	return beerName(f.Rand)
}

func beerName(r *rand.Rand) string {
	return getRandValue(r, []string{"beer", "name"})
}

// BeerStyle will return a random beer style
func BeerStyle() string {
	return beerStyle(globalFaker.Rand)
}

// BeerStyle will return a random beer style
func (f *Faker) BeerStyle() string {
	return beerStyle(f.Rand)
}

func beerStyle(r *rand.Rand) string {
	return getRandValue(r, []string{"beer", "style"})
}

// BeerHop will return a random beer hop
func BeerHop() string {
	return beerHop(globalFaker.Rand)
}

// BeerHop will return a random beer hop
func (f *Faker) BeerHop() string {
	return beerHop(f.Rand)
}

func beerHop(r *rand.Rand) string {
	return getRandValue(r, []string{"beer", "hop"})
}

// BeerYeast will return a random beer yeast
func BeerYeast() string {
	return beerYeast(globalFaker.Rand)
}

// BeerYeast will return a random beer yeast
func (f *Faker) BeerYeast() string {
	return beerYeast(f.Rand)
}

func beerYeast(r *rand.Rand) string {
	return getRandValue(r, []string{"beer", "yeast"})
}

// BeerMalt will return a random beer malt
func BeerMalt() string {
	return beerMalt(globalFaker.Rand)
}

// BeerMalt will return a random beer malt
func (f *Faker) BeerMalt() string {
	return beerMalt(f.Rand)
}

func beerMalt(r *rand.Rand) string {
	return getRandValue(r, []string{"beer", "malt"})
}

// BeerAlcohol will return a random beer alcohol level between 2.0 and 10.0
func BeerAlcohol() string {
	return beerAlcohol(globalFaker.Rand)
}

// BeerAlcohol will return a random beer alcohol level between 2.0 and 10.0
func (f *Faker) BeerAlcohol() string {
	return beerAlcohol(f.Rand)
}

func beerAlcohol(r *rand.Rand) string {
	return strconv.FormatFloat(float64Range(r, 2.0, 10.0), 'f', 1, 64) + "%"
}

// BeerIbu will return a random beer ibu value between 10 and 100
func BeerIbu() string {
	return beerIbu(globalFaker.Rand)
}

// BeerIbu will return a random beer ibu value between 10 and 100
func (f *Faker) BeerIbu() string {
	return beerIbu(f.Rand)
}

func beerIbu(r *rand.Rand) string {
	return strconv.Itoa(randIntRange(r, 10, 100)) + " IBU"
}

// BeerBlg will return a random beer blg between 5.0 and 20.0
func BeerBlg() string {
	return beerBlg(globalFaker.Rand)
}

// BeerBlg will return a random beer blg between 5.0 and 20.0
func (f *Faker) BeerBlg() string {
	return beerBlg(f.Rand)
}

func beerBlg(r *rand.Rand) string {
	return strconv.FormatFloat(float64Range(r, 5.0, 20.0), 'f', 1, 64) + "°Blg"
}

func addBeerLookup() {
	AddFuncLookup("beername", Info{
		Display:     "Beer Name",
		Category:    "beer",
		Description: "Name for a beer, inspired by real beer names",
		Example:     "Duvel",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			return beerName(r), nil
		},
	})

	AddFuncLookup("beerstyle", Info{
		Display:     "Beer Style",
		Category:    "beer",
		Description: "Style of beer, covering beer types from around the world",
		Example:     "European Amber Lager",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			return beerStyle(r), nil
		},
	})

	AddFuncLookup("beerhop", Info{
		Display:     "Beer Hop",
		Category:    "beer",
		Description: "Beer hop variety, covering different hop types used in brewing",
		Example:     "Glacier",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			return beerHop(r), nil
		},
	})

	AddFuncLookup("beeryeast", Info{
		Display:     "Beer Yeast",
		Category:    "beer",
		Description: "Beer yeast, covering various yeast strains used in beer fermentation",
		Example:     "1388 - Belgian Strong Ale",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			return beerYeast(r), nil
		},
	})

	AddFuncLookup("beermalt", Info{
		Display:     "Beer Malt",
		Category:    "beer",
		Description: "Beer malt, covering various malt types used in the brewing process",
		Example:     "Munich",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			return beerMalt(r), nil
		},
	})

	AddFuncLookup("beeralcohol", Info{
		Display:     "Beer Alcohol",
		Category:    "beer",
		Description: "Alcohol percentage",
		Example:     "2.7%",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			return beerAlcohol(r), nil
		},
	})

	AddFuncLookup("beeribu", Info{
		Display:     "Beer IBU",
		Category:    "beer",
		Description: "International Bitterness Unit (IBU) value, indicating the bitterness level in beer",
		Example:     "29 IBU",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			return beerIbu(r), nil
		},
	})

	AddFuncLookup("beerblg", Info{
		Display:     "Beer BLG",
		Category:    "beer",
		Description: "Generates a random beer Balling (°Blg) value, representing the beer's original gravity or sugar content",
		Example:     "6.4°Blg",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			return beerBlg(r), nil
		},
	})
}
