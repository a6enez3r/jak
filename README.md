# blacjac ![build](https://github.com/abmamo/blacjac/workflows/build/badge.svg?branch=main)
cli blackjack game written in go (mainly to explore cross platform builds in `golang`)

## quickstart
download the binary [for your platform using] `curl`
```
    curl -L  https://github.com/abmamo/BlacJac/raw/main/builds/BlacJac-darwin-amd64 >> BlacJac && chmod +x ./BlacJac
```
and run it
```
    ./blacjac-<whatever platform you want>
```

## develop
usage:
  make <cmd>

cmds:
  help                 show help
  save-local           save changes locally using git
  save-remote          save changes to remote using git
  pull-remote          pull changes from remote
  tag                  create new tag, recreate if it exists
  deps-dev             install deps [dev]
  build                cross platform build
  run                  run package
  test                 test package
  benchmark            benchmark package
  coverage             test coverage
  vet                  vet modules
  lint                 lint package
  format               format package
  scan-duplicate       scan package for duplicate code [dupl]
  scan-errors          scan package for errors [errcheck]
  scan-security        scan package for security issues [gosec]
  build-env            build docker env
  up-env               start docker env
  exec-env             exec. into docker env
  purge-env            remove docker env
  init-env             init env + install common tools
```
