package controllers

import (
    "encoding/json"
    "myapp/config"
    "myapp/models"
    "myapp/services"
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "strconv"
)

// RegisterPatient - Controller to register a new patient
func RegisterPatient(w http.ResponseWriter, r *http.Request) {
    var patient models.Patient
    if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := services.RegisterPatient(patient); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(patient)
}

// GetPatient - Controller to get a patient by ID
func GetPatient(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    patient, err := services.GetPatientByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(patient)
}

// UpdatePatient - Controller to update a patient's contact information
func UpdatePatient(w http.ResponseWriter, r *http.Request) {
    // Get the patient ID from the URL
    vars := mux.Vars(r)
    patientID := vars["id"]

    var patient models.Patient

    // Decode the incoming request body to get updated contact details
    err := json.NewDecoder(r.Body).Decode(&patient)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Update the patient's contact information in the database
    query := `UPDATE patients SET contact = $1 WHERE id = $2`
    _, err = config.DB.Exec(query, patient.Contact, patientID)
    if err != nil {
        log.Printf("Database error: %v", err)
        http.Error(w, "Failed to update patient contact", http.StatusInternalServerError)
        return
    }

    // Respond with success message
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Patient contact updated successfully"))
}

// DeletePatient - Controller to delete a patient's record by ID
func DeletePatient(w http.ResponseWriter, r *http.Request) {
    // Get the patient ID from the URL
    vars := mux.Vars(r)
    patientID := vars["id"]

    // Execute the DELETE query to remove the patient from the database
    query := `DELETE FROM patients WHERE id = $1`
    _, err := config.DB.Exec(query, patientID)
    if err != nil {
        log.Printf("Database error: %v", err)
        http.Error(w, "Failed to delete patient", http.StatusInternalServerError)
        return
    }

    // Respond with a success message
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Patient deleted successfully"))
}
