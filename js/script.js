// Copyright (c) Bonnie Zhu
//
// Created by: Bonnie Zhu
// Created on: May 2023
// This file contains the JS functions for index.html
"use strict"

function calculate() {
  const age = document.getElementById("age").value

  if (age >= 18) {
    document.getElementById("answer").innerHTML =
      "You can see G, PG and R rated movies!"
  } else if (age >= 13) {
    document.getElementById("answer").innerHTML =
      "You can see G and PG rated movies!"
  } else if (age >= 5) {
    document.getElementById("answer").innerHTML = "You can see G rated movies!"
  } else {
    document.getElementById("answer").innerHTML =
      "you can hardly watch any movies"
  }
  console.log("Thanks for verifying you're age!")
}