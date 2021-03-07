# Line Server

A simple line server implemented in golang.

Usage:
```
line-server help
NAME:
   line-server - Serves lines from a file via an http interface

USAGE:
   line-server [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file_path value, --fp value    path to the file to serve [$FILE_PATH]
   --server_port value, --sp value  Server port (default: 8889) [$SERVER_PORT]
   --help, -h                       show help (default: false)
```

## How does the system work?

This implementation of the line server its loosely based on the process of memory pagination.
It uses a index to map each "page" / line on a block of storage. Given an page index
the service looks for its start and end in the index to seek its position on the block.

### Creating an index of all lines

When the line server starts it sweeps across the target file to build an index in memory
containing the file path, the total number of lines found, and an index of all line starts
and endings.

The structure looks like this in memory:
```
{
    "path to file"                      // path to file
    7                                   // total number of lines in the file
    [0:[start, end],[1:start, end] ...] // a continues block of memory with an index of all lines
}
```

### Serving the lines

Given a positive line number we check if the number is with the lines range
we navigate the index to the current line position, read the start and end values for the requested
line.

After this we can seek the line start on the file and read it until the end of the line.

If the number is invalid a 404 or 413 http code is return to the user.

## How will your system perform with a 1 GB file? a 10 GB file? a 100 GB file?

Each line index its stored in 16 bytes of memory. So we should be able to store 67,108,864 lines indexes per gigabyte of memory.
Since we do not store any parts of the file in memory and perform seeks to read the file the velocity of the system tied to the
velocity of its storage.

## How will your system perform with 100 users? 10000 users? 1000000 users?

For each request the line server opens a new file pointer to read the file. Depending on the configuration
of your OS with more concurrent request the server my reach the maximum permitted files open.

## What documentation, websites, papers, etc did you consult in doing this assignment?

For this assignment I visit the wikipedia page about memory pagination and the golang standard library documentation.

- https://en.wikipedia.org/wiki/Memory_paging
- https://golang.org/pkg/


## What third-party libraries or other tools does the system use? How did you choose each library or framework you used?

This project uses the standard golang project layout.
https://github.com/golang-standards/project-layout

Quoting:
> This is a basic layout for Go application projects. It's not an official standard defined by the core Go dev team;
> however, it is a set of common historical and emerging project layout patterns in the Go ecosystem.
> Some of these patterns are more popular than others. It also has a number of small enhancements along
> with several supporting directories common to any large enough real world application.

A library to facilitate the usage of input parameters.
https://github.com/urfave/cli

A library to convert unicode to ascii.
https://github.com/mozillazg/go-unidecode


I choose to use those to libraries because dealing with input parameters is not the core domain of our application and
converting unicode to ascii could be cover by an well tested external library vs home brewing it in such a short time.


## How long did you spend on this exercise? If you had unlimited more time to spend on this,
how would you spend it and how would you prioritize each item?

The exercise was realised in one lazy Sunday.

For this project we only implement the first step of memory pagination. Every request is served as a page fault,
we locate the lines on disk and return it immediately to the user. If I had more time I would continue to implement
the process of memory pagination fully. The next step would be store the lines in memory to faster serve the users.

This would also mitigate the issue of number of files open.

Lastly I would address the number of files open, using channels and goroutines.

Please note that the number of allowed open files in a system is configurable.


## If you were to critique your code, what would you have to say about it?

In this project I tried to implement a solid foundation that would let the project continue
to implement a line server using memory pagination model.

For small files it could be better to just load the all file into memory instead of all
the hassle of using memory pagination. The real benefit only occurs if a file is bigger than the
available memory of the system.

The code is well organized, tries to be KISS and SOLID friendly, its test all our required business scenarios, uses dependency
injection. This will allowing us, in the future to replace, our line reader for something else if needed.
