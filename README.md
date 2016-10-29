# Tny url shortener

This is a url shortener with the following stack:
        Frontend -> Vue.js
        Backend -> Go

I built this as a fun project to teach myself Vue.js and Go. It is currently still
in development and is not ready for production, but I wanted to release the code
publically.

This runs ontop of Nginx, and routes everything except static files, api calls, and /:linkid
calls to the Vue.js instance which then routes them to any pages which have been defined
in the route config.

All /api/ and /:linkid routes are handled by the Go backend and routed to specified
handlers in the main.go file. In order to understand the backend it would be easiest
to just follow that file.

The frontend runs ontop of Vue2.0 and utilizes Vue Router, and Vuex for handling state.

When this project is ready for deployment I will upload it to my server and post
the url in this repo.