# PR Status Notifier

Do you find yourself waiting for the little green bulb to turn green so you can merge a PR? Sometimes the tests and associated hooks can take a long time to run and GitHub doesn't send you an email when the mergeable status changes.

Enter `pr_status_notifier`!

`pr_status_notifier` is a simple command line utility that will check a given pull request to see if it can be merged. If it can be merged it will pop up a native notification in your Operating System (currently works for MacOS, Windows and Linux).

Looks something [like this](https://cl.ly/1J2z161G3S1k).

## Usage
```shell
# pr_status_notifier <owner> <repo> <issueNumber>
pr_status_notifier jazzdan pr_status_notifier 1
```

## Installation

### From release
```shell
wget something
untar
```

Then add to path

### From source
Prerequisites:
* Go (tested with 1.9)

```shell
git clone https://www.github.com/jazzdan/pr_status_notifier
cd pr status_notifier
go buold
```
