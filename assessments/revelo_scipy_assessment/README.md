# Revelo's Assessment

Do **NOT** open pull requests to this repository. If you do, your application will be immediately discarded.

## Thoughts

- I started by taking a high-level look at the assessment's instructions.
- I was warned that the Dockerfile has a few issues, so let's check that out by running `docker run --rm -it $(docker build -q .)`
- 11:47: Okay, there's an issue while installing dependencies. Let's run a container from the base image and try to install dependencies manually: `docker run -it python:3.8-slim-buster /bin/bash`
- 11:48: Maybe add a `-y` at the end of the `apt-get install` command is all we need, but let's try to install dependencies manually anyway
- 11:55: Bingo!, `-y` solved the issue
- 11:56: Next, let's find out why `sed -i \"s/ISRELEASED = False/ISRELEASED = True/\" setup.py` is failing...
- 12:32: In summary, the Dockerfile issues were fixed by using `-y` in the `apt-get install`, fixing the wrongly escaped 'double-quotes' in the sed commands and adding a line `COPY . .` right after `WORKDIR /usr/src/app`.
