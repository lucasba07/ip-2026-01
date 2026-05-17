package utils

import (
	"fmt"
	"log"
)

func UpdatePatient(name, cpf, birthDate, bloodType, phone, email, medicalNotes string) error {
	db, err := ConnectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `UPDATE patients 
			  SET name = $1, birth_date = $2, blood_type = $3, phone = $4, email = $5, medical_notes = $6
			  WHERE cpf = $7`
	result, err := db.Exec(query, name, birthDate, bloodType, phone, email, medicalNotes, cpf)
	if err != nil {
		log.Println("Error updating patient:", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("nenhum paciente encontrado com o CPF informado")
	}

	return nil
}
