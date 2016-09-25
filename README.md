# Focus
## Usage
 - `focus add "Coding kernel-bug"`

Add a new item. Clock is ticking from this point.

 - `focus now`

Print what are you working now.

 - `focus today`

Your timesheet for today.

 - `focus yesterday`

Your timesheet for yesterday.

 - `focus list`

List all elements

## Install
**Currently not available**
[Download](http://github.com/gulyasm/focus) and install the binaries.

If you have go installed you can install with go install.
```bash
go install github.com/gulyasm/focus
```
## Development
`Vagrantfile` is provided for development. To run it:
```bash
vagrant up && vagrant ssh
```
On the vagrant machine you can build and run Focus.
```bash
go build
./focus add Coding...
./focus add Coding...
./focus add Coding...
./focus list
```
## Contribution
Feel free to submit your PR.
