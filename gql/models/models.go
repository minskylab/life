package models

import (
	"io"
	"time"
)

type ID string
type DateTime time.Time
type Time time.Time
type Map map[string]interface{}
type Upload io.Reader

// JobProfileExtraRequeriments is kind OBJECT
type JobProfileExtraRequeriments struct {
	Id string
}

func (jobProfileExtraRequeriments *JobProfileExtraRequeriments) Class() (JobClass, error) {
	panic("unimplemented builder")
}

func (jobProfileExtraRequeriments *JobProfileExtraRequeriments) TravelAvailability() (TravelAvailability, error) {
	panic("unimplemented builder")
}

func (jobProfileExtraRequeriments *JobProfileExtraRequeriments) LicenceRequeriment() (LicenceRequirement, error) {
	panic("unimplemented builder")
}

func (jobProfileExtraRequeriments *JobProfileExtraRequeriments) ContractKind() (ContractKind, error) {
	panic("unimplemented builder")
}

func (jobProfileExtraRequeriments *JobProfileExtraRequeriments) WorkKind() (WorkKind, error) {
	panic("unimplemented builder")
}

type JobFunctionFrequency string

const JOB_FUNCTION_FREQUENCY_DIARY JobFunctionFrequency = "DIARY"
const JOB_FUNCTION_FREQUENCY_WEEKLY JobFunctionFrequency = "WEEKLY"
const JOB_FUNCTION_FREQUENCY_MONTHLY JobFunctionFrequency = "MONTHLY"

// DepartmentCreateOrUse is kind INPUT_OBJECT
type DepartmentCreateOrUse struct {
	FromID *string
	create *NewDepartment
}

// MissionTool is kind OBJECT
type MissionTool struct {
	Id        string
	Protected bool
	Value     string
}

func (missionTool *MissionTool) Creator() (User, error) {
	panic("unimplemented builder")
}

func (missionTool *MissionTool) Organization() (Organization, error) {
	panic("unimplemented builder")
}

func (missionTool *MissionTool) Mission() (JobMission, error) {
	panic("unimplemented builder")
}

// JobCampus is kind OBJECT
type JobCampus struct {
	Id        string
	Protected bool
	Value     string
}

func (jobCampus *JobCampus) Creator() (User, error) {
	panic("unimplemented builder")
}

func (jobCampus *JobCampus) Organization() (Organization, error) {
	panic("unimplemented builder")
}

type JobSortParam string

const JOB_SORT_PARAM_ID JobSortParam = "ID"
const JOB_SORT_PARAM_NAME JobSortParam = "NAME"
const JOB_SORT_PARAM_LAST_UPDATE JobSortParam = "LAST_UPDATE"

// JobPayloadInput is kind INPUT_OBJECT
type JobPayloadInput struct {
	Code              string
	Name              string
	Sites             *int
	department        DepartmentCreateOrUse
	occupationalGroup *OccupationalGroupCreateOrUse
	family            *FamilyCreateOrUse
	campus            *CampusCreateOrUse
}

// NewFundamentalResponse is kind OBJECT
type NewFundamentalResponse struct {
	Id    string
	Value *string
}

func (newFundamentalResponse *NewFundamentalResponse) FundamentalKind() (FundamentalKind, error) {
	panic("unimplemented builder")
}

func (newFundamentalResponse *NewFundamentalResponse) Metadata() (*Map, error) {
	panic("unimplemented builder")
}

// Worker is kind OBJECT
type Worker struct {
	Id           string
	Code         string
	Email        string
	Name         string
	Lastname     string
	Nif          string
	Gender       string
	Phone        string
	AnotherPhone string
	Direction    string
	Country      string
}

func (worker *Worker) Creator() (User, error) {
	panic("unimplemented builder")
}

func (worker *Worker) Supervisor() (Job, error) {
	panic("unimplemented builder")
}

func (worker *Worker) Job() (Job, error) {
	panic("unimplemented builder")
}

func (worker *Worker) BiologicalSex() (BiologicalSex, error) {
	panic("unimplemented builder")
}

func (worker *Worker) Birthdate() (Time, error) {
	panic("unimplemented builder")
}

func (worker *Worker) IncorporationDate() (Time, error) {
	panic("unimplemented builder")
}

func (worker *Worker) Metadata() (Map, error) {
	panic("unimplemented builder")
}

// Sector is kind OBJECT
type Sector struct {
	Id        string
	Protected bool
	Value     string
}

func (sector *Sector) Creator() (User, error) {
	panic("unimplemented builder")
}

func (sector *Sector) Organization() (Organization, error) {
	panic("unimplemented builder")
}

// LicenceRequirement is kind OBJECT
type LicenceRequirement struct {
	Id        string
	Protected bool
	Value     string
}

func (licenceRequirement *LicenceRequirement) Creator() (User, error) {
	panic("unimplemented builder")
}

func (licenceRequirement *LicenceRequirement) Organization() (Organization, error) {
	panic("unimplemented builder")
}

type JobFunctionErrorConsequence string

const JOB_FUNCTION_ERROR_CONSEQUENCE_HIGH JobFunctionErrorConsequence = "HIGH"
const JOB_FUNCTION_ERROR_CONSEQUENCE_MEDIUM JobFunctionErrorConsequence = "MEDIUM"
const JOB_FUNCTION_ERROR_CONSEQUENCE_LOW JobFunctionErrorConsequence = "LOW"

// Role is kind OBJECT
type Role struct {
	Id   string
	Name string
}

func (role *Role) Kind() (RoleRealm, error) {
	panic("unimplemented builder")
}

func (role *Role) Users() ([]User, error) {
	panic("unimplemented builder")
}

func (role *Role) Organization() (Organization, error) {
	panic("unimplemented builder")
}

func (role *Role) Capabilities() ([]Access, error) {
	panic("unimplemented builder")
}

// FileUpload is kind INPUT_OBJECT
type FileUpload struct {
	file Upload
}

// ProcessLevel is kind OBJECT
type ProcessLevel struct {
	Id        string
	Protected bool
	Value     string
}

func (processLevel *ProcessLevel) Creator() (User, error) {
	panic("unimplemented builder")
}

func (processLevel *ProcessLevel) Organization() (Organization, error) {
	panic("unimplemented builder")
}

// JobClass is kind OBJECT
type JobClass struct {
	Id        string
	Protected bool
	Value     string
}

func (jobClass *JobClass) Creator() (User, error) {
	panic("unimplemented builder")
}

func (jobClass *JobClass) Organization() (Organization, error) {
	panic("unimplemented builder")
}

// JobFunctionAssessment is kind OBJECT
type JobFunctionAssessment struct {
	Id string
}

func (jobFunctionAssessment *JobFunctionAssessment) Frequency() (JobFunctionFrequency, error) {
	panic("unimplemented builder")
}

func (jobFunctionAssessment *JobFunctionAssessment) ErrorConsequence() (JobFunctionErrorConsequence, error) {
	panic("unimplemented builder")
}

func (jobFunctionAssessment *JobFunctionAssessment) Complexity() (JobFunctionComplexity, error) {
	panic("unimplemented builder")
}

func (jobFunctionAssessment *JobFunctionAssessment) Kpis() ([]JobFunctionKPI, error) {
	panic("unimplemented builder")
}

// LoginCredentials is kind INPUT_OBJECT
type LoginCredentials struct {
	Username string
	Password string
}

// Competency is kind OBJECT
type Competency struct {
	Id        string
	Protected bool
	Value     string
}

func (competency *Competency) Creator() (User, error) {
	panic("unimplemented builder")
}

func (competency *Competency) Organization() (Organization, error) {
	panic("unimplemented builder")
}

func (competency *Competency) Metadata() (Map, error) {
	panic("unimplemented builder")
}

// DownloadableFile is kind OBJECT
type DownloadableFile struct {
	BlobURL     string
	ContentType string
	Hash        *string
}

func (downloadableFile *DownloadableFile) CreatedAt() (Time, error) {
	panic("unimplemented builder")
}

func (downloadableFile *DownloadableFile) ExpireAt() (Time, error) {
	panic("unimplemented builder")
}

// LogEntry is kind OBJECT
type LogEntry struct {
	Id string
}

// NewDepartment is kind INPUT_OBJECT
type NewDepartment struct {
	Code   string
	Name   string
	Parent *string
}

// JobOccupationalGroup is kind OBJECT
type JobOccupationalGroup struct {
	Id        string
	Protected bool
	Value     string
}

func (jobOccupationalGroup *JobOccupationalGroup) Creator() (User, error) {
	panic("unimplemented builder")
}

func (jobOccupationalGroup *JobOccupationalGroup) Organization() (Organization, error) {
	panic("unimplemented builder")
}

// StructureSummary is kind OBJECT
type StructureSummary struct {
	TotalDepartments int
	TotalJobs        int
	TotalSites       int
}

// LanguageProficiency is kind OBJECT
type LanguageProficiency struct {
	Id        string
	Protected bool
	Value     string
}

func (languageProficiency *LanguageProficiency) Creator() (User, error) {
	panic("unimplemented builder")
}

func (languageProficiency *LanguageProficiency) Organization() (Organization, error) {
	panic("unimplemented builder")
}

// JobProfile is kind OBJECT
type JobProfile struct {
	Id string
}

func (jobProfile *JobProfile) Education() (JobProfileEducation, error) {
	panic("unimplemented builder")
}

func (jobProfile *JobProfile) ComplementaryEducation() (JobProfileComplementaryEducation, error) {
	panic("unimplemented builder")
}

func (jobProfile *JobProfile) Experience() (JobProfileExperience, error) {
	panic("unimplemented builder")
}

func (jobProfile *JobProfile) Risk() (RiskKind, error) {
	panic("unimplemented builder")
}

func (jobProfile *JobProfile) Extra() (JobProfileExtraRequeriments, error) {
	panic("unimplemented builder")
}

func (jobProfile *JobProfile) Competencies() (JobProfileCompetency, error) {
	panic("unimplemented builder")
}

// ValorationFactor is kind OBJECT
type ValorationFactor struct {
	Id          string
	Creator     string
	Name        string
	Description string
}

// User is kind OBJECT
type User struct {
	Id    string
	Name  string
	Email string
}

func (user *User) Roles() ([]Role, error) {
	panic("unimplemented builder")
}

type FundamentalKind string

const FUNDAMENTAL_KIND_MI_SSION_VERB FundamentalKind = "MiSSION_VERB"
const FUNDAMENTAL_KIND_MISSION_TOOL FundamentalKind = "MISSION_TOOL"
const FUNDAMENTAL_KIND_RESPONSIBILITY_LEVEL FundamentalKind = "RESPONSIBILITY_LEVEL"
const FUNDAMENTAL_KIND_PROCESS_LEVEL FundamentalKind = "PROCESS_LEVEL"
const FUNDAMENTAL_KIND_INSTRUCTION_GRADUATE FundamentalKind = "INSTRUCTION_GRADUATE"
const FUNDAMENTAL_KIND_INSTRUCTION_SPECIALITY FundamentalKind = "INSTRUCTION_SPECIALITY"
const FUNDAMENTAL_KIND_MASTER_AND_DOCTORATE FundamentalKind = "MASTER_AND_DOCTORATE"
const FUNDAMENTAL_KIND_COURSE_AND_GRADUATE FundamentalKind = "COURSE_AND_GRADUATE"
const FUNDAMENTAL_KIND_LANGUAGE_PROFICIENCY FundamentalKind = "LANGUAGE_PROFICIENCY"
const FUNDAMENTAL_KIND_SOFTWARE FundamentalKind = "SOFTWARE"
const FUNDAMENTAL_KIND_EXPERIENCE_TIME FundamentalKind = "EXPERIENCE_TIME"
const FUNDAMENTAL_KIND_AREA FundamentalKind = "AREA"
const FUNDAMENTAL_KIND_SECTOR FundamentalKind = "SECTOR"
const FUNDAMENTAL_KIND_JOB_CLASS FundamentalKind = "JOB_CLASS"
const FUNDAMENTAL_KIND_TRAVEL_AVAILABILITY FundamentalKind = "TRAVEL_AVAILABILITY"
const FUNDAMENTAL_KIND_LICENCE_REQUIREMENT FundamentalKind = "LICENCE_REQUIREMENT"
const FUNDAMENTAL_KIND_CONTRACT_KIND FundamentalKind = "CONTRACT_KIND"
const FUNDAMENTAL_KIND_WORK_KIND FundamentalKind = "WORK_KIND"
const FUNDAMENTAL_KIND_RISK_KIND FundamentalKind = "RISK_KIND"
const FUNDAMENTAL_KIND_JOB_FAMILY FundamentalKind = "JOB_FAMILY"
const FUNDAMENTAL_KIND_JOB_OCCUPATIONAL_GROUP FundamentalKind = "JOB_OCCUPATIONAL_GROUP"
const FUNDAMENTAL_KIND_JOB_CAMPUS FundamentalKind = "JOB_CAMPUS"
const FUNDAMENTAL_KIND_GUIDE FundamentalKind = "GUIDE"

// JobFunction is kind OBJECT
type JobFunction struct {
	Id        string
	Actions   []string
	Objective string
}

func (jobFunction *JobFunction) Assessment() (JobFunctionAssessment, error) {
	panic("unimplemented builder")
}

// JobFunctionKPI is kind OBJECT
type JobFunctionKPI struct {
	Id    string
	Name  string
	Value string
}

// OccupationalGroupCreateOrUse is kind INPUT_OBJECT
type OccupationalGroupCreateOrUse struct {
	FromID   *string
	NewValue *string
}

// OrganizationBrand is kind OBJECT
type OrganizationBrand struct {
	Id             string
	Photo          string
	BrandName      string
	PrimaryColor   *string
	SecondaryColor *string
}

// OrganizationAsset is kind OBJECT
type OrganizationAsset struct {
	Id      string
	BlobURL string
	Name    string
}

func (organizationAsset *OrganizationAsset) Parent() (Organization, error) {
	panic("unimplemented builder")
}

// ResponsibilityLevel is kind OBJECT
type ResponsibilityLevel struct {
	Id        string
	Protected bool
	Value     string
}

func (responsibilityLevel *ResponsibilityLevel) Creator() (User, error) {
	panic("unimplemented builder")
}

func (responsibilityLevel *ResponsibilityLevel) Organization() (Organization, error) {
	panic("unimplemented builder")
}

func (responsibilityLevel *ResponsibilityLevel) Function() (JobFunction, error) {
	panic("unimplemented builder")
}

// JobFamily is kind OBJECT
type JobFamily struct {
	Id        string
	Protected bool
	Value     string
}

func (jobFamily *JobFamily) Creator() (User, error) {
	panic("unimplemented builder")
}

func (jobFamily *JobFamily) Organization() (Organization, error) {
	panic("unimplemented builder")
}

// Job is kind OBJECT
type Job struct {
	Id       string
	Name     string
	Code     string
	Sites    int
	Category int
}

func (job *Job) CreatedAt() (Time, error) {
	panic("unimplemented builder")
}

func (job *Job) UpdatedAt() (Time, error) {
	panic("unimplemented builder")
}

func (job *Job) Metadata() (Map, error) {
	panic("unimplemented builder")
}

func (job *Job) PlacedIn() (Department, error) {
	panic("unimplemented builder")
}

func (job *Job) OccupationalGroup() (*JobOccupationalGroup, error) {
	panic("unimplemented builder")
}

func (job *Job) Family() (*JobFamily, error) {
	panic("unimplemented builder")
}

func (job *Job) Campus() (*JobCampus, error) {
	panic("unimplemented builder")
}

func (job *Job) Mission() (JobMission, error) {
	panic("unimplemented builder")
}

func (job *Job) Functions() ([]JobFunction, error) {
	panic("unimplemented builder")
}

func (job *Job) Profile() (JobProfile, error) {
	panic("unimplemented builder")
}

func (job *Job) Workers() ([]Worker, error) {
	panic("unimplemented builder")
}

// CampusCreateOrUse is kind INPUT_OBJECT
type CampusCreateOrUse struct {
	FromID   *string
	NewValue *string
}

// MissionVerb is kind OBJECT
type MissionVerb struct {
	Id        string
	Protected bool
	Value     string
}

func (missionVerb *MissionVerb) Creator() (User, error) {
	panic("unimplemented builder")
}

func (missionVerb *MissionVerb) Organization() (Organization, error) {
	panic("unimplemented builder")
}

func (missionVerb *MissionVerb) Mission() (JobMission, error) {
	panic("unimplemented builder")
}

// Area is kind OBJECT
type Area struct {
	Id        string
	Protected bool
	Value     string
}

func (area *Area) Creator() (User, error) {
	panic("unimplemented builder")
}

func (area *Area) Organization() (Organization, error) {
	panic("unimplemented builder")
}

// TravelAvailability is kind OBJECT
type TravelAvailability struct {
	Id        string
	Protected bool
	Value     string
}

func (travelAvailability *TravelAvailability) Creator() (User, error) {
	panic("unimplemented builder")
}

func (travelAvailability *TravelAvailability) Organization() (Organization, error) {
	panic("unimplemented builder")
}

// GenericFundamental is kind OBJECT
type GenericFundamental struct {
	Id    string
	Value *string
}

func (genericFundamental *GenericFundamental) Kind() (FundamentalKind, error) {
	panic("unimplemented builder")
}

func (genericFundamental *GenericFundamental) ValueMap() (*Map, error) {
	panic("unimplemented builder")
}

// RiskKind is kind OBJECT
type RiskKind struct {
	Id        string
	Protected bool
	Value     string
}

func (riskKind *RiskKind) Creator() (User, error) {
	panic("unimplemented builder")
}

func (riskKind *RiskKind) Organization() (Organization, error) {
	panic("unimplemented builder")
}

func (riskKind *RiskKind) Metadata() (Map, error) {
	panic("unimplemented builder")
}

// Guide is kind OBJECT
type Guide struct {
	Id        string
	Protected bool
	Value     string
}

func (guide *Guide) Creator() (User, error) {
	panic("unimplemented builder")
}

func (guide *Guide) Organization() (Organization, error) {
	panic("unimplemented builder")
}

// JobDescription is kind OBJECT
type JobDescription struct {
	Id string
}

// JobMission is kind OBJECT
type JobMission struct {
	Id    string
	Goal  string
	Scope string
}

func (jobMission *JobMission) ActionAndScope() ([]MissionVerb, error) {
	panic("unimplemented builder")
}

func (jobMission *JobMission) Guides() ([]MissionTool, error) {
	panic("unimplemented builder")
}

// LoginResponse is kind OBJECT
type LoginResponse struct {
	Token string
}

func (loginResponse *LoginResponse) User() (User, error) {
	panic("unimplemented builder")
}

// Department is kind OBJECT
type Department struct {
	Id                    string
	Code                  string
	Name                  string
	ExpectedNumberOfSeats int
}

func (department *Department) Parent() (*Department, error) {
	panic("unimplemented builder")
}

func (department *Department) Jobs() ([]Job, error) {
	panic("unimplemented builder")
}

func (department *Department) Metadata() (Map, error) {
	panic("unimplemented builder")
}

// Error is kind OBJECT
type Error struct {
	Id      string
	Message string
	Source  *string
}

func (error *Error) At() (*Time, error) {
	panic("unimplemented builder")
}

// InstructionSpeciality is kind OBJECT
type InstructionSpeciality struct {
	Id        string
	Protected bool
	Value     string
}

func (instructionSpeciality *InstructionSpeciality) Creator() (User, error) {
	panic("unimplemented builder")
}

func (instructionSpeciality *InstructionSpeciality) Organization() (Organization, error) {
	panic("unimplemented builder")
}

// JobProfileEducation is kind OBJECT
type JobProfileEducation struct {
	Id string
}

func (jobProfileEducation *JobProfileEducation) InstructionLevel() (InstructionGraduate, error) {
	panic("unimplemented builder")
}

func (jobProfileEducation *JobProfileEducation) InstructionSpecialities() ([]InstructionSpeciality, error) {
	panic("unimplemented builder")
}

// JobProfileComplementaryEducation is kind OBJECT
type JobProfileComplementaryEducation struct {
	Id string
}

func (jobProfileComplementaryEducation *JobProfileComplementaryEducation) PostGraduate() (MasterAndDoctorate, error) {
	panic("unimplemented builder")
}

func (jobProfileComplementaryEducation *JobProfileComplementaryEducation) Courses() ([]CourseAndGraduate, error) {
	panic("unimplemented builder")
}

func (jobProfileComplementaryEducation *JobProfileComplementaryEducation) Languages() ([]LanguageProficiency, error) {
	panic("unimplemented builder")
}

func (jobProfileComplementaryEducation *JobProfileComplementaryEducation) Softwares() ([]Software, error) {
	panic("unimplemented builder")
}

// StructureTree is kind OBJECT
type StructureTree struct {
	IsRoot *bool
}

func (structureTree *StructureTree) Value() (Department, error) {
	panic("unimplemented builder")
}

func (structureTree *StructureTree) Children() ([]Department, error) {
	panic("unimplemented builder")
}

// ExperienceTime is kind OBJECT
type ExperienceTime struct {
	Id        string
	Protected bool
	Value     string
}

func (experienceTime *ExperienceTime) Creator() (User, error) {
	panic("unimplemented builder")
}

func (experienceTime *ExperienceTime) Organization() (Organization, error) {
	panic("unimplemented builder")
}

// JobProfileExperience is kind OBJECT
type JobProfileExperience struct {
	Id     string
	Others string
}

func (jobProfileExperience *JobProfileExperience) ExperienceTime() (ExperienceTime, error) {
	panic("unimplemented builder")
}

func (jobProfileExperience *JobProfileExperience) Areas() ([]Area, error) {
	panic("unimplemented builder")
}

func (jobProfileExperience *JobProfileExperience) Sectors() ([]Sector, error) {
	panic("unimplemented builder")
}

// FamilyCreateOrUse is kind INPUT_OBJECT
type FamilyCreateOrUse struct {
	FromID   *string
	NewValue *string
}

// Organization is kind OBJECT
type Organization struct {
	Id   string
	Name string
}

func (organization *Organization) Admin() (User, error) {
	panic("unimplemented builder")
}

func (organization *Organization) Brand() (OrganizationBrand, error) {
	panic("unimplemented builder")
}

func (organization *Organization) Fundamentals() (OrganizationFundamentals, error) {
	panic("unimplemented builder")
}

func (organization *Organization) Jobs() ([]Job, error) {
	panic("unimplemented builder")
}

func (organization *Organization) Workers() ([]Worker, error) {
	panic("unimplemented builder")
}

func (organization *Organization) Files() ([]OrganizationAsset, error) {
	panic("unimplemented builder")
}

// SalarialStudy is kind OBJECT
type SalarialStudy struct {
	Id string
}

type RoleRealm string

const ROLE_REALM_ADMIN RoleRealm = "ADMIN"
const ROLE_REALM_MODERATOR RoleRealm = "MODERATOR"
const ROLE_REALM_AMBASSADOR RoleRealm = "AMBASSADOR"
const ROLE_REALM_ORGANIZATION RoleRealm = "ORGANIZATION"
const ROLE_REALM_NORMAL RoleRealm = "NORMAL"

// InstructionGraduate is kind OBJECT
type InstructionGraduate struct {
	Id        string
	Protected bool
	Value     string
}

func (instructionGraduate *InstructionGraduate) Creator() (User, error) {
	panic("unimplemented builder")
}

func (instructionGraduate *InstructionGraduate) Organization() (Organization, error) {
	panic("unimplemented builder")
}

// CourseAndGraduate is kind OBJECT
type CourseAndGraduate struct {
	Id        string
	Protected bool
	Value     string
}

func (courseAndGraduate *CourseAndGraduate) Creator() (User, error) {
	panic("unimplemented builder")
}

func (courseAndGraduate *CourseAndGraduate) Organization() (Organization, error) {
	panic("unimplemented builder")
}

// JobProfileCompetency is kind OBJECT
type JobProfileCompetency struct{}

func (jobProfileCompetency *JobProfileCompetency) Organizationals() ([]Competency, error) {
	panic("unimplemented builder")
}

func (jobProfileCompetency *JobProfileCompetency) Functionals() ([]Competency, error) {
	panic("unimplemented builder")
}

// WorkKind is kind OBJECT
type WorkKind struct {
	Id        string
	Protected bool
	Value     string
}

func (workKind *WorkKind) Creator() (User, error) {
	panic("unimplemented builder")
}

func (workKind *WorkKind) Organization() (Organization, error) {
	panic("unimplemented builder")
}

// Software is kind OBJECT
type Software struct {
	Id        string
	Protected bool
	Value     string
}

func (software *Software) Creator() (User, error) {
	panic("unimplemented builder")
}

func (software *Software) Organization() (Organization, error) {
	panic("unimplemented builder")
}

// ContractKind is kind OBJECT
type ContractKind struct {
	Id        string
	Protected bool
	Value     string
}

func (contractKind *ContractKind) Creator() (User, error) {
	panic("unimplemented builder")
}

func (contractKind *ContractKind) Organization() (Organization, error) {
	panic("unimplemented builder")
}

type JobFunctionComplexity string

const JOB_FUNCTION_COMPLEXITY_HIGH JobFunctionComplexity = "HIGH"
const JOB_FUNCTION_COMPLEXITY_MEDIUM JobFunctionComplexity = "MEDIUM"
const JOB_FUNCTION_COMPLEXITY_LOW JobFunctionComplexity = "LOW"

// JobPayloadUpdate is kind INPUT_OBJECT
type JobPayloadUpdate struct {
	Code              *string
	Name              *string
	Sites             *int
	department        *DepartmentCreateOrUse
	occupationalGroup *OccupationalGroupCreateOrUse
	family            *FamilyCreateOrUse
	campus            *CampusCreateOrUse
}

// MasterAndDoctorate is kind OBJECT
type MasterAndDoctorate struct {
	Id        string
	Protected bool
	Value     string
}

func (masterAndDoctorate *MasterAndDoctorate) Creator() (User, error) {
	panic("unimplemented builder")
}

func (masterAndDoctorate *MasterAndDoctorate) Organization() (Organization, error) {
	panic("unimplemented builder")
}

// OrganizationFundamentals is kind OBJECT
type OrganizationFundamentals struct{}

func (organizationFundamentals *OrganizationFundamentals) MissionVerbs() ([]MissionVerb, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) MissionTools() ([]MissionTool, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) ResponsibilityLevels() ([]ResponsibilityLevel, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) ProcessLevels() ([]ProcessLevel, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) InstructionGraduates() ([]InstructionGraduate, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) InstructionSpecialities() ([]InstructionSpeciality, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) MasterAndDoctorates() ([]MasterAndDoctorate, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) CourseAndGraduates() ([]CourseAndGraduate, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) LanguageProficiencies() ([]LanguageProficiency, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) Softwares() ([]Software, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) ExperienceTimes() ([]ExperienceTime, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) Areas() ([]Area, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) Sectors() ([]Sector, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) JobClasss() ([]JobClass, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) TravelAvailabilities() ([]TravelAvailability, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) LicenceRequirements() ([]LicenceRequirement, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) ContractKinds() ([]ContractKind, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) WorkKinds() ([]WorkKind, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) RiskKinds() ([]RiskKind, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) JobFamilies() ([]JobFamily, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) JobOccupationalGroups() ([]JobOccupationalGroup, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) JobCampus() ([]JobCampus, error) {
	panic("unimplemented builder")
}

func (organizationFundamentals *OrganizationFundamentals) Guides() ([]Guide, error) {
	panic("unimplemented builder")
}

type BiologicalSex string

const BIOLOGICAL_SEX_MALE BiologicalSex = "MALE"
const BIOLOGICAL_SEX_FEMALE BiologicalSex = "FEMALE"
const BIOLOGICAL_SEX_INTERSEXUAL BiologicalSex = "INTERSEXUAL"
const BIOLOGICAL_SEX_UNKNOWN BiologicalSex = "UNKNOWN"

// Access is kind OBJECT
type Access struct {
	Id       string
	Resource string
	Write    bool
	Read     bool
	Delete   bool
}

func (access *Access) Role() (Role, error) {
	panic("unimplemented builder")
}

// OrganizationAndJobsUpdate is kind OBJECT
type OrganizationAndJobsUpdate struct {
	UpdateID     string
	TotalNewRows int
	Preview      string
}

// PromotedOrganizationAndJobsUpdate is kind OBJECT
type PromotedOrganizationAndJobsUpdate struct {
	UpdateID    string
	AddedRows   int
	UpdatedRows int
	DeletedRows int
}
