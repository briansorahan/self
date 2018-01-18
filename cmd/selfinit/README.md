# cligen

Generate Go programs that run in tiny docker images.

## Install

```shell
go get -u git.drillinginfo.com/brian-sorahan/cligen
```

## Usage

```shell
mkdir $GOPATH/src/git.drillinginfo.com/my-org/new-repo
cd $GOPATH/src/git.drillinginfo.com/my-org/new-repo
cligen -name new-repo -org my-org
git init
git remote add origin git@git.drillinginfo.com:my-org/new-repo.git
git add -A
git commit -m 'Initial commit'
git push origin master
```

## Next Steps

### Build

```
make
```

### Test

```
make test
```

### Release

**Pushes to docker hub: use with care!**

```
make push
```
