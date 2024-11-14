package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time" //init gnér de nb alé avq un seed 
)

func main() {
	fichierMots, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fichierMots.Close()

	var lignes []string
	scanner := bufio.NewScanner(fichierMots) //ouvre le fichier de mots 
	for scanner.Scan() {
		lignes = append(lignes, scanner.Text()) //va chercher une ligne au hasard et recupère le texte qui a dedans 
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}  // annonce l'erreur en cas d'erreur 

	rand.Seed(time.Now().UnixNano())    //génère un numero aléatoire 
	motAleatoire := lignes[rand.Intn(len(lignes))] // le mot aleatoire sera = au mot qui est dans la ligne du numero aleatoire
	motAleatoire = strings.ToLower(motAleatoire)  // le mot aleatoire sera = au mot aleatoire peut importe si il est en majuscule ou en minuscule 

	mot := make([]rune, len(motAleatoire)) // parcours du mot sélectionne ses 3dernières lettres et placer des tirets devant selon son nombre de lettres total
	for i := range mot {
		if i >= len(motAleatoire)-3 {
			mot[i] = rune(motAleatoire[i])
		} else {
			mot[i] = '_'
		}
	} //boucle for parcours et compte les l, si le mot a + que 3l toutes les lettres sont remplacees par t du bas 

	fmt.Println("Mot à deviner :", string(mot)) // initialise le jeu et donne le mot caché 

	tentatives := 10 // donne le nb de tentatives 
	etapes := lireEtapesPendu("hangmanposition.txt") // va chercher dans le fichier hangman position les differentes etapes pour print la bonne etape

	for tentatives > 0 {
		var proposition string
		fmt.Print("Entrez une lettre ou devinez le mot complet : ")
		fmt.Scanln(&proposition)
		proposition = strings.ToLower(proposition)

		if len(proposition) > 1 {
			if proposition == motAleatoire {
				fmt.Println("Félicitations ! Vous avez deviné le mot :", motAleatoire)
				return
			} else {
				fmt.Println("Mot incorrect !", tentatives, "tentatives restantes")
				tentatives--
				if tentatives < len(etapes) {
					fmt.Println(etapes[len(etapes)-tentatives-1])
				}
			}
		} else {
			bonneProposition := false
			for i := 0; i < len(motAleatoire)-3; i++ {
				if rune(proposition[0]) == rune(motAleatoire[i]) && mot[i] == '_' {
					mot[i] = rune(motAleatoire[i])
					bonneProposition = true
				}
			}

			if bonneProposition {
				fmt.Println("Bonne lettre !", string(mot))
			} else {
				fmt.Println("Lettre incorrecte !", tentatives, "tentatives restantes")
				tentatives--
				if tentatives < len(etapes) {
					fmt.Println(etapes[len(etapes)-tentatives-1])
				}
			}

			if string(mot) == motAleatoire {
				fmt.Println("Félicitations ! Vous avez deviné le mot :", motAleatoire)
				return
			}
		}
	}

	fmt.Println("Fin de partie ! Le mot était :", motAleatoire)
}

func lireEtapesPendu(nomFichier string) []string {
	fichier, err := os.Open(nomFichier)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return nil
	}
	defer fichier.Close()

	scanner := bufio.NewScanner(fichier)
	var etapes []string
	var etapeActuelle strings.Builder

	for scanner.Scan() {
		ligne := scanner.Text()
		if ligne == "" {
			etapes = append(etapes, etapeActuelle.String())
			etapeActuelle.Reset()
		} else {
			etapeActuelle.WriteString(ligne + "\n")
		}
	}

	if etapeActuelle.Len() > 0 {
		etapes = append(etapes, etapeActuelle.String())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur de scanner :", err)
	}

	return etapes
}
