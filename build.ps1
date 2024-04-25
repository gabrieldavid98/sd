param(
    [switch] $Idl,
    [switch] $LintIdl
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


if ($Idl) {
    Generate-Idl
    return
}

if ($LintIdl) {
    Lint-Idl
    return
}