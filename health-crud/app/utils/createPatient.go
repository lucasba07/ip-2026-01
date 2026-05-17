package utils

import (
	"log"
)

func CreatePatient(name, cpf, birthDate, bloodType, phone, email, medicalNotes string) error {
	db, err := ConnectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO patients (name, cpf, birth_date, blood_type, phone, email, medical_notes) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = db.Exec(query, name, cpf, birthDate, bloodType, phone, email, medicalNotes)
	if err != nil {
		log.Println("Error inserting patient:", err)
		return err
	}
	return nil
}
