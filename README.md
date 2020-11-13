# Simple password generator

Generating a password list based on the one password.

![alt text](assets/passgen.png "passgen")

## Build & Clean

````
$ make
$ make clean
````

## Usage

````
$ ./passgen
$ ./passgen 20 // count of passwords, ≥ 1
$ ./passgen 20 32 // length of password, 1 ≤ length ≤ 32
````