// Code generated by entc, DO NOT EDIT.

package patient

const (
	// Label holds the string label denoting the patient type in the database.
	Label = "patient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHospitalNumber holds the string denoting the hospitalnumber field in the database.
	FieldHospitalNumber = "hospital_number"
	// FieldPatientName holds the string denoting the patientname field in the database.
	FieldPatientName = "patient_name"
	// FieldDrugAllergy holds the string denoting the drugallergy field in the database.
	FieldDrugAllergy = "drug_allergy"

	// EdgePatientDetails holds the string denoting the patient_details edge name in mutations.
	EdgePatientDetails = "patient_details"
	// EdgeTriageResult holds the string denoting the triageresult edge name in mutations.
	EdgeTriageResult = "triageResult"
	// EdgePatientToAppointmentResults holds the string denoting the patienttoappointmentresults edge name in mutations.
	EdgePatientToAppointmentResults = "PatientToAppointmentResults"
	// EdgePatientToMedicalProcedure holds the string denoting the patienttomedicalprocedure edge name in mutations.
	EdgePatientToMedicalProcedure = "PatientToMedicalProcedure"
	// EdgePatientToRightToTreatment holds the string denoting the patienttorighttotreatment edge name in mutations.
	EdgePatientToRightToTreatment = "PatientToRightToTreatment"
	// EdgePatientToDiagnosis holds the string denoting the patienttodiagnosis edge name in mutations.
	EdgePatientToDiagnosis = "PatientToDiagnosis"

	// Table holds the table name of the patient in the database.
	Table = "patients"
	// PatientDetailsTable is the table the holds the patient_details relation/edge.
	PatientDetailsTable = "patient_details"
	// PatientDetailsInverseTable is the table name for the PatientDetail entity.
	// It exists in this package in order to avoid circular dependency with the "patientdetail" package.
	PatientDetailsInverseTable = "patient_details"
	// PatientDetailsColumn is the table column denoting the patient_details relation/edge.
	PatientDetailsColumn = "patient_patient_details"
	// TriageResultTable is the table the holds the triageResult relation/edge.
	TriageResultTable = "triage_results"
	// TriageResultInverseTable is the table name for the TriageResult entity.
	// It exists in this package in order to avoid circular dependency with the "triageresult" package.
	TriageResultInverseTable = "triage_results"
	// TriageResultColumn is the table column denoting the triageResult relation/edge.
	TriageResultColumn = "patient_triage_result"
	// PatientToAppointmentResultsTable is the table the holds the PatientToAppointmentResults relation/edge.
	PatientToAppointmentResultsTable = "appointment_results"
	// PatientToAppointmentResultsInverseTable is the table name for the AppointmentResults entity.
	// It exists in this package in order to avoid circular dependency with the "appointmentresults" package.
	PatientToAppointmentResultsInverseTable = "appointment_results"
	// PatientToAppointmentResultsColumn is the table column denoting the PatientToAppointmentResults relation/edge.
	PatientToAppointmentResultsColumn = "patient_patient_to_appointment_results"
	// PatientToMedicalProcedureTable is the table the holds the PatientToMedicalProcedure relation/edge.
	PatientToMedicalProcedureTable = "medical_procedures"
	// PatientToMedicalProcedureInverseTable is the table name for the MedicalProcedure entity.
	// It exists in this package in order to avoid circular dependency with the "medicalprocedure" package.
	PatientToMedicalProcedureInverseTable = "medical_procedures"
	// PatientToMedicalProcedureColumn is the table column denoting the PatientToMedicalProcedure relation/edge.
	PatientToMedicalProcedureColumn = "patient_patient_to_medical_procedure"
	// PatientToRightToTreatmentTable is the table the holds the PatientToRightToTreatment relation/edge.
	PatientToRightToTreatmentTable = "right_to_treatments"
	// PatientToRightToTreatmentInverseTable is the table name for the RightToTreatment entity.
	// It exists in this package in order to avoid circular dependency with the "righttotreatment" package.
	PatientToRightToTreatmentInverseTable = "right_to_treatments"
	// PatientToRightToTreatmentColumn is the table column denoting the PatientToRightToTreatment relation/edge.
	PatientToRightToTreatmentColumn = "patient_patient_to_right_to_treatment"
	// PatientToDiagnosisTable is the table the holds the PatientToDiagnosis relation/edge.
	PatientToDiagnosisTable = "diagnoses"
	// PatientToDiagnosisInverseTable is the table name for the Diagnosis entity.
	// It exists in this package in order to avoid circular dependency with the "diagnosis" package.
	PatientToDiagnosisInverseTable = "diagnoses"
	// PatientToDiagnosisColumn is the table column denoting the PatientToDiagnosis relation/edge.
	PatientToDiagnosisColumn = "patient_patient_to_diagnosis"
)

// Columns holds all SQL columns for patient fields.
var Columns = []string{
	FieldID,
	FieldHospitalNumber,
	FieldPatientName,
	FieldDrugAllergy,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// HospitalNumberValidator is a validator for the "hospitalNumber" field. It is called by the builders before save.
	HospitalNumberValidator func(string) error
	// PatientNameValidator is a validator for the "patientName" field. It is called by the builders before save.
	PatientNameValidator func(string) error
	// DrugAllergyValidator is a validator for the "drugAllergy" field. It is called by the builders before save.
	DrugAllergyValidator func(string) error
)
