package utils

import (
	"database/sql"
	"log"
)

type Patient struct {
	ID           int
	Name         string
	CPF          string
	BirthDate    string
	BloodType    string
	Phone        string
	Email        string
	MedicalNotes string
	CreatedAt    string
}

func GetPatientByCPF(cpf string) (*Patient, error) {
	db, err := ConnectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var p Patient
	var phone, email, medicalNotes sql.NullString

	query := `SELECT id, name, cpf, TO_CHAR(birth_date, 'YYYY-MM-DD'), blood_type, phone, email, medical_notes, created_at 
			  FROM patients WHERE cpf = $1`
	err = db.QueryRow(query, cpf).Scan(
		&p.ID, &p.Name, &p.CPF, &p.BirthDate, &p.BloodType, &phone, &email, &medicalNotes, &p.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found
		}
		log.Println("Error fetching patient:", err)
		return nil, err
	}

	if phone.Valid { p.Phone = phone.String }
	if email.Valid { p.Email = email.String }
	if medicalNotes.Valid { p.MedicalNotes = medicalNotes.String }

	return &p, nil
}
