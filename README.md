# quikon-api

Schema of the API types that are served by quikon(渠成).

## Purpose

This library is the canonical location of the Quikon API definition.

We recommend using the go types in this repo. You may serialize them directly to JSON.

## Compatibility matrix

| Kubernetes Version in your Project | Qucheng(Quikon) Version in your Project  | Import Quikon-api >= v0.10 |
|------------------------------------|----------------------------|----------------------------|
| >= 1.18                            | v1.x.y (x>1)          |                                      |

## Where does it come from?

`quikon-api` is synced from [https://github.com/easysoft/qucheng-operator/tree/master/apis](https://github.com/easysoft/qucheng-operator/tree/master/apis).
Code changes are made in that location, merged into `easysoft/qucheng-operator` and later synced here.

## Things you should NOT do

[https://github.com/easysoft/qucheng-operator/tree/master/apis](https://github.com/easysoft/qucheng-operator/tree/master/apis) is synced to here.
All changes must be made in the former. The latter is read-only.
