package config

import "fmt"

func Init() {
	asciiInitialize :=
		"__        __   _    ____  _       _     _   \n" +
			"\\ \\      / /__| |__/ ___|(_) __ _| |__ | |_ \n" +
			" \\ \\ /\\ / / _ \\ '_ \\___ \\| |/ _` | '_ \\| __|\n" +
			"  \\ V  V /  __/ |_) |__) | | (_| | | | | |_ \n" +
			"   \\_/\\_/ \\___|_.__/____/|_|\\__, |_| |_|\\__|\n" +
			"                            |___/           \n" +
			"                                               v0.1.0"

	sentence := "\n" +
		"author: dawnaliens\n" +
		"github:https://github.com/dawnaliens/WebSight\n" +
		"\n" +
		"\n" +
		"\n" +
		"Use with caution. You are responsible for your actions" +
		"\n"

	fmt.Println(asciiInitialize)
	fmt.Println(sentence)
}
