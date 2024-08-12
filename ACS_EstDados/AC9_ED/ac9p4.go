import "fmt"

func main() {
    // Criando o mapa com as frases de Natal em diferentes idiomas
    frases := map[string]string{
        "brasil":         "Feliz Natal!",
        "alemanha":       "Frohliche Weihnachten!",
        "austria":        "Frohe Weihnacht!",
        "coreia":         "Chuk Sung Tan!",
        "espanha":        "Feliz Navidad!",
        "grecia":         "Kala Christougena!",
        "estados-unidos": "Merry Christmas!",
        "inglaterra":     "Merry Christmas!",
        "australia":      "Merry Christmas!",
        "portugal":       "Feliz Natal!",
        "suecia":         "God Jul!",
        "turquia":        "Mutlu Noeller",
        "argentina":      "Feliz Navidad!",
        "chile":          "Feliz Navidad!",
        "mexico":         "Feliz Navidad!",
        "antardida":      "Merry Christmas!",
        "canada":         "Merry Christmas!",
        "irlanda":        "Nollaig Shona Dhuit!",
        "belgica":        "Zalig Kerstfeest!",
        "italia":         "Buon Natale!",
        "libia":          "Buon Natale!",
        "siria":          "Milad Mubarak!",
        "marrocos":       "Milad Mubarak!",
        "japao":          "Merii Kurisumasu!",
        "india":          "Merry Christmas!",
    }

    // Testando o aplicativo com diferentes nomes de países
    paises := []string{"brasil", "alemanha", "india", "italia", "australia", "canada"}

    for _, pais := range paises {
        if frase, ok := frases[pais]; ok {
            fmt.Println(frase)
        } else {
            fmt.Println("--- NOT FOUND ---")
        }
    }
}