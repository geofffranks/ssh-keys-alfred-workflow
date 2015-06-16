# SSH-Keys Workflow for [Alfred 2](http://www.alfredapp.com/)

This [Alfred](http://www.alfredapp.com/) workflow looks up public SSH keys for users
on [github.com](https://github.com/), allowing you to search
for the user intuitively, and copies all of that user's SSH keys
to the clipboard.

To get started, simply [download the workflow](blob/master/SSH-Keys.alfredworkflow), and double-click
to have Alfred install it for you. You should then be able
to use it like this:

![ssh keys demo](ssh-keys-workflow-animation.gif)

## Commands

Search for a users's keys:
- `keys <username|partial username>`

Log in/Save github API credentials for higher query rates + better suggestions: **Not Implemented**
- `keys login`

Log out/destroy locally saved github API credentials: **Not Implemented**
- `keys logout`
