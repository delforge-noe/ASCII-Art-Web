package gene

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// --- Variables globales ---
var alphabets = make(map[string]map[rune][]string)

// Liste des alphabets disponibles
var fonts = []string{"standard", "shadow", "thinkertoy"}

// Chargement automatique au démarrage
func init() {
	loadAlphabets()
}

// -----------------------------
//     LECTURE DES FICHIERS
// -----------------------------

// Lit un fichier et renvoie une map[rune][]string
func readBannerFile(filename string) (map[rune][]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir %s : %w", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	charMap := make(map[rune][]string)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines)%9 != 0 {
		return nil, fmt.Errorf("fichier %s mal formaté", filename)
	}

	// Chaque caractère fait 8 lignes + 1 ligne vide
	for i := 0; i < len(lines); i += 9 {
		charIndex := i / 9
		r := rune(32 + charIndex) // caractères ASCII du 32 au 126
		charMap[r] = lines[i : i+8]
	}

	return charMap, nil
}

// Charge tous les alphabets du dossier banners/
func loadAlphabets() {
	for _, f := range fonts {
		path := "banners/" + f + ".txt"

		charMap, err := readBannerFile(path)
		if err != nil {
			fmt.Println("Erreur chargement :", err)
			continue
		}

		alphabets[f] = charMap
	}

	fmt.Println("✔ Alphabets chargés :", len(alphabets))
}

// -----------------------------
//     GENERATION ASCII-ART
// -----------------------------

func GenerateASCIIArt(text string, font string) (string, error) {

	charMap := alphabets[font]

	text = strings.ReplaceAll(text, "\r\n", "\n")
	lines := strings.Split(text, "\n")

	var result strings.Builder

	for _, line := range lines {

		if strings.TrimSpace(line) == "" {
			// 8 retours à la ligne si vide
			for i := 0; i < 8; i++ {
				result.WriteString("\n")
			}
			continue
		}

		for row := 0; row < 8; row++ {
			for _, r := range line {

				ascii, ok := charMap[r]
				if !ok {
					return "", fmt.Errorf("caractère '%c' non supporté dans '%s'", r, font)
				}

				result.WriteString(ascii[row])
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}
