# Yoti coding exercise

Client / Server application for encrypting values and storing them with ids.

## Prerequesites

golang and set up `$GOPATH`

## Compiling

run `make`

## Running

run `./run.sh`

## Interactive commands

the commands are:

- `store <ID> yourstringvalue`
- `retreive <ID> <ABOVE_RESULT>`

please be aware that the lexing algorithm is brutally primitive, and won't
accept spaces in the string to encrypt it tokenises on space characters, so the
spaces in the input will have to match this format exactly

## Remarks

This was a very fun exercise for me, and I have been able to discover lots of
areas of computer science which had been living behind a veil of _sort of
knowing_ about them, keeping them unexplored. Especially reading about
encryption algorithms has been enjoyable and challenging.

I appreciate there are plenty of areas I have glossed over in my
implementation, but I think it will serve as a very good starting point for
further discussion. It has been a real pleasure working with the protobuf /
grpc tool chain, and I am very impressed by the maturity of the project, and
flexibility it enables for polyglot service architectures.

