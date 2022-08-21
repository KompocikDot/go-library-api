# GO LIBRARY API

## Project purpose
Project was made to train my GO skills and to be more familiarized with `fiber` framework

#
## Technology/Frameworks used
- Fiber - [Github](https://github.com/gofiber/fiber)
- Postgresql
- PGX - [Github](https://github.com/jackc/pgx)

#
## Requirements
- Postgresql
- GO
- Make(only if you want to run makefile commands)

#
## How to use
1. Clone project `git clone  git@github.com:KompocikDot/go-library-api.git`
2. Change dir to project `cd go=library-api/cmd`
3. Create `.env` file with `DB_URL` variable set to postgresql url
    - If db is not created create one and create table called books
4. Run makefile command `make build`
