Go must be installed on the device running the CLI

Once installed cd into the repo root directory so that you can run the program.

Run the `main.go` file directly by specifying its path:

```bash
go run src/main.go
```
Testing Info:

The program simply asks for a user role, valid roles include employee, guest, admin. Once 
a valid role is provided it will ask for a valid action, q to quit, or r to change roles.
Valid actions include read, write, or delete. Once both are selected and valid it will tell 
you if you have access to the action based on the role. Since the logic loops until you quit 
or interrupt the program it should ideally be pretty easy to test by testing each role and actions.
