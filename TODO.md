# TODO

## Frontend/visual
- [ ] Create actual flower svgs (Unsure how to do this programmatically with genetically determining number of petals, petal shape, etc; for now just use color DNA/data)
- [ ] Decide how to style app

## Backend/data
- [ ] Structure backend more (not just main.go)
  - [ ] Create models for Boxes (flower pots/patches) and Flowers, etc; Boxes have a Flower field and there are default 8 Boxes
- [ ] Store data in local storage + DB (user login)
  - [ ] Create database | decide database
  - [ ] User login/authentication
- [ ] Create Go function to breed two flowers (weighted average towards middle of 2 colors); Later will have mutation % as part of Flower DNA etc etc

## Etc
- [ ] Add more to TODO
- [x] Create dummy local app
  - [x] Dummy frontend: Display colored squares from JSON data
  - [x] Dummy backend: Serve dummy JSON data on localhost
- [x] Create TODO