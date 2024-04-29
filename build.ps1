param(
    [switch] $Idl,
    [switch] $LintIdl,
    [switch] $Run,
    [switch] $Clean
)

function Lint-Idl {
    Push-Location "$pwd\idl"
    docker run --volume "$(pwd):/workspace" --workdir /workspace bufbuild/buf lint
    Pop-Location
}

function Generate-Idl {
    Push-Location "$pwd\idl"
    docker run --volume "$(pwd):/workspace" --workdir /workspace bufbuild/buf generate proto
    Pop-Location
}

function Run {
    docker-compose up --force-recreate --build
}

function Remove {
    docker-compose down --rmi local
}


if ($Idl) {
    Generate-Idl
    return
}

if ($LintIdl) {
    Lint-Idl
    return
}

if ($Run) {
    Run
    return
}

if ($Clean) {
    Remove
    return
}