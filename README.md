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
```
    list of available app commands

    lint            - lint app.
    test            - test app.
    test            - test app.
    vet             - vet app.
    build           - build app.
    serve           - serve app.
    benchmark       - benchmark app.

    docker commands
    up-env          - start dev container
    up-cli          - start cli container
    exec-env        - exec into dev container
    exec-cli        - exec into cli container
    purge-env       - purge dev container
    purge-cli       - purge cli container
```
