# go-module-template

On the other repositories you have to add this template repository as a remote.

`git remote add template https://github.com/Dana-Team/go-module-template.git`

Then run git fetch to update the changes

`git fetch --all`

Then is possible to merge another branch from the new remote to your current one.

`git merge template/[branch to merge]`

# Run a GitHub actions locally

After clone the template to your local computer you can use `act` command to make actions on your module like tests etc...

Here you can see how to use `act` command: https://github.com/nektos/act
