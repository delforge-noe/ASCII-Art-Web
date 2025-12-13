package main

import (
	"ascii-art-web/gene"
	"html/template"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var tpl *template.Template

func Init() {
	rand.Seed(time.Now().UnixNano())
}

var PlaceHolderImages = []string{"image1.png", "image2.png", "image3.png"}

func main() {

	// Chargement du template au démarrage.
	tpl = template.Must(template.ParseFiles("templates/index.html"))

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Route principale gérée par la fonction handler
	http.HandleFunc("/", handler)

	// Lancer le serveur
	println("Serveur lancé : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	// 1. GESTION DU 404 Not Found
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found: La page demandée n'existe pas", http.StatusNotFound)
		return
	}

	RandIndex := rand.Intn(len(PlaceHolderImages))
	RandImage := PlaceHolderImages[RandIndex]

	data := struct {
		Input        string
		Alphabet     string
		Result       string
		Error        string
		InputIsEmpty bool // Ajout du flag pour le rendu personnalisé
		RandomImage  string
	}{
		Alphabet:    "standard",
		RandomImage: RandImage,
	}

	// GESTION DU POST
	if r.Method == http.MethodPost {

		// 1. GESTION DU 400 Bad Request - Parsing
		if err := r.ParseForm(); err != nil {
			http.Error(w, "400 Bad Request: Données de formulaire illisibles", http.StatusBadRequest)
			return
		}

		data.Input = strings.TrimSpace(r.FormValue("input_text"))
		data.Alphabet = r.FormValue("alphabet_choisi")

		// 2. LOGIQUE PERSONNALISÉE pour input vide
		if data.Input == "" {
			data.InputIsEmpty = true
			// Le code continue pour exécuter le template et afficher l'erreur personnalisée

		} else {
			// 3. GESTION DU 400 Bad Request - Validation des bannières
			validBanners := map[string]bool{"standard": true, "shadow": true, "thinkertoy": true}
			if !validBanners[data.Alphabet] {
				http.Error(w, "400 Bad Request: Bannière non valide.", http.StatusBadRequest)
				return
			}

			// 4. Appel au générateur ASCII
			result, err := gene.GenerateASCIIArt(data.Input, data.Alphabet)
			if err != nil {
				// Erreur de génération (mauvaise donnée utilisateur) -> 400.
				http.Error(w, "400 Bad Request: Erreur de génération ASCII Art ("+err.Error()+")", http.StatusBadRequest)
				return
			}

			// Si tout est OK
			data.Result = result
		}
	}

	// 5. GESTION DU 500 Internal Server Error & 200 OK
	err := tpl.Execute(w, data)
	if err != nil {
		http.Error(w, "500 Internal Server Error: Erreur lors du rendu de la page.", http.StatusInternalServerError)
		return
	}
}
