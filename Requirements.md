ACTORS:
	Teachers -> regualar teachers of a school
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
		download self submitted questions as excel

	SUB-ADMIN:
		creates teacher account
		update teacher account
		activate/deactivate teacher account on joining or leaving school
		approves/rejects questions based on assigned class range and subject range
		view teacher under his mapping
		transfer ownership from 1 teacher to another

	SCHOOL ADMIN:
		upate self account
		update school profile
		creates sub-admins account based on class & subject range
		update sub-admin account
		acivate/deactivate subadmin account
		view account of all subadmins
		transfer ownership from 1 sub admin to another
		view all teachers of school ( read only )

	GROUP ADMIN:
		update self account
		add/update/remove schools in a group (initial setup by us)
		create school admin account (initial setup by us)
		add/update curriculum (eg. subjects,classes,chapters)  (initial setup by us)
		view all schools,school admins (all access)

	BOARD ADMIN:
		view all data under this group (read only)


	OUR ROLE:
		create a grooup
		create a group admin
		add schools in group
		create school admins
		map school admins under group admin
		help creating meta data such as classes,subjects,chapters
		
		monthly check group stats - for pricing and other purposes (1st n 15th of evry month)
		get notified as soon as a school is addded/deleted from a group








