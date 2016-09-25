# ![](logo.png)
## Build
![](https://travis-ci.org/gulyasm/focus.svg?branch=master)
## Usage
 - `focus start "Coding kernel-bug"`

Add a new item. Clock is ticking from this point.

 - `focus stop`

Stops the currently running element

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

## Integrations
### tmux
To add your current running element to the right side of your tmux status bar, add the following line to your `tmux.conf`.
```
set-option -g status-right '⤇ #(focus now)'
```
## Development
`Vagrantfile` is provided for development. To run it:
```bash
vagrant up && vagrant ssh
```
On the vagrant machine you can build and run Focus.
```bash
go build
./focus start Coding...
./focus start Coding...
./focus start Coding...
./focus list
```
## Contribution
Feel free to submit your PR.
