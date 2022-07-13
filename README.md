# Shape 

## Problem description

A library to calculate perimeter and area of rectangle with sides of lengths a and b.  

## Development environment setup

After checking out the repo, import `github.com/stretchr/testify` to use `assert`.

### Prerequisite

- Go (1.18.2)

## Build instructions

To build this project use the following command:

    $ go build


## Test instructions

To run the test use the following command:

    $ go test

## Library usage

To create a new rectangle

    NewRectangle{height, width}

To create a new square

    NewSquare{side}

To find the perimeter of a rectangle

    NewRectangle{height, width}.Perimeter()

To find the area of a rectangle

    NewRectangle{height, width}.Area()

To find the perimeter of a square

    NewSquare{side}.Perimeter()

To find the area of a square

    NewSquare{side}.Area()