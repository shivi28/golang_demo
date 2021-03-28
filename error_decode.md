1. If in http call, encounter with `unsupported protocol scheme "" ` error then add `http://` in the URL


2. Want to see stash file changes without applying it. 
```
git stash list
git stash show -p // to see recent stash 
```

3. If while applying stash, faced some error like
```
.idea/sonarIssues.xml already exists, no checkout
.idea/workspace.xml already exists, no checkout
error: could not restore untracked files from stash
```
Solution:
 ```
 git checkout stash -- .
 ```
