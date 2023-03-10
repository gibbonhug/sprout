# TODO

## Frontend/visual
- [ ] Decide how to style app
- [ ] Display Flowers
  - [ ] Programatically shaped SVGs
  - [x] Create initial flower SVG and use in App, fill color taken from props
- [ ] State management, esp with currently selected box & being able to clear selected boxes
- [x] Initial logic to fetch, render and style Boxes and Flowers
- [x] Code for clicking on a box
- [x] Decide file/folder structure for nested react components' SCSS


## Backend/data
- [ ] Create logic, endpoints for breeding & cloning flowers, adding to screen
  - [ ] Pair logic & API
    - [ ] Create POST endpoint for new pair, which calls pairing logic
    - [ ] Create Go function to breed two flowers (weighted average towards middle of 2 colors); Later will have mutation % as part of Flower DNA etc etc
      - [x] Create initial Go func to breed two flowers (50/50 blend of flower petal DNA)
        - [x] Create initial Go func to combine two colors
        - [x] Find library to combine colors
  - [ ] Clone logic & API
    - [x] Create Go func to create a clone child
    - [ ] Create POST endpoint for new clone, which calls cloning logic
  - [x] Create POST endpoint for creating a flower
  - [x] Create GET endpoints for flower, rlns etc by id
- [ ] Finish structuring backend more (not just main.go)
  - [ ] Create final structs for Boxes (flower pots/patches) and Flowers, relationships, etc
  - [x] Design and create temporary structs
  - [x] Create new flower package
  - [x] Move HTTP handlers outside of main.go to declutter
- [ ] Store data in DB (user login)
  - [ ] Complete rework of current code to use database
    - [x] Figure out database connection .env
    - [x] Use pgx global connection
    - [x] Fetch flowers from db
    - [ ] ...Fetch other data from DB
    - [ ] Change ugly temp fetching from DB to better functions
  - [x] Connect to database
    - [x] Fill database with initial tables and rows (copy of JSON data), finish setup
      - [x] Decide database
  - [ ] User login/authentication
- [x] Temporary JSON database

## Etc
- [ ] Add more more to TODO
- [x] Add more to TODO, clean TODO
- [x] Create dummy local app
  - [x] Dummy frontend: Display colored squares from JSON data
  - [x] Dummy backend: Serve dummy JSON data on localhost
- [x] Create TODO