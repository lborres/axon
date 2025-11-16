# Axon Development

This document will briefly explain how to begin with development work on this project.

## Project Structure
```
axon/
├── apps/           # Main Application Code
│   └── client/     # Frontend. Tanstack Start with ReactJS
│   └── server/     # Backend. Go Server
├── migrations/     # Relational Database Migration Files
├── packages/       # Reusable packages. Reserved for both TS and Go.
├── tooling/        # Development Tooling
```

At the root level, this project has a few key subdirectories.
The `apps/` directory contains all the application code for both the ReactJS `client/` and the Go Server `server/`.

## Setup
### NodeJS
Please have NodeJS installed on your system. Please check the ITShared directory. Use any Node LTS version v24 and up.

### pnpm
The project uses [turborepo](https://turborepo.com/) to power the monolithic project structure.
The project also uses [pnpm](https://pnpm.io/installation) to power JS workspaces.

Please have `pnpm` installed.
You may install using `npm`.
```sh
npm install -g pnpm@latest-10
```

### Golang
The API Server runs on the Go programming language.
Please check the ITShared directory under requests for the installer.

### Post-setup
Once the above are installed, please install the project dependencies by running the following at the project root.
```sh
pnpm install
# or
pnpm i
```

This installs the nodejs dependencies.
Go will download dependencies when the api server is run.

## Usage
Thanks to turborepo and pnpm workspaces, most of the commands can be run at the root of the project.

```sh
# At the project root
pnpm dev
```
The `dev` script will run the development environment for the client. Running the server will be done separately for this case.

```sh
# Under apps/server
go run ./cmd/api
```

This is intended for working when development is being done on the server. Where the developer will need to run and stop the server manually multiple times during development.
(Hot Module Reloading for go can be considered in the future)

### Frontend Development
If the developer doesn't intend to modify the API server code, a script is prepared to allow both the client and the server to run in one terminal.

```sh
# At the project root
pnpm dev:client
```

The `dev:client` script is intended for development on the client only. This is a convenience script to allow the developer to avoid needing to open another terminal for the api server.

