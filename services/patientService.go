package services

import (
    "errors"
    "myapp/config"
    "myapp/models"
)

// RegisterPatient - Register a new patient in the database
func RegisterPatient(patient models.Patient) error {
    db := config.DB

    // SQL query to insert patient data
    _, err := db.Exec("INSERT INTO patients (name, age, gender, contact) VALUES ($1, $2, $3, $4)", 
                      patient.Name, patient.Age, patient.Gender, patient.Contact)
    if err != nil {
        return err
    }

    return nil
}

// GetPatientByID - Get a patient by their ID
func GetPatientByID(id int) (models.Patient, error) {
    var patient models.Patient
    db := config.DB

    // SQL query to retrieve patient data by ID
    row := db.QueryRow("SELECT id, name, age, gender, contact FROM patients WHERE id=$1", id)

    // Scan the result into the patient model
    if err := row.Scan(&patient.ID, &patient.Name, &patient.Age, &patient.Gender, &patient.Contact); err != nil {
        if err.Error() == "no rows in result set" {
            return patient, errors.New("patient not found")
        }
        return patient, err
    }

    return patient, nil
}

// UpdatePatient - Update patient information
func UpdatePatient(id int, patient models.Patient) error {
    db := config.DB

    // SQL query to update patient data
    _, err := db.Exec("UPDATE patients SET name=$1, age=$2, gender=$3, contact=$4 WHERE id=$5", 
                      patient.Name, patient.Age, patient.Gender, patient.Contact, id)
    if err != nil {
        return err
    }

    return nil
}

// DeletePatient - Delete a patient by ID
func DeletePatient(id int) error {
    db := config.DB

    // SQL query to delete patient by ID
    _, err := db.Exec("DELETE FROM patients WHERE id=$1", id)
    if err != nil {
        return err
    }

    return nil
}
