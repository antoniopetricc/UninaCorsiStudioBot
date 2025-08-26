package services

import (
	"encoding/json"
	"fmt"
	"main/models"

	"net/http"
)

func GetCourses(page int) (models.CoursesResponse, error) {
	url := fmt.Sprintf("https://www.corsi.unina.it/corsidistudio-be/v1/courses?page=%d&size=10", page)

	resp, err := http.Get(url)
	if err != nil {
		return models.CoursesResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.CoursesResponse{}, fmt.Errorf("failed to get courses: %s", resp.Status)
	}

	var result models.CoursesResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return models.CoursesResponse{}, err
	}

	return result, nil
}

/*

{"id":"3D4C36B40A317032E063C91DE18F4C64","content":"Il Corso di studio in Archeologia del Mediterraneo/Mediterranean Archaeology (MediA) (classe LM-02) coniuga la consolidata tradizione di studi e ricerche nelle discipline archeologiche e storico-artistiche dell'Ateneo federiciano, con le nuove linee metodologiche e gli orientamenti del dibattito scientifico più recenti. L'ambito di riferimento è costituito dal bacino del Mediterraneo, un articolato palinsesto storico, archeologico e culturale, analizzato in un ampio arco cronologico, compreso tra l'età preistorica e quella tardo antica e medievale.\nHa come obiettivo la formazione, a livello interdisciplinare, di specialisti nel settore dei beni archeologici che, muovendo da una già acquisita conoscenza delle diverse problematiche dei beni culturali, maturino avanzate competenze di carattere teorico, storico, e critico-metodologico nelle diverse aree e negli ambiti cronologici relativi allo sviluppo dell'archeologia, nonché abilità in ordine alle strategie di conservazione e valorizzazione del patrimonio culturale.\nIl corso prepara alle professioni di: archeologi, curatori e conservatori di musei, esperti d'arte, redattori di testi tecnici, revisore di testi, registar,\nricercatore e tecnico laureato nelle scienze dell'antichità, filologico-letterarie e storico-artistiche, organizzatori di eventi e manifestazioni culturali.\n\n\nL'offerta formativa si sviluppa secondo un piano biennale, strutturato in due semestri per anno. Nel primo anno gli studenti devono superare n. 7 esami, per un totale di 66 cfu, affinché possano consolidare le proprie conoscenze e competenze in relazione alle discipline storiche, linguistiche, letterarie, archeologiche e alla formazione tecnica, scientifica e giuridica, secondo i seguenti ambiti disciplinari:\n1) Storia antica e medievale: potranno sostenere 1 esame a scelta di 6 cfu\n2) Lingue e letterature antiche e medievali: potranno sostenere 1 esame a scelta da 12 cfu\n3) Archeologia e antichità classiche e medievali: potranno sostenere 2 esami a scelta da 12 cfu\n4) Archeologia e antichità classiche e medievali: potranno sostenere 2 esami a scelta da 6 cfu\n5) Formazione tecnica, scientifica e giuridica: potranno sostenere 1 esame a scelta da 12 cfu\n\n\n\nNel secondo anno sono pervisti n. 3 esami (2 esami e la prova finale): un esame affine o integrativo da 12 cfu o 6+6 cfu a scelta tra esami di archeologia, storia, museografia, storia della filosofia e storia dell'arte, affinché ciascun studente, secondo le proprie inclinazioni e interessi personali, possa approfondire una disciplina in particolare; un altro esame da 12 cfu tra gli insegnamenti a scelta in tutto l'Ateneo, fra i quali si consigliano in particolare quelli che consentano l'approfondimento di una lingua straniera. Ultimo esame è la prova finale, pari a 18 cfu, per la quale gli studenti sono tenuti a produrre un elaborato scritto originale (tesi di laurea), frutto di una ricerca approfondita nella disciplina scelta, su tematiche attinenti agli ambiti disciplinari caratterizzanti il Corso di Studio Magistrale e coerenti con gli obiettivi generali della classe.\n\n\nÈ previsto, inoltre, un tirocinio obbligatorio di n. 6 cfu (pari a n. 150 ore): gli studenti, sulla base di un progetto formativo calibrato sugli obiettivi specifici del CdSM e approvato dalla commissione di Coordinamento didattico, possono decidere di effettuare l'attività di tirocinio o stage sia extramoenia, presso enti o aziende italiani o stranieri convenzionati con l'Ateneo, sia intramoenia, sotto la guida di un docente del CdSM. I\nLo studente è tenuto a frequentare il Laboratorio di Digital Humanities da 6 cfu che fornirà agli studenti una comprensione approfondita delle principali metodologie e degli strumenti impiegati per la documentazione digitale, la digitalizzazione e la gestione del patrimonio culturale.\n\n\n","fieldName":"DESC_COR_BRE","title":"Il Corso di Studio in breve","codCorso":"DL6","validityYear":2025,"idSuaCorsi":18024}
*/

func GetCourseDescription(cod string) (models.CourseDescriptionResponse, error) {
	url := fmt.Sprintf("https://www.corsi.unina.it/corsidistudio-be/v1/courses/details/%s/sua/latest/description", cod)

	resp, err := http.Get(url)
	if err != nil {
		return models.CourseDescriptionResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.CourseDescriptionResponse{}, fmt.Errorf("failed to get course description: %s", resp.Status)
	}

	var result models.CourseDescriptionResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return models.CourseDescriptionResponse{}, err
	}

	return result, nil
}
