User
	
Role
	
Group

School

Standard/Class

Subject

Chapter

Question

----------------------


Class User:
	Long userID
	String First_Name
	String Last_Name
	String Mobile
	String Email
	String Password
	Group group
	School school
	List<Class> classes
	ROLE role
 	STATUS status
	ACCESS access

Class Group:
	Long groupID
	String groupName
	String groupCode
	String Board
	List<Users> users
	List<School> schools
	List<Class> classes
	List<Question> questionBank
	List<Question> auxQuestionBank

Class Location:
	String latitude
	String longitude

Class Address:
	String street
	String city
	String state
	String pincode
	Location location

Class Contact:
	String[] landline
	String[] mobile
	String fax
	String email
	String websiteURL

Class School:
	Long schoolID
	String schoolName
	String schoolCode
	Address address
	Contact contact
	STATUS status

Class Standard/Class:
	Long classID
	String class
	List<Subject> subjects
	STATUS status

Class Subject:
	Long subjectID
	String subjectCode
	String subjectName
	List<Chapters> chapters
	STATUS status

Class Chapter:
	Long chapterID
	String chapterSeq
	String chapterName
	STATUS status

Class Question:
	Long quesID
	String description
	Chapter chapter
	Subject subject
	List<Tag> tags
	SEVERITY severity
	Integer marks
	QUES_APPROVAL_STATUS approval_status
	String approval_description
	STATUS status

Class Tag:
	Long tagID
	String descrption

Enum QUES_APPROVAL_STATUS:
	APPROVED , REJECTED , WAITING 

Enum STATUS:
	ACTIVE , DORMENT

Enum ACCESS:
	ALL, NONE, READ-ONLY

Enum ROLE:
	BOARD, GROUP, SCHOOL, SUBORDINATE, TEACHER
	
Enum SEVERITY:
	EASY, MODERATE, HARD



