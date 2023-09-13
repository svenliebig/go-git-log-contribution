# go-git-log-contribution

Wonder how long you have been commiting to projects? Not anymore!

## install

```sh
git clone git@github.com:svenliebig/go-git-log-contribution.git
cd go-git-log-contribution
go build
go install .
```

## usage

add `<YOUR GIT NAME>` please and execute:

```sh
git log --encoding=utf-8 --full-history --reverse "--format=format:%cs;%an;%ae" --author "<YOUR GIT NAME>" | git-log-contribution
```

prints out the duration that you have been consecutively commiting to this project (max gap 45 days):

```sh
16.03.2021 - 15.08.2023
```

🚀
