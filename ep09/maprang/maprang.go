package main

func main() {
	langs := map[string]int{
		"Go":      1,
		"Rust":    2,
		"Elixir":  3,
		"Pythoon": 4, //comma here not required if } is in this line
	}

	// keys and values
	for k, v := range langs {
		println(k, v)
	}

	//only values - Key dont care
	for _, v := range langs {
		println(v)
	}

	//only keys _ for value is not required
	for k := range langs {
		println(k)
	}

}
