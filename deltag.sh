#!/usr/bin/env bash

#
# Delete all tag form local and remote repository
#

listtag=`git tag`

git tag -d $listtag

for i in $listtag; do
  git push -d origin $i;
done

	


