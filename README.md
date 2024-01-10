# Hangman Web

## Contents

- I. [ About the project](#i-about-the-project)
- II. [ Installation](#ii-installation)
- III. [ Features](#iii-features)
    - [1. Main](#1-main)
    - [2. Bring-to-death](#2-bring-to-death)
    - [3. Gamify](#3-gamify)
    - [4. Level](#4-level)

## I. About the project

The Hangman Web project aims to recreate the Hangman game on a web interface without JavaScript, using only Golang, HTML and CSS.

## II. Installation

To launch the project,

1. Clone the project:

    ```bash
    https://ytrack.learn.ynov.com/git/gyael/hangman-web
    ```

2. Change into the project directory:

    ```bash
    cd hangman-web    
    ```
All you have to do now is run the command `go run server.go`, and use your browser to go to the following link: [http://localhost:8080](http://localhost:8080).

## III. Features

### A. Main

The main task was to create a web interface for the Hangman game, using the Golang programming language.


### B. Bring-to-death

"Bring-to-death" consists in creating a hangman that is updated as the game progresses.

Each time a wrong suggestion is made (letter or word), the hanged man appears more and more.


### C. Gamify

"Gamify" involves making the interface more like a video game.

Create new interfaces such as: a login page (to choose your Username), an end-of-game screen (win/lose), a button to play again, a scoreboard, a list of letters already used during the game, etc.


### D. Level

"Level" adds a choice of difficulty before starting a game, allowing the user to choose between three difficulties: EASY / MEDIUM / HARD.