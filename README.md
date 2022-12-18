# Sprout
Flower breeding game built with Go and TypeScript/React

## Running the app
The app can currently be run locally by serving the frontend with Vite on port 5173 with 'npm run dev' command; and serving the backend on port 3000 by running the main.go file. 

Currently, this will show two colored divs when visiting the 5173 localhost. These divs take color data (hex colors) from the port 3000/flowers GET endpoint.

## Project Structure
### Overall
The base of the directory is split into 2 directories: /client for frontend/client-side code and /server for backend/server-side code.

### React components & prop interfaces
All React components should have their .tsx file stored in the /components directory inside /client/src. If they have props, prop interfaces are found inside the interfaces.tsx file found inside /client/src, which should be imported into the relevant component.

### Frontend styling
All SCSS/CSS should remain in the /scss directory, except for inline dynamic styling taken from data (for example, color data used to style Flowers). For React components, each component should have an associated file in the /scss/react_components directory that is imported to the component's .tsx file. Other SCSS files should be private (prepended with an underscore '_') and @imported into the main.scss file.

### Server-side
Currently, all backend code is found inside the main.go file.

## Inspiration
Sprout is heavily inspired by the game [Seed](https://www.noio.nl/2007/12/seed/), by Thomas van den Berg. Seed is one of my favorite Flash games and in the past I have run a tab of Seed in the background for many hours at a time. I decided to create a Seed-inspired game with additional features as a tribute, for fun, and as a way to build my skills.