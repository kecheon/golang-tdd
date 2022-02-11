package hello

import "fmt"

const(
  helloPrefix = "Hello "
  helloPrefixKo = "안녕 "
  helloPrefixFr = "Bonjour "
)

func Hello(name string, lang string) string {
  if name == "" {
    name = "world"
  }

  prefix := greet(lang)

  return prefix + name
}


// declare named return value and return without value returns prefix
// func name starts with lowercase means not public but private func
func greet(lang string) (prefix string) {

  switch lang {
  case "French":
    prefix = helloPrefixFr
  case "Korean":
    prefix = helloPrefixKo
  default:
    prefix = helloPrefix
  }

  return
}

func main() {
	fmt.Println(Hello("world", "Korean"))
}
