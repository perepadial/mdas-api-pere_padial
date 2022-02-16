package user

type PokemonIdCollection struct {
	elements []PokemonId
}

func (collection *PokemonIdCollection) Add(id PokemonId) {
	collection.elements = append(collection.elements, id)
}

func (collection *PokemonIdCollection) Has(id PokemonId) bool {
	elements := collection.elements

	if len(elements) == 0 {
		return true
	}

	for _, pokemonId := range elements {
		if pokemonId.GetValue() == id.GetValue() {
			return false
		}
	}

	return true
}

func CreatePokemonIdCollection(favoritePokemonList []PokemonId) PokemonIdCollection {
	return PokemonIdCollection{favoritePokemonList}
}

func (collection PokemonIdCollection) GetValues() []PokemonId {
	return collection.elements
}
