ACTORS:
	Teachers -> regular teachers of a school
	Sub-admin -> managing teachers for a specific class range
	School admin -> managing sub admins
	Group admin -> managing school admins
	Board Admin -> read access to all data of schools under 1 group

ACTIONS / OPERATIONS / USE CASES:
	Teachers:
		login account
		update account profile
		add/update/discard questions
		view questions (segregated into 2 tabs - self , all)
		generate question papers
		download self-submitted questions as excel

SUB-ADMIN:
		approves/rejects questions based on assigned class range and subject range
		view teacher under his mapping
	
SCHOOL ADMIN:
		update self-account
		update school profile
		creates sub-admins/teacher account based on class & subject range
		update sub-admin/teacher account
		activate/deactivate subadmin/teacher account  on joining or leaving school
		view account of all subadmins/teacher
		transfer ownership from 1 sub admin/teacher to another
		view all teachers of school ( read only )

GROUP ADMIN:
		update self-account
		add/update/remove schools in a group (initial setup by us)
		create school admin account (initial setup by us)
		add/update curriculum (e.g. subjects,classes,chapters)  (initial setup by us)
		view all schools, school admins (all access)

BOARD ADMIN:
		view all data under this group (read only)
		
OUR ROLE:
		create a group
		create a group admin
		add schools in group
		create school admins
		map school admins under group admin
		help creating meta data such as classes,subjects,chapters
		
		monthly check group stats - for pricing and other purposes (1st n 15th of Every month)
		get notified as soon as a school is added/deleted from a group
