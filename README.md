# FUN-project
## General
    This is an implementation of the article [Skiing Is Easy, Gymnastics Is Hard: Complexity of Routine Construction in Olympic Sports](https://drops.dagstuhl.de/opus/volltexte/2022/15987/pdf/LIPIcs-FUN-2022-17.pdf).
    The program calculates value of a routine according to the rules described in the article.
    Additionaly, a generator of all possible sequences up to the desired length is also available.
## Usage
### Run
1. Run
    ```shell
   go run main.go -file=<configFilePath> -seq=<sequence> -gen=<length>
    ```
### Build
1. Build
    ```shell
   go build
    ```
2. Run
    ```shell
   ./fun-project -file=<configFilePath> -seq=<sequence> -gen=<length>
    ```

### Help
1. Flags
    * -file (Required) : points to config file of the sport
    * -seq : single sequence to evaluate
    * -gen : upper length of sequences to generate and test, returns sequences witht the highest value
    * -gen-min : minimum length of sequences to generate, default: true
    * -print-all : print all optimal sequences instead of one, default: false

2. Sequences \
    Sequence elements must be delimited with a dot. \
    Example: A.B.A.C
