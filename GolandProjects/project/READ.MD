Documentation of Major Project
i)Add new Task
ii)Assign above task to a new employee
iii)Above task which is created is by mistake assigned a wrong project_id update that project_id as 126
iv)Get all tasks based on month filter client will pass the month i.e. integer say 2 you have to fetch data which are older than say 2 months
v)Get all tasks based on from and to date filters along with project and employee details


* Add new task-
   The given task is to add a new record in task table.
   We have used POST http method to insert a new record into the task table in the given database.
   The endpoint for this API is /createnewtask/

  
* Assign above task to a new employee-
  The given API is to assign the above new task to the new employee.
  We have used POST http method to insert a new record into the employee table in the given database.
  The endpoint for this API is /createnewemp/

* Above task which is created is by mistake assigned a wrong project_id update that project_id as 126
  The given API is to update the project_id as 126 in the previously created new record in task table.
  We have used PUT http method to update the project_id in task table in the given database.
  The endpoint for this API is /updateprojectid/:projectid/

* Get all tasks based on month filter client will pass the month i.e. integer say 2 you have to fetch data which are older than say 2 months
  The given API is to fetch the records if we have entered 2 we will have to fetch the data which are older than 2 months.
  We have used GET http method tO perform the above API.
  The endpoint for this API is /task4/

  
* Get all tasks based on from and to date filters along with project and employee details-
  The given API is to get all the tasks from the table in the database based on the from and to date we pass.
  We have used GET http method tO perform the above API.
  The endpoint for this API is /task5/


  T













