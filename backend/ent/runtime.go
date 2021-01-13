// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/team06/app/ent/appointmentresults"
	"github.com/team06/app/ent/bloodtype"
	"github.com/team06/app/ent/department"
	"github.com/team06/app/ent/diagnosis"
	"github.com/team06/app/ent/doctor"
	"github.com/team06/app/ent/gender"
	"github.com/team06/app/ent/hospital"
	"github.com/team06/app/ent/medicalrecord"
	"github.com/team06/app/ent/nurse"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/prefix"
	"github.com/team06/app/ent/righttotreatmenttype"
	"github.com/team06/app/ent/room"
	"github.com/team06/app/ent/schema"
	"github.com/team06/app/ent/treatmenttype"
	"github.com/team06/app/ent/triageresult"
	"github.com/team06/app/ent/urgencylevel"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	appointmentresultsFields := schema.AppointmentResults{}.Fields()
	_ = appointmentresultsFields
	// appointmentresultsDescAddtimeSave is the schema descriptor for addtimeSave field.
	appointmentresultsDescAddtimeSave := appointmentresultsFields[1].Descriptor()
	// appointmentresults.DefaultAddtimeSave holds the default value on creation for the addtimeSave field.
	appointmentresults.DefaultAddtimeSave = appointmentresultsDescAddtimeSave.Default.(func() time.Time)
	bloodtypeFields := schema.BloodType{}.Fields()
	_ = bloodtypeFields
	// bloodtypeDescBloodValue is the schema descriptor for bloodValue field.
	bloodtypeDescBloodValue := bloodtypeFields[0].Descriptor()
	// bloodtype.BloodValueValidator is a validator for the "bloodValue" field. It is called by the builders before save.
	bloodtype.BloodValueValidator = bloodtypeDescBloodValue.Validators[0].(func(string) error)
	departmentFields := schema.Department{}.Fields()
	_ = departmentFields
	// departmentDescDepartmentName is the schema descriptor for departmentName field.
	departmentDescDepartmentName := departmentFields[0].Descriptor()
	// department.DepartmentNameValidator is a validator for the "departmentName" field. It is called by the builders before save.
	department.DepartmentNameValidator = departmentDescDepartmentName.Validators[0].(func(string) error)
	diagnosisFields := schema.Diagnosis{}.Fields()
	_ = diagnosisFields
	// diagnosisDescSymptom is the schema descriptor for symptom field.
	diagnosisDescSymptom := diagnosisFields[0].Descriptor()
	// diagnosis.SymptomValidator is a validator for the "symptom" field. It is called by the builders before save.
	diagnosis.SymptomValidator = diagnosisDescSymptom.Validators[0].(func(string) error)
	// diagnosisDescOpinionresult is the schema descriptor for Opinionresult field.
	diagnosisDescOpinionresult := diagnosisFields[1].Descriptor()
	// diagnosis.OpinionresultValidator is a validator for the "Opinionresult" field. It is called by the builders before save.
	diagnosis.OpinionresultValidator = diagnosisDescOpinionresult.Validators[0].(func(string) error)
	// diagnosisDescDiagnosisDate is the schema descriptor for diagnosisDate field.
	diagnosisDescDiagnosisDate := diagnosisFields[2].Descriptor()
	// diagnosis.DefaultDiagnosisDate holds the default value on creation for the diagnosisDate field.
	diagnosis.DefaultDiagnosisDate = diagnosisDescDiagnosisDate.Default.(func() time.Time)
	doctorFields := schema.Doctor{}.Fields()
	_ = doctorFields
	// doctorDescDoctorName is the schema descriptor for doctorName field.
	doctorDescDoctorName := doctorFields[0].Descriptor()
	// doctor.DoctorNameValidator is a validator for the "doctorName" field. It is called by the builders before save.
	doctor.DoctorNameValidator = doctorDescDoctorName.Validators[0].(func(string) error)
	// doctorDescDoctorUsername is the schema descriptor for doctorUsername field.
	doctorDescDoctorUsername := doctorFields[1].Descriptor()
	// doctor.DoctorUsernameValidator is a validator for the "doctorUsername" field. It is called by the builders before save.
	doctor.DoctorUsernameValidator = doctorDescDoctorUsername.Validators[0].(func(string) error)
	// doctorDescDoctorPassword is the schema descriptor for doctorPassword field.
	doctorDescDoctorPassword := doctorFields[2].Descriptor()
	// doctor.DoctorPasswordValidator is a validator for the "doctorPassword" field. It is called by the builders before save.
	doctor.DoctorPasswordValidator = doctorDescDoctorPassword.Validators[0].(func(string) error)
	genderFields := schema.Gender{}.Fields()
	_ = genderFields
	// genderDescGenderValue is the schema descriptor for genderValue field.
	genderDescGenderValue := genderFields[0].Descriptor()
	// gender.GenderValueValidator is a validator for the "genderValue" field. It is called by the builders before save.
	gender.GenderValueValidator = genderDescGenderValue.Validators[0].(func(string) error)
	hospitalFields := schema.Hospital{}.Fields()
	_ = hospitalFields
	// hospitalDescHospitalName is the schema descriptor for HospitalName field.
	hospitalDescHospitalName := hospitalFields[0].Descriptor()
	// hospital.HospitalNameValidator is a validator for the "HospitalName" field. It is called by the builders before save.
	hospital.HospitalNameValidator = hospitalDescHospitalName.Validators[0].(func(string) error)
	medicalrecordFields := schema.MedicalRecord{}.Fields()
	_ = medicalrecordFields
	// medicalrecordDescEmail is the schema descriptor for email field.
	medicalrecordDescEmail := medicalrecordFields[0].Descriptor()
	// medicalrecord.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	medicalrecord.EmailValidator = medicalrecordDescEmail.Validators[0].(func(string) error)
	// medicalrecordDescPassword is the schema descriptor for password field.
	medicalrecordDescPassword := medicalrecordFields[1].Descriptor()
	// medicalrecord.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	medicalrecord.PasswordValidator = medicalrecordDescPassword.Validators[0].(func(string) error)
	// medicalrecordDescName is the schema descriptor for name field.
	medicalrecordDescName := medicalrecordFields[2].Descriptor()
	// medicalrecord.NameValidator is a validator for the "name" field. It is called by the builders before save.
	medicalrecord.NameValidator = medicalrecordDescName.Validators[0].(func(string) error)
	nurseFields := schema.Nurse{}.Fields()
	_ = nurseFields
	// nurseDescNurseName is the schema descriptor for nurseName field.
	nurseDescNurseName := nurseFields[0].Descriptor()
	// nurse.NurseNameValidator is a validator for the "nurseName" field. It is called by the builders before save.
	nurse.NurseNameValidator = nurseDescNurseName.Validators[0].(func(string) error)
	// nurseDescNurseUsername is the schema descriptor for nurseUsername field.
	nurseDescNurseUsername := nurseFields[1].Descriptor()
	// nurse.NurseUsernameValidator is a validator for the "nurseUsername" field. It is called by the builders before save.
	nurse.NurseUsernameValidator = nurseDescNurseUsername.Validators[0].(func(string) error)
	// nurseDescNursePassword is the schema descriptor for nursePassword field.
	nurseDescNursePassword := nurseFields[2].Descriptor()
	// nurse.NursePasswordValidator is a validator for the "nursePassword" field. It is called by the builders before save.
	nurse.NursePasswordValidator = nurseDescNursePassword.Validators[0].(func(string) error)
	patientFields := schema.Patient{}.Fields()
	_ = patientFields
	// patientDescPersonalID is the schema descriptor for personalID field.
	patientDescPersonalID := patientFields[0].Descriptor()
	// patient.PersonalIDValidator is a validator for the "personalID" field. It is called by the builders before save.
	patient.PersonalIDValidator = patientDescPersonalID.Validators[0].(func(int) error)
	// patientDescPatientName is the schema descriptor for patientName field.
	patientDescPatientName := patientFields[1].Descriptor()
	// patient.PatientNameValidator is a validator for the "patientName" field. It is called by the builders before save.
	patient.PatientNameValidator = patientDescPatientName.Validators[0].(func(string) error)
	// patientDescAge is the schema descriptor for age field.
	patientDescAge := patientFields[2].Descriptor()
	// patient.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	patient.AgeValidator = patientDescAge.Validators[0].(func(int) error)
	// patientDescHospitalNumber is the schema descriptor for hospitalNumber field.
	patientDescHospitalNumber := patientFields[3].Descriptor()
	// patient.HospitalNumberValidator is a validator for the "hospitalNumber" field. It is called by the builders before save.
	patient.HospitalNumberValidator = patientDescHospitalNumber.Validators[0].(func(string) error)
	// patientDescDrugAllergy is the schema descriptor for drugAllergy field.
	patientDescDrugAllergy := patientFields[4].Descriptor()
	// patient.DrugAllergyValidator is a validator for the "drugAllergy" field. It is called by the builders before save.
	patient.DrugAllergyValidator = patientDescDrugAllergy.Validators[0].(func(string) error)
	// patientDescAddedDate is the schema descriptor for addedDate field.
	patientDescAddedDate := patientFields[5].Descriptor()
	// patient.DefaultAddedDate holds the default value on creation for the addedDate field.
	patient.DefaultAddedDate = patientDescAddedDate.Default.(func() time.Time)
	prefixFields := schema.Prefix{}.Fields()
	_ = prefixFields
	// prefixDescPrefixValue is the schema descriptor for prefixValue field.
	prefixDescPrefixValue := prefixFields[0].Descriptor()
	// prefix.PrefixValueValidator is a validator for the "prefixValue" field. It is called by the builders before save.
	prefix.PrefixValueValidator = prefixDescPrefixValue.Validators[0].(func(string) error)
	righttotreatmenttypeFields := schema.RightToTreatmentType{}.Fields()
	_ = righttotreatmenttypeFields
	// righttotreatmenttypeDescTypeName is the schema descriptor for TypeName field.
	righttotreatmenttypeDescTypeName := righttotreatmenttypeFields[0].Descriptor()
	// righttotreatmenttype.TypeNameValidator is a validator for the "TypeName" field. It is called by the builders before save.
	righttotreatmenttype.TypeNameValidator = righttotreatmenttypeDescTypeName.Validators[0].(func(string) error)
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescRoomName is the schema descriptor for roomName field.
	roomDescRoomName := roomFields[0].Descriptor()
	// room.RoomNameValidator is a validator for the "roomName" field. It is called by the builders before save.
	room.RoomNameValidator = roomDescRoomName.Validators[0].(func(string) error)
	treatmenttypeFields := schema.TreatmentType{}.Fields()
	_ = treatmenttypeFields
	// treatmenttypeDescType is the schema descriptor for Type field.
	treatmenttypeDescType := treatmenttypeFields[0].Descriptor()
	// treatmenttype.TypeValidator is a validator for the "Type" field. It is called by the builders before save.
	treatmenttype.TypeValidator = treatmenttypeDescType.Validators[0].(func(string) error)
	triageresultFields := schema.TriageResult{}.Fields()
	_ = triageresultFields
	// triageresultDescSymptom is the schema descriptor for symptom field.
	triageresultDescSymptom := triageresultFields[0].Descriptor()
	// triageresult.SymptomValidator is a validator for the "symptom" field. It is called by the builders before save.
	triageresult.SymptomValidator = triageresultDescSymptom.Validators[0].(func(string) error)
	// triageresultDescTriageDate is the schema descriptor for triageDate field.
	triageresultDescTriageDate := triageresultFields[1].Descriptor()
	// triageresult.DefaultTriageDate holds the default value on creation for the triageDate field.
	triageresult.DefaultTriageDate = triageresultDescTriageDate.Default.(func() time.Time)
	urgencylevelFields := schema.UrgencyLevel{}.Fields()
	_ = urgencylevelFields
	// urgencylevelDescUrgencyName is the schema descriptor for urgencyName field.
	urgencylevelDescUrgencyName := urgencylevelFields[0].Descriptor()
	// urgencylevel.UrgencyNameValidator is a validator for the "urgencyName" field. It is called by the builders before save.
	urgencylevel.UrgencyNameValidator = urgencylevelDescUrgencyName.Validators[0].(func(string) error)
}
