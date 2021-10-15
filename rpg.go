package main

import (
	"fmt"
	"os"
	"time"
)

const (
	fourruredeLoup       = "fourrure de loup"
	peaudetroll          = "peau de troll"
	cuirdesanglier       = "cuir de sanglier"
	plumedecorbeau       = "plume de corbeau"
	tuniquedelaventurire = "tunique de l'aventurier"
	chapeaudelaventurier = "chapeau de l'aventurier"
	bottedelaventurier   = "botte de l'aventurier"
	potiondevie          = "potion de vie"
	potiondepoison       = "potion de poison"
	bouledefeu           = "boule de feu"
	tsunami              = "tsunami"
	tonerre              = "tonerre"
	lamedair             = "lame d'air"
	seisme               = "seisme"
	ether                = "ether"
	sword                = "sword"
	coupdepoing          = "coup de poing"
	arc                  = "arc"
	Bahamut              = "Bahamut"
	analyse              = "analyse"
	pvplus               = "pv+"
	pmplus               = "pm+"
	ifrit                = "Ifrit"
	shiva                = "Shiva"
)

type equipement struct {
	tete  string
	torse string
	pieds string
}

func (e *equipement) init(tete string, torse string, pieds string) {
	e.tete = tete
	e.torse = torse
	e.pieds = pieds
}

type perso struct {
	nom        string
	niv        int
	exp        int
	materia    []string
	comp       [80]string
	sort       [80]string
	pvmax      int
	pv         int
	pmmax      int
	pm         int
	pa         int
	pma        int
	vit        int
	gils       int
	inventaire [10]string
	equipement equipement
}

func (p *perso) init(nom string, class string, niv int, exp int, materia []string, comp [80]string, sort [80]string, pvmax int, pv int, pmmax int, pm int, pa int, pma int, vit int, gils int, inventaire [10]string, tete string, torse string, pieds string) {
	p.nom = nom
	p.niv = niv
	p.exp = exp
	p.materia = materia
	p.comp = comp
	p.sort = sort
	p.pvmax = pvmax
	p.pv = pv
	p.pmmax = pmmax
	p.pm = pm
	p.pa = pa
	p.pma = pma
	p.vit = vit
	p.gils = gils
	p.inventaire = inventaire
	p.equipement.tete = tete
	p.equipement.torse = torse
	p.equipement.pieds = pieds
}

type mob struct {
	nom2      string
	pvmax     int
	pva       int
	pa        int
	faiblesse string
	vit       int
	gils      int
	exp       int
}

func (pm *mob) init(nom string, pvmax int, pva int, pa int, faiblesse string, vit int, gils int, exp int) {
	pm.nom2 = nom
	pm.pvmax = pvmax
	pm.pva = pva
	pm.pa = pa
	pm.faiblesse = faiblesse
	pm.vit = vit
	pm.gils = gils
	pm.exp = exp
}

var s int = 3

func (p *perso) menu(pm *mob) {
	var menu string
	fmt.Println("cmd : commande du jeu")
	fmt.Println("")
	fmt.Println("helpm : commande du materia")
	fmt.Println("")
	fmt.Scan(&menu)
	switch menu {
	case "helpm":
		p.materiashelp()
		p.menu(pm)
	case "perso":
		fmt.Println("fiche du perso")
		p.displayinfo()
	case "stuff":
		fmt.Println("equipement")
		p.armure()
	case "forge":
		var test55 int
		fmt.Println("la forge")
		fmt.Println("1 :chapeau de l'aventurier = plume de corbeau, cuir de sanglier et 5 gils ")
		fmt.Println("")
		fmt.Println("2: tunique de l'aventurier = fourrure de loup X2, peau de troll et 5 gils")
		fmt.Println("")
		fmt.Println("3: botte de l'aventurier = fourrure de loup, cuir de sanglier et 5 gils ")
		fmt.Scan(&test55)
		p.forgeron(test55)
	case "cmd":
		fmt.Println(" Les différentes commandes : ")
		fmt.Println("")
		fmt.Println("Information du personnage = perso")
		fmt.Println("")
		fmt.Println("Ouvrir l'inventaire = inv")
		fmt.Println("")
		fmt.Println("liste de materia = materia")
		fmt.Println("")
		fmt.Println("liste de sort = sort")
		fmt.Println("")
		fmt.Println("liste de comp = comp")
		fmt.Println("")
		fmt.Println("faire un combat = fight")
		fmt.Println("")
		fmt.Println("Se rendre chez le marchand = shop")
		fmt.Println("")
		fmt.Println("Se rendre chez le forgeron = forge")
		fmt.Println("")
		fmt.Println("Utiliser une potion de vie = vie")
		fmt.Println("")
		fmt.Println("Utiliser un ether = pm")
		fmt.Println("")
		fmt.Println("Ouvrir le menu d'équipement = stuff")
		fmt.Println("")
		fmt.Println("Equipé un chapeau = tete")
		fmt.Println("")
		fmt.Println("Equipé une tunique = torse")
		fmt.Println("")
		fmt.Println("Equipé des bottes = pieds")
		fmt.Println("")
		fmt.Println("Quitter = fin")
	case "inv":
		fmt.Println("l'inventaire")
		p.accesInventory()
	case "shop":
		var test54 int
		fmt.Println("la boutique")
		fmt.Println("")
		fmt.Println("1 : potion de vie = 3 gils")
		fmt.Println("")
		fmt.Println("2 : potion de poison = 6 gils")
		fmt.Println("")
		fmt.Println("3 : boule de feu = 25 gils")
		fmt.Println("")
		fmt.Println("4 : fourrure de loup = 4 gils")
		fmt.Println("")
		fmt.Println("5 : peau de troll = 7 gils")
		fmt.Println("")
		fmt.Println("6 : cuir de sanglier = 3 gils")
		fmt.Println("")
		fmt.Println("7 : plume de corbeau = 1 gils")
		fmt.Println("")
		fmt.Println("8 : ether = 10 gils")
		fmt.Println("")
		fmt.Println("9 : tonerre = 25 gils")
		fmt.Println("")
		fmt.Println("10 : lame d'air = 25 gils")
		fmt.Println("")
		fmt.Println("11 : tsunami = 25 gils")
		fmt.Println("")
		fmt.Println("12 : seisme = 25 gils")
		fmt.Println("")
		fmt.Println("13 :analyse = 15 gils")
		fmt.Println("")
		fmt.Println("14 : pv+ = 15 gils")
		fmt.Println("")
		fmt.Println("15 : pm+ = 15 gils")
		fmt.Scan(&test54)
		p.marchand(test54)
	case "tete":
		p.Tete()
	case "torse":
		p.Torse()
	case "pieds":
		p.Pieds()
	case "vie":
		fmt.Println(potiondevie)
		p.takepot()
	case "pm":
		fmt.Println(ether)
		p.takeether()
	case "fight":
		var trainingfight string
		fmt.Scan(trainingfight)
		fmt.Println("La bagarre")
		p.trainingfight(pm)
		fmt.Println("Pour retourner a la guerre mon général, go refaire fight esclave!!")
		time.Sleep(3 * time.Second)
	case "Bahamut":
		p.addInventory(Bahamut)
		fmt.Println("félicitation, tu n'as aucun honneur")
		fmt.Println("tappe Bahamut avec t'es pm au max pour invoquer le roi des dragons, il fait 9999 de dégat")
	case "dragon":
		p.speelbahamut()
	case "-dragon":
		p.speelbahamut()
	case "Ifrit":
		p.speelifrit()
	case "-Ifrit":
		p.speelifrit()
	case "Shiva":
		p.speelshiva()
	case "-Shiva":
		p.speelshiva()
	case "feu":
		p.speelfire()
	case "-feu":
		p.removefire()
	case "eau":
		p.speeleau()
	case "-eau":
		p.removeeau()
	case "terre":
		p.speelterre()
	case "-terre":
		p.removeterre()
	case "air":
		p.speelair()
	case "-air":
		p.removeair()
	case "elec":
		p.speelelec()
	case "-elec":
		p.removeelec()
	case "info":
		p.speelinfo()
	case "-info":
		p.removeinfo()
	case "pvplus":
		p.speelpv()
	case "-pvplus":
		p.removepv()
	case "pmplus":
		p.speelpm()
	case "-pmplus":
		p.removepm()
	case "sort":
		fmt.Println("liste de sort")
		p.displaysort()
	case "materia":
		fmt.Println("liste de materia")
		p.displaymateria()
	case "comp":
		fmt.Println("liste de comp")
		p.displaycomp()
	case "fin":
		os.Exit(3)
	}
	p.menu(pm)
}

func main() {
	var al perso
	al.init("Cecil", "", 1, 0, []string{""}, [80]string{""}, [80]string{""}, 100, 50, 100, 50, 5, 16, 10, 100, [10]string{potiondevie}, "", "", "")
	var m mob
	m.init("Gobelin d'entraînement", 40, 40, 5, "feu", 5, 5, 100)
	al.menu(&m)
}

func (p *perso) displayinfo() {
	fmt.Println(" ")
	fmt.Println("/------------------------------------------/")
	fmt.Println("    Écrire 'back' pour revenir au menu ")
	fmt.Println("/------------------------------------------/")
	fmt.Println("nom :", p.nom)
	fmt.Println("niv :", p.niv)
	fmt.Println("pvmax :", p.pvmax)
	fmt.Println("pv :", p.pv)
	fmt.Println("pmmax :", p.pmmax)
	fmt.Println("pm :", p.pm)
	fmt.Println("pa :", p.pa)
	fmt.Println("vit :", p.vit)
	fmt.Println("gils:", p.gils)
}
func (p *perso) displaysort() {
	fmt.Println("sort:", p.sort)
}

func (p *perso) displaymateria() {
	fmt.Println("materia", p.materia)
}

func (p *perso) displaycomp() {
	fmt.Println("comp", p.comp)
}

func (m *mob) displayinfomob() {
	fmt.Println("nom :", m.nom2)
	fmt.Println("pvmax :", m.pvmax)
	fmt.Println("pv :", m.pva)
	fmt.Println("pa :", m.pa)
	fmt.Println("vit :", m.vit)
}

func (p *perso) armure() {
	fmt.Println(" ")
	fmt.Println("/------------------------------------------/")
	fmt.Println("    Écrire 'back' pour revenir au menu ")
	fmt.Println("/------------------------------------------/")
	fmt.Println("tete :", p.equipement.tete)
	fmt.Println("torse :", p.equipement.torse)
	fmt.Println("pieds :", p.equipement.pieds)
}

func (p *perso) levelup() {
	if p.exp >= 100 {
		fmt.Println("félicitation, vous passé niv 2")
		p.niv = 2
		p.pa += 3
		p.vit *= 2
		s += 2
	}
}

func (p *perso) materiashelp() {
	fmt.Println("Les matérias pour les nuls : ")
	fmt.Println("")
	fmt.Println("Définition de matérias : ")
	fmt.Println("")
	fmt.Println("Les matérias sont des objets qui donnent au joueur qui les utilise la possibilité d'utiliser des magies, d'invoquer,d'obtenir des bonus de vies ou d'apprendre des nouvelles compétences..")
	fmt.Println("")
	fmt.Println("Où trouver ses matérias ?")
	fmt.Println("")
	fmt.Println("Les matérias se trouvent chez le marchand")
	fmt.Println("")
	fmt.Println("Quand utiliser ces matérias ?")
	fmt.Println("")
	fmt.Println("Les matérias sont utilisés après les avoirs achetés et équipé. Elles vous serviront durant les combats.")
	fmt.Println("")
	fmt.Println("Et si j'ai pas envie d'en acheter des matérias, tu vas faire quoi ? ")
	fmt.Println("")
	fmt.Println("En effet, je vais absolument rien faire, mais toi non plus. Les matérias sont indispensables. Sans elles, tu n'iras pas loins. Ce sont des bonus qui permettent d'améliorer ton perso. Donc au lieu de faire le grand, profite de ces aides chez le marchand.")
	fmt.Println("")
	fmt.Println("Matérias magies : ")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Boule de feu :")
	fmt.Println("")
	fmt.Println("Equipé : feu")
	fmt.Println("")
	fmt.Println("Déséquipé : -feu")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Tsunami : ")
	fmt.Println("")
	fmt.Println("Equipé : eau")
	fmt.Println("")
	fmt.Println("Déséquipe : -eau")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Séisme :")
	fmt.Println("")
	fmt.Println("Equipé : terre")
	fmt.Println("")
	fmt.Println("Déséquipe : -terre")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Lame d'air :")
	fmt.Println("")
	fmt.Println("Equipé : air")
	fmt.Println("")
	fmt.Println("Déséquipe : -air")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Tonerre :")
	fmt.Println("")
	fmt.Println("Equipé : elec")
	fmt.Println("")
	fmt.Println("Déséquipe : -elec")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Analyse :")
	fmt.Println("")
	fmt.Println("Equipé : info")
	fmt.Println("")
	fmt.Println("Déséquipe : -info")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("PV+ :")
	fmt.Println("")
	fmt.Println("Equipé : pvplus")
	fmt.Println("")
	fmt.Println("Déséquipe : -pvplus")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("PM+ :")
	fmt.Println("")
	fmt.Println("Equipé : pmplus")
	fmt.Println("")
	fmt.Println("Déséquipe : -pmplus")
	fmt.Println("")
	fmt.Println("Bahamut :")
	fmt.Println("")
	fmt.Println("Equipé : dragon")
	fmt.Println("")
	fmt.Println("Déséquipe : -dragon")
	fmt.Println("")
	fmt.Println("Ifrit :")
	fmt.Println("")
	fmt.Println("Equipé :Ifrit")
	fmt.Println("")
	fmt.Println("Déséquipe : -Ifrit")
	fmt.Println("")
	fmt.Println("Shiva :")
	fmt.Println("")
	fmt.Println("Equipé : Shiva")
	fmt.Println("")
	fmt.Println("Déséquipe : -Shiva")
	fmt.Println("")

}

func (p *perso) accesInventory() {
	fmt.Println(" ")
	fmt.Println("/------------------------------------------/")
	fmt.Println("    Écrire 'back' pour revenir au menu ")
	fmt.Println("/------------------------------------------/")
	fmt.Println("Inventaire :")
	fmt.Print(p.inventaire)
}

func (p *perso) takepot() {
	for i := 0; i < len(p.inventaire); i++ {
		if p.pv >= p.pvmax {
			fmt.Println("dsl tu as toute ta vie")
			break
		} else if p.pvmax-p.pv < 50 {
			if p.inventaire[i] == potiondevie {
				p.pv += p.pvmax - p.pv
				fmt.Println("vous regagnez votre vie")
				p.removeInventory(potiondevie)
			} else {
				fmt.Println("dsl tu n'as pas de potion")
			}
			break
		} else if p.pv <= 0 {
			fmt.Println("c'est terminé frère, tape dead pour continuer sauf si t'es vraiment éclaté au sol, alors tape fin")
			break
		} else {
			if p.inventaire[i] == potiondevie {
				p.pv += 50
				fmt.Println("vous regagnez votre vie")
				p.removeInventory(potiondevie)
			} else {
				fmt.Println("dsl tu n'as pas de potion")
			}
			break
		}
	}
}

func (p *perso) takeether() {
	var count2 int = 0
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == ether {
			count2 += 1
			if count2 >= 1 {
				if p.pm >= p.pmmax {
					fmt.Println("dsl tu as toute tes pm")
					break
				} else if p.pmmax-p.pm < 30 {
					if p.inventaire[i] == ether {
						p.pm += p.pmmax - p.pm
						fmt.Println("vous regagnez vos pm")
						p.removeInventory(ether)
					}
				}
				if p.inventaire[i] == ether {
					p.pm += 30
					fmt.Println("vous regagnez votre vie")
					p.removeInventory(ether)
				}
			} else {
				if count2 == 0 {
					fmt.Println("dsl tu n'as pas d'ether")
					break
				}
			}
		}
	}
}

func (p *perso) poisonPot(pm *mob) {
	w := 0
	count := 0
	for i := 0; i < len(p.inventaire); i++ {
		if pm.pva > 0 {
			if p.inventaire[i] == potiondepoison {
				p.removeInventory(potiondepoison)
				for count < 3 {
					pm.pva -= 10
					time.Sleep(1 * time.Second)
					fmt.Println("le monstre à perdue 30 pv")
					w++
					count += 1
				}
				break
			}
		} else if pm.pva <= 0 {
			fmt.Println("mourir empoisoné c'est triste... Mais c'est rigolo pour nous ! ")
			break
		}
	}
}

func (p *perso) dead(pm *mob) {
	if p.pv == 0 {
		fmt.Println("c'est honteux")
		fmt.Println("")
		fmt.Println("si tu galère trop, tappe Bahamut pour te rendre la vie facile")
		fmt.Println("mais bon, t'es pas une loser, t'as quand même un peu d'honneur")
		p.pv = p.pvmax / 2
		pm.pva = pm.pvmax
		p.menu(pm)
	}
}

func (p *perso) addInventory(item string) {
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == "" {
			p.inventaire[i] = item
			break
		}
	}
}

func (p *perso) removeInventory(item string) {
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == item {
			p.inventaire[i] = ""
			break
		}
	}
}

func (p *perso) addsort(item string) {
	for i := 0; i < len(p.sort); i++ {
		if p.sort[i] == "" {
			p.sort[i] = item
			break
		}
	}
}

func (p *perso) removesort(item string) {
	for i := 0; i < len(p.sort); i++ {
		if p.sort[i] == item {
			p.sort[i] = ""
			break
		}
	}
}

func (p *perso) addmateria(item string) {
	if len(p.materia) <= s {
		p.materia = append(p.materia, item)
		p.removeInventory(item)
	}
}

func (p *perso) removemateria(item string) {
	for i := 0; i < len(p.materia); i++ {
		if p.materia[i] == item {
			p.materia[i] = ""
			break
		}
	}
}

func (p *perso) addcomp(item string) {
	for i := 0; i < len(p.comp); i++ {
		if p.comp[i] == "" {
			p.comp[i] = item
			break
		}
	}
}

func (p *perso) removecomp(item string) {
	for i := 0; i < len(p.comp); i++ {
		if p.comp[i] == item {
			p.comp[i] = ""
			break
		}
	}
}

func (p *perso) speelbahamut() {
	for i := 0; i < len(p.inventaire); i++ {
		if len(p.materia) <= s {
			if p.inventaire[i] == "Bahamut" {
				p.removeInventory("Bahamut")
				p.addmateria("Bahamut")
				p.addsort("Bahamut")
				fmt.Println("materia équipé")
			}
		}
	}
}

func (p *perso) removebahamut() {
	for i := 0; i < len(p.materia); i++ {
		if p.materia[i] == "Bahamut" {
			p.addInventory("Bahamut")
			p.removemateria("Bahamut")
			p.removesort("Bahamut")
			fmt.Println("materia déséquipé")
		}
	}
}

func (p *perso) speelifrit() {
	for i := 0; i < len(p.inventaire); i++ {
		if len(p.materia) <= s {
			if p.inventaire[i] == "Ifrit" {
				p.removeInventory("Ifrit")
				p.addmateria("Ifrit")
				p.addsort("Ifrit")
				fmt.Println("materia équipé")
			}
		}
	}
}

func (p *perso) removeifrit() {
	for i := 0; i < len(p.materia); i++ {
		if p.materia[i] == "Ifrit" {
			p.addInventory("Ifrit")
			p.removemateria("Ifrit")
			p.removesort("Ifrit")
			fmt.Println("materia déséquipé")
		}
	}
}

func (p *perso) speelshiva() {
	for i := 0; i < len(p.inventaire); i++ {
		if len(p.materia) <= s {
			if p.inventaire[i] == "Shiva" {
				p.removeInventory("Shiva")
				p.addmateria("Shiva")
				p.addsort("Shiva")
				fmt.Println("materia équipé")
			}
		}
	}
}

func (p *perso) removeshiva() {
	for i := 0; i < len(p.materia); i++ {
		if p.materia[i] == "Shiva" {
			p.addInventory("Shiva")
			p.removemateria("Shiva")
			p.removesort("Shiva")
			fmt.Println("materia déséquipé")
		}
	}
}

func (p *perso) speelfire() {
	for i := 0; i < len(p.inventaire); i++ {
		if len(p.materia) <= s {
			if p.inventaire[i] == "boule de feu" {
				p.removeInventory("boule de feu")
				p.addmateria("boule de feu")
				p.addsort("boule de feu")
				fmt.Println("materia équipé")
			}
		}
	}
}

func (p *perso) removefire() {
	for i := 0; i < len(p.materia); i++ {
		if p.materia[i] == "boule de feu" {
			p.addInventory("boule de feu")
			p.removemateria("boule de feu")
			p.removesort("boule de feu")
			fmt.Println("materia déséquipé")
		}
	}
}

func (p *perso) speelelec() {
	for i := 0; i < len(p.inventaire); i++ {
		if len(p.materia) <= s {
			if p.inventaire[i] == "tonerre" {
				p.removeInventory("tonerre")
				p.addmateria("tonerre")
				p.addsort("tonerre")
				fmt.Println("materia équipé")
			}
		}
	}
}

func (p *perso) removeelec() {
	for i := 0; i < len(p.materia); i++ {
		if p.materia[i] == ("tonerre") {
			p.addInventory("tonerre")
			p.removemateria("tonerre")
			p.removesort("tonerre")
			fmt.Println("materia déséquipé")
		}
	}
}

func (p *perso) speelair() {
	for i := 0; i < len(p.inventaire); i++ {
		if len(p.materia) <= s {
			if p.inventaire[i] == ("lame d'air") {
				p.removeInventory("lame d'air")
				p.addmateria("lame d'air")
				p.addsort("lame d'air")
				fmt.Println("materia équipé")
			}
		}
	}
}

func (p *perso) removeair() {
	for i := 0; i < len(p.materia); i++ {
		if p.materia[i] == ("lame d'air") {
			p.addInventory("lame d'air")
			p.removemateria("lame d'air")
			p.removesort("lame d'air")
			fmt.Println("materia déséquipé")
		}
	}
}

func (p *perso) speelterre() {
	for i := 0; i < len(p.inventaire); i++ {
		if len(p.materia) <= s {
			if p.inventaire[i] == ("seisme") {
				p.removeInventory("seisme")
				p.addmateria("seisme")
				p.addsort("seisme")
				fmt.Println("materia équipé")
			}
		}
	}
}

func (p *perso) removeterre() {
	for i := 0; i < len(p.materia); i++ {
		if p.materia[i] == ("seisme") {
			p.addInventory("seisme")
			p.removemateria("seisme")
			p.removesort("seisme")
			fmt.Println("materia déséquipé")
		}
	}
}

func (p *perso) speeleau() {
	for i := 0; i < len(p.inventaire); i++ {
		if len(p.materia) <= s {
			if p.inventaire[i] == ("tsunami") {
				p.removeInventory("tsunami")
				p.addmateria("tsunami")
				p.addsort("tsunami")
				fmt.Println("materia équipé")
			}
		}
	}
}

func (p *perso) removeeau() {
	for i := 0; i < len(p.materia); i++ {
		if p.materia[i] == ("tsunami") {
			p.addInventory("tsunami")
			p.removemateria("tsunami")
			p.removesort("tsunami")
			fmt.Println("materia déséquipé")
		}
	}
}

func (p *perso) speelinfo() {
	for i := 0; i < len(p.inventaire); i++ {
		if len(p.materia) <= s {
			if p.inventaire[i] == ("analyse") {
				p.removeInventory("analyse")
				p.addmateria("analyse")
				p.addcomp("analyse")
				fmt.Println("materia équipé")
			}
		}
	}
}

func (p *perso) removeinfo() {
	for i := 0; i < len(p.materia); i++ {
		if p.materia[i] == ("analyse") {
			p.addInventory("analyse")
			p.removemateria("analyse")
			p.removecomp("analyse")
			fmt.Println("materia déséquipé")
		}
	}
}

func (p *perso) speelpv() {
	for i := 0; i < len(p.inventaire); i++ {
		if len(p.materia) <= s {
			if p.inventaire[i] == ("pv+") {
				p.removeInventory("pv+")
				p.addmateria("pv+")
				fmt.Println("materia équipé")
				p.pvpv()
			}
		}
	}
}

func (p *perso) removepv() {
	for i := 0; i < len(p.materia); i++ {
		if p.materia[i] == ("pv+") {
			p.addInventory("pv+")
			p.removemateria("pv+")
			fmt.Println("materia déséquipé")
			p.pvpv()
		}
	}
}

func (p *perso) speelpm() {
	for i := 0; i < len(p.inventaire); i++ {
		if len(p.materia) <= s {
			if p.inventaire[i] == ("pm+") {
				p.removeInventory("pm+")
				p.addmateria("pm+")
				fmt.Println("materia équipé")
				p.pmpm()
			}
		}
	}
}

func (p *perso) removepm() {
	for i := 0; i < len(p.materia); i++ {
		if p.materia[i] == ("pm+") {
			p.addInventory("pm+")
			p.removemateria("pm+")
			fmt.Println("materia déséquipé")
			p.pmpm()
		}
	}
}

func (p *perso) marchand(d int) {
	fmt.Println(" ")
	fmt.Println("/------------------------------------------/")
	fmt.Println("    Écrire 'back' pour revenir au menu ")
	fmt.Println("/------------------------------------------/")
	if d == 1 {
		if p.gils >= 3 {
			p.gils = p.gils - 3
			p.addInventory(potiondevie)
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emploi")
		}
	}
	if d == 2 {
		if p.gils >= 6 {
			p.gils = p.gils - 6
			p.addInventory(potiondepoison)
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emploi")
		}
	}
	if d == 3 {
		if p.gils >= 25 {
			p.gils = p.gils - 25
			p.addInventory("boule de feu")
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
		}
	}
	if d == 4 {
		if p.gils >= 4 {
			p.gils = p.gils - 4
			p.addInventory(fourruredeLoup)
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
		}
	}
	if d == 5 {
		if p.gils >= 7 {
			p.gils = p.gils - 7
			p.addInventory(peaudetroll)
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
		}
	}
	if d == 6 {
		if p.gils >= 3 {
			p.gils = p.gils - 3
			p.addInventory(cuirdesanglier)
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
		}
	}
	if d == 7 {
		if p.gils >= 1 {
			p.gils = p.gils - 1
			p.addInventory(plumedecorbeau)
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
		}
	}
	if d == 8 {
		if p.gils >= 10 {
			p.gils = p.gils - 10
			p.addInventory(ether)
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emploi")
		}
	}
	if d == 9 {
		if p.gils >= 25 {
			p.gils = p.gils - 25
			p.addInventory("tonerre")
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
		}
	}
	if d == 10 {
		if p.gils >= 25 {
			p.gils = p.gils - 25
			p.addInventory("ame d'air")
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
		}
	}
	if d == 11 {
		if p.gils >= 25 {
			p.gils = p.gils - 25
			p.addInventory("sunami")
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
		}
	}
	if d == 12 {
		if p.gils >= 25 {
			p.gils = p.gils - 25
			p.addInventory("seisme")
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
		}
	}
	if d == 13 {
		if p.gils >= 15 {
			p.gils = p.gils - 15
			p.addInventory("analyse")
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
		}
	}
	if d == 14 {
		if p.gils >= 15 {
			p.gils = p.gils - 15
			p.addInventory("pv+")
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
		}
	}
	if d == 15 {
		if p.gils >= 15 {
			p.gils = p.gils - 15
			p.addInventory("pm+")
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
		}
	}
	if d == 16 {
		if p.gils >= 50 {
			p.gils = p.gils - 50
			p.addInventory("Ifrit")
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
		}
	}
	if d == 17 {
		if p.gils >= 50 {
			p.gils = p.gils - 50
			p.addInventory("Shiva")
			fmt.Println("achat réussi")
		} else {
			fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
		}
	}
}

func (p *perso) forgeron(d int) {
	fmt.Println(" ")
	fmt.Println("/------------------------------------------/")
	fmt.Println("    Écrire 'back' pour revenir au menu ")
	fmt.Println("/------------------------------------------/")
	countplumedecorbeau := 0
	countfourruredeloup := 0
	countpeaudetroll := 0
	countcuirdesanglier := 0
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == plumedecorbeau {
			countplumedecorbeau++
		} else if p.inventaire[i] == fourruredeLoup {
			countfourruredeloup++
		} else if p.inventaire[i] == cuirdesanglier {
			countcuirdesanglier++
		} else if p.inventaire[i] == peaudetroll {
			countpeaudetroll++
		}
	}
	for i := 0; i < len(p.inventaire); i++ {
		if d == 1 {
			if p.gils >= 5 {
				if countplumedecorbeau >= 1 && countcuirdesanglier >= 1 {
					p.gils = p.gils - 5
					p.removeInventory(plumedecorbeau)
					p.removeInventory(cuirdesanglier)
					p.addInventory(chapeaudelaventurier)
					fmt.Println("objet forgé")
				} else {
					fmt.Println("dsl, t'as pas les objets frere")
				}
				break
			} else {
				fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
			}
			break
		}
		if d == 2 {
			if p.gils >= 5 {
				if countfourruredeloup >= 2 && countpeaudetroll >= 1 {
					p.gils = p.gils - 5
					p.removeInventory(fourruredeLoup)
					p.removeInventory(fourruredeLoup)
					p.removeInventory(peaudetroll)
					p.addInventory(tuniquedelaventurire)
					fmt.Println("objet forgé")
				} else {
					fmt.Println("dsl, t'as pas les objets frere")
				}
				break
			} else {
				fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
			}
			break
		}
		if d == 3 {
			if p.gils >= 5 {
				if countfourruredeloup >= 1 && countcuirdesanglier >= 1 {
					p.gils = p.gils - 5
					p.removeInventory(fourruredeLoup)
					p.removeInventory(cuirdesanglier)
					p.addInventory(bottedelaventurier)
					fmt.Println("objet forgé")
				} else {
					fmt.Println("dsl, t'as pas les objets frere")
				}
				break
			} else {
				fmt.Println("dsl t'es un pauvre, vas pointer à pole emplois")
			}
			break
		}
	}
}

func (p *perso) Tete() {
	countchapeaudelaventurier := 0
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == chapeaudelaventurier {
			if countchapeaudelaventurier < 1 {
				countchapeaudelaventurier++
				p.removeInventory(chapeaudelaventurier)
				p.equipement.tete = chapeaudelaventurier
				fmt.Println("chapeau équipé")
			}
		}
	}
	for i := 0; i < len(p.equipement.tete); i++ {
		if countchapeaudelaventurier == 1 {
			p.pvmax += 10
			break
		}
	}
}

func (p *perso) Torse() {
	counttuniquedelaventurier := 0
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == tuniquedelaventurire {
			if counttuniquedelaventurier < 1 {
				counttuniquedelaventurier++
				p.removeInventory(tuniquedelaventurire)
				p.equipement.torse = tuniquedelaventurire
				fmt.Println("tunique équipé")
			}
		}
	}
	for i := 0; i < len(p.equipement.torse); i++ {
		if counttuniquedelaventurier == 1 {
			p.pvmax += 25
			break
		}
	}
}

func (p *perso) Pieds() {
	countbottelaventurier := 0
	for i := 0; i < len(p.inventaire); i++ {
		if p.inventaire[i] == bottedelaventurier {
			if countbottelaventurier < 1 {
				countbottelaventurier++
				p.removeInventory(bottedelaventurier)
				p.equipement.pieds = bottedelaventurier
				fmt.Println("botte équipé")
			}
		}
	}
	for i := 0; i < len(p.equipement.pieds); i++ {
		if countbottelaventurier == 1 {
			p.pvmax += 15
			break
		}
	}
}

func (p *perso) trainingfight(pm *mob) {
	fmt.Println("")
	fmt.Println("           It's time to la bagarre")
	fmt.Println("")
	fmt.Println("/------------------------------------------/")
	fmt.Println("\n")
	if p.vit >= pm.vit {
		p.charturn(pm)
	} else {
		p.goblinpattern(pm)
	}
}

func (p *perso) charturn(pm *mob) {
	var trainingfight string
	fmt.Scan(&trainingfight)
	switch trainingfight {
	case "attack":
		pm.pva -= p.pa
		fmt.Println("vous attaquez")
		fmt.Println("-", p.pa, " pv")
		fmt.Println("")
		p.dead(pm)
		p.win(pm)
		p.goblinpattern(pm)
	case "info":
		fmt.Println("fiche du mob")
		p.info(pm)
		p.dead(pm)
		p.win(pm)
		p.goblinpattern(pm)
	case "feu":
		fmt.Println("sort lancé")
		fmt.Println("-", p.pma, "pv")
		fmt.Println("")
		p.bouledefeu(pm)
		p.dead(pm)
		p.win(pm)
		p.goblinpattern(pm)
	case "eau":
		fmt.Println("sort lancé")
		fmt.Println("-", p.pma, "pv")
		fmt.Println("")
		p.tsunami(pm)
		p.dead(pm)
		p.win(pm)
		p.goblinpattern(pm)
	case "terre":
		fmt.Println("sort lancé")
		fmt.Println("-", p.pma, "pv")
		fmt.Println("")
		p.seisme(pm)
		p.dead(pm)
		p.win(pm)
		p.goblinpattern(pm)
	case "air":
		fmt.Println("sort lancé")
		fmt.Println("-", p.pma, "pv")
		fmt.Println("")
		p.lamedair(pm)
		p.dead(pm)
		p.win(pm)
		p.goblinpattern(pm)
	case "elec":
		fmt.Println("sort lancé")
		fmt.Println("-", p.pma, "pv")
		fmt.Println("")
		p.tonerre(pm)
		p.dead(pm)
		p.win(pm)
		p.goblinpattern(pm)
	case "sword":
		fmt.Println("sort lancé")
		fmt.Println("-8 pv")
		p.epée(pm)
		p.dead(pm)
		p.win(pm)
		p.goblinpattern(pm)
	case "Bahamut":
		fmt.Println("sort lancé")
		fmt.Println("-9999 pv")
		p.dragon(pm)
		p.dead(pm)
		p.win(pm)
		p.goblinpattern(pm)
	case "Ifrit":
		fmt.Println("sort lancé")
		fmt.Println("-30 pv")
		p.demon(pm)
		p.dead(pm)
		p.win(pm)
		p.goblinpattern(pm)
	case "Shiva":
		fmt.Println("sort lancé")
		fmt.Println("-30 pv")
		p.reine(pm)
		p.dead(pm)
		p.win(pm)
		p.goblinpattern(pm)
	case "vie":
		fmt.Println("potion de vie")
		p.takepot()
		p.dead(pm)
		p.win(pm)
		p.goblinpattern(pm)
	case "pm":
		fmt.Println("ether")
		p.takeether()
		p.dead(pm)
		p.win(pm)
		p.goblinpattern(pm)
	case "poi":
		fmt.Println("potion de poison")
		p.poisonPot(pm)
		p.dead(pm)
		p.win(pm)
		p.goblinpattern(pm)
	case "inv":
		fmt.Println("l'inventaire")
		p.accesInventory()
		p.charturn(pm)
	case "perso":
		fmt.Println("fiche du personnage")
		p.displayinfo()
		p.charturn(pm)
	case "cmd":
		fmt.Println("Fiche du personage = perso")
		fmt.Println("")
		fmt.Println("attaquer= attack")
		fmt.Println("")
		fmt.Println("Fiche du mob= info")
		fmt.Println("")
		fmt.Println("potion de vie = vie")
		fmt.Println("")
		fmt.Println("ether = pm")
		fmt.Println("")
		fmt.Println("l'inventaire = inv")
		fmt.Println("")
		fmt.Println("potion de poison = poi")
		fmt.Println("")
		fmt.Println("sword = sword")
		fmt.Println("")
		fmt.Println("boule de feu = fire")
		fmt.Println("")
		fmt.Println("fuir = back")
		fmt.Println("")
		p.charturn(pm)
	case "back":
		fmt.Println("vous fuyez comme un lache")
		pm.pva = pm.pvmax
		p.menu(pm)
	}
}

func (p *perso) win(pm *mob) {
	if pm.pva <= 0 {
		fmt.Println("bravo t'as gagné")
		p.gils += pm.gils
		p.exp += pm.exp
		fmt.Println("tu as gagné", pm.gils, " gils")
		fmt.Println("")
		fmt.Println("+", pm.exp, " exp")
		pm.pva = pm.pvmax
		p.levelup()
		p.menu(pm)
	}
}

var test2 int = 0
var tours int

func (p *perso) goblinpattern(pm *mob) {
	if tours%3 == 2 {
		p.pv -= 10
		fmt.Println("le gobelin vous attaque")
		fmt.Println("-10 pv")
		fmt.Println("")
		tours = 0
		p.charturn(pm)
	} else {
		p.pv -= 5
		fmt.Println("le gobelin vous attaque")
		fmt.Println("-5 pv")
		tours++
		p.charturn(pm)
	}
}

func (p *perso) bouledefeu(pm *mob) {
	countbouledefeu := 0
	for i := 0; i < len(p.sort); i++ {
		if p.sort[i] == bouledefeu {
			countbouledefeu++
		}
	}
	for i := 0; i < len(p.sort); i++ {
		if p.pm >= 20 {
			if countbouledefeu >= 1 {
				pm.pva -= p.pma
				p.pm -= 20
			} else {
				fmt.Println("dsl achete le sort")
			}
			break
		} else {
			fmt.Println("dsl t'as plus de pm")
		}
		break
	}
}

func (p *perso) epée(pm *mob) {
	countsword := 0
	for i := 0; i < len(p.sort); i++ {
		if p.sort[i] == sword {
			countsword++
		}
	}
	for i := 0; i < len(p.sort); i++ {
		if p.pm >= 10 {
			if countsword >= 1 {
				pm.pva -= 8
				p.pm -= 10
			} else {
				fmt.Println("dsl achete le sort")
			}
			break
		} else {
			fmt.Println("dsl t'as plus de pm")
		}
		break
	}
}

func (p *perso) dragon(pm *mob) {
	countBahamut := 0
	for i := 0; i < len(p.sort); i++ {
		if p.sort[i] == Bahamut {
			countBahamut++
		}
	}
	for i := 0; i < len(p.sort); i++ {
		if p.pm >= 100 {
			if countBahamut >= 1 {
				pm.pva -= 9999
				p.pm -= 100
			} else {
				fmt.Println("dsl achete le sort")
			}
			break
		} else {
			fmt.Println("dsl t'as plus de pm")
		}
		break
	}
}

func (p *perso) demon(pm *mob) {
	countifrit := 0
	for i := 0; i < len(p.sort); i++ {
		if p.sort[i] == ifrit {
			countifrit++
		}
	}
	for i := 0; i < len(p.sort); i++ {
		if p.pm >= 50 {
			if countifrit >= 1 {
				pm.pva -= 35
				p.pm -= 50
			} else {
				fmt.Println("dsl achete le sort")
			}
			break
		} else {
			fmt.Println("dsl t'as plus de pm")
		}
		break
	}
}

func (p *perso) reine(pm *mob) {
	countshiva := 0
	for i := 0; i < len(p.sort); i++ {
		if p.sort[i] == shiva {
			countshiva++
		}
	}
	for i := 0; i < len(p.sort); i++ {
		if p.pm >= 50 {
			if countshiva >= 1 {
				pm.pva -= 35
				p.pm -= 50
			} else {
				fmt.Println("dsl achete le sort")
			}
			break
		} else {
			fmt.Println("dsl t'as plus de pm")
		}
		break
	}
}

func (p *perso) tonerre(pm *mob) {
	counttonerre := 0
	for i := 0; i < len(p.sort); i++ {
		if p.sort[i] == tonerre {
			counttonerre++
		}
	}
	for i := 0; i < len(p.sort); i++ {
		if p.pm >= 20 {
			if counttonerre >= 1 {
				pm.pva -= p.pma
				p.pm -= 20
			} else {
				fmt.Println("dsl achete le sort")
			}
			break
		} else {
			fmt.Println("dsl t'as plus de pm")
		}
		break
	}
}

func (p *perso) lamedair(pm *mob) {
	countlamedair := 0
	for i := 0; i < len(p.sort); i++ {
		if p.sort[i] == lamedair {
			countlamedair++
		}
	}
	for i := 0; i < len(p.sort); i++ {
		if p.pm >= 20 {
			if countlamedair >= 1 {
				pm.pva -= p.pma
				p.pm -= 20
			} else {
				fmt.Println("dsl achete le sort")
			}
			break
		} else {
			fmt.Println("dsl t'as plus de pm")
		}
		break
	}
}

func (p *perso) tsunami(pm *mob) {
	counttsunami := 0
	for i := 0; i < len(p.sort); i++ {
		if p.sort[i] == tsunami {
			counttsunami++
		}
	}
	for i := 0; i < len(p.sort); i++ {
		if p.pm >= 20 {
			if counttsunami >= 1 {
				pm.pva -= p.pma
				p.pm -= 20
			} else {
				fmt.Println("dsl achete le sort")
			}
			break
		} else {
			fmt.Println("dsl t'as plus de pm")
		}
		break
	}
}

func (p *perso) seisme(pm *mob) {
	countseisme := 0
	for i := 0; i < len(p.sort); i++ {
		if p.sort[i] == seisme {
			countseisme++
		}
	}
	for i := 0; i < len(p.sort); i++ {
		if p.pm >= 20 {
			if countseisme >= 1 {
				pm.pva -= p.pma
				p.pm -= 20
			} else {
				fmt.Println("dsl achete le sort")
			}
			break
		} else {
			fmt.Println("dsl t'as plus de pm")
		}
		break
	}
}

func (p *perso) info(pm *mob) {
	countanalyse := 0
	for i := 0; i < len(p.comp); i++ {
		if p.comp[i] == analyse {
			countanalyse++
		}
	}
	for i := 0; i < len(p.comp); i++ {
		if countanalyse >= 1 {
			pm.displayinfomob()
		}
		break
	}
}

func (p *perso) pvpv() {
	countpvplus := 0
	for i := 0; i < len(p.materia); i++ {
		if p.materia[i] == pvplus {
			countpvplus++
			break
		}
	}
	for i := 0; i < len(p.materia); i++ {
		if countpvplus == 1 {
			p.pvmax += 5
		} else {
			if countpvplus == 2 {
				p.pvmax += 30
			} else {
				if countpvplus == 3 {
					p.pvmax += 30
				}
			}
		}
	}
}

func (p *perso) pmpm() {
	countpmplus := 0
	for i := 0; i < len(p.materia); i++ {
		if p.materia[i] == pmplus {
			countpmplus++
			break
		}
	}
	for i := 0; i < len(p.materia); i++ {
		if countpmplus == 1 {
			p.pvmax += 5
		} else {
			if countpmplus == 2 {
				p.pvmax += 30
			} else {
				if countpmplus == 3 {
					p.pvmax += 30
				}
			}
		}
	}
}
