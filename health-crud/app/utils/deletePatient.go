package utils

import (
	"fmt"
	"log"
)

func DeletePatient(cpf string) error {
	db, err := ConnectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `DELETE FROM patients WHERE cpf = $1`
	result, err := db.Exec(query, cpf)
	if err != nil {
		log.Println("Error deleting patient:", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("nenhum paciente encontrado com o CPF informado")
	}

	return nil
}
