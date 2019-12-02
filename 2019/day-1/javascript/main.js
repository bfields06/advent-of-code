const fs = require('fs')

let total = 0
let totalWithFuel = 0

var numsByLine = fs.readFileSync('input.txt').toString().split('\n').map(Number)

// function to calculate fuel based on mass
let fuelByMass = function (mass) {
  fuel = Math.floor(mass / 3) - 2

  if (fuel > 0) return fuel

  return 0
}

// Part 1
numsByLine.forEach(element => {
  f = fuelByMass(element)
  total += f
});

// Part 2
numsByLine.forEach(element => {
  mass = element
  do {
    f = fuelByMass(mass)
    totalWithFuel += f
    mass = f
  } while (f != 0)
});

console.log(total)
console.log(totalWithFuel)