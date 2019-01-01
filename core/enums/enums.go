package enums

type (
	//RoleType enum type
	RoleType int16
	//Status enum type
	Status int16
	//Access enum type
	Access int16
	//Severity enum type
	Severity int16
	//Approval enum type
	Approval int16
)

//enum type RoleType
const (
	BoardType       RoleType = 1
	GroupType       RoleType = 2
	SchoolType      RoleType = 3
	SubOrdinateType RoleType = 4
	TeacherType     RoleType = 5
)

//enum type Status
const (
	Active  Status = 1
	Dorment Status = 0
)

//enum tye Access
const (
	None     Access = 0
	ReadOnly Access = 1
	All      Access = 2
)

//enum type Severity
const (
	Easy     Severity = 1
	Moderate Severity = 2
	Hard     Severity = 3
)

//enum type Approval
const (
	Approved Approval = 1
	Waiting  Approval = 2
	Rejected Approval = 3
)
