# TODO

## Frontend/visual
- [ ] Programatically shaped SVGs
- [ ] Fetch, render and style Boxes and Flowers
- [ ] Decide file/folder structure for nested react components' SCSS
- [ ] Decide how to style app
- [x] Code for clicking on a box
- [x] Create flower SVG and use in App, fill color taken from props

## Backend/data
- [ ] Create POST endpoints for breeding flowers
- [ ] Finish structuring backend more (not just main.go)
  - [ ] Create final structs for Boxes (flower pots/patches) and Flowers, relationships, etc
  - [x] Design and create temporary structs
  - [x] Create new flower package
- [ ] Store data in local storage + DB (user login)
  - [ ] Create database | decide database
  - [ ] User login/authentication
- [ ] Create Go function to breed two flowers (weighted average towards middle of 2 colors); Later will have mutation % as part of Flower DNA etc etc
- [x] Temporary JSON database

## Etc
- [ ] Add more to TODO
- [x] Create dummy local app
  - [x] Dummy frontend: Display colored squares from JSON data
  - [x] Dummy backend: Serve dummy JSON data on localhost
- [x] Create TODO