# Petal fork specific documentation

## Description

The fork was done to unblock using Camel-k on Azure AKS.

## Status

We only need the container image with the appropriate fix, at the latest version currently
available (v.2.2.0).

The fix has been merge in the upstream main and will be available in version v2.3.0.

## Changes

## CI/CD

* We created a new release-petal](.github/actions/release-petal/action.yml) Github action to release the container image on Petal's registry.
* We added a (release-petal](.github/workflows/release-petal.yml) Github workflow to release the container image on the petal specific tags.
