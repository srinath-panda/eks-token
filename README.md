# eks-token
Repository to generate EKS token 

This is going to do exactly what `aws eks get-token` is going to do, but with a twist of golang.

## Why?

Bcaz we need to get the token to call EKS and aws eks get token needs python and aws cli installed in atlantis and atlantis doesnt have both and its a pain to install aws cli in alpine (atlantis) image. 


## How to use?

download and Mount the ./bin/eks-token<arch> in the atlanits container using extra voulmes. to the atlantis image and see [example](./terraform/main.tf) for usage

