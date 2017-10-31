#!/bin/sh

currentDir="${PWD##*/}"
#echo $currentDir

if [ -d "vendor" ]; then
    rm -rf "vendor"
fi

#DIRS=`find ./* -name '*.go'  | awk -F'/' '{print $2}' | grep -v *.go | sort -u`
DIRS=`find . -name '*.go' | awk 'BEGIN{FS=OFS="/"}{sub(/.\//,"",$0);print}'|awk 'BEGIN{FS=OFS="/"}{sub(/[^\/ ]*.go/,"",$0);print}' | sort -u`
#DIRS=`find . -type d -print | grep -v *.go `
#echo $DIRS

List=()


for DIR in $DIRS
do
    #echo $DIR
    imports=$(go list -f '{{ join .Imports "\n" }}' bitbucket.org/tekion/$currentDir/$DIR . | grep bitbucket.org/tekion | grep -v bitbucket.org/tekion/$currentDir  )

	#echo $imports
    for i in $imports
    do
        tmp=$( echo $i | awk -F'/' '{print $3}')
        List+=" $tmp"
        #echo $tmp
    done
    #echo  ${DIR}
done
#echo ${List[*]}

GOIMPORTSLIST=$(echo "${List[@]}" | tr ' ' '\n' | sort -u | tr '\n' ' ')

echo ${GOIMPORTSLIST[*]}

#export REPOBASEURL=http://bitbucket.org/tekion
#export GITNAME=.git
#
#mkdir vendor
#mkdir vendor/bitbucket.org
#mkdir vendor/bitbucket.org/tekion
#cd vendor/bitbucket.org/tekion
#
##imports=$(echo $GOIMPORTSLIST | tr "," "\n")
#for i in $GOIMPORTSLIST
#do
#   #echo "$i"
#   git clone -b $1 $REPOBASEURL/$i$GITNAME
#done
#
#cd $GOPATH/src/bitbucket.org/tekion/$currentDir