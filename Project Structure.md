Project Structure

applaunch
	- main.go

Repo
	- userHandler.go
		- addUserHandler().   // common sign-in/sign-up mappings
		- addGroupHandler()
		- addSchoolHandler()
		- addSubHandler()
		- addTeacherHandler()
		- addBoardHandler()
service
	- userService.go
	- groupService.go
	- schoolService.go
	- subService.go
	- teacherService.go
	- boardService.go

repo
	- dbrepo.go
		- insert()
        - fetch()
        - delete()
        - filter() // should contain generic args such as the query which will be formed in the service

model
	- userModel.go
	- roleModel.go
	- groupModel.go
	- locationModel.go
	- addressModel.go
	- contactModel.go
	- schoolModel.go
	- classModel.go
	- subjectModel.go
	- chapterModel.go
	- questionModel.go
	- tagModel.go

common
	-utility.go. // commonly used utility methods used across all files should be here

enums
	- QUES_APPROVAL_STATUS
	- STATUS
	- ACCESS
	- ROLE
	- SEVERITY




