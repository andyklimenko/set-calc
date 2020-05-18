# set-calc build instruction

# Clone
git clone git@github.com:andyklimenko/set-calc.git

# Build and run tests
    $ cd set-calc
    $ make
    
# Usage
    $ ./set-calc "[expression]"
Where expression can be:
* name of file that contains numbers
* set-operation SUM | INT | DIF with list of file-names to operate with
* combination of previous two points

Examples:
* `$ ./set-calc [ DIF a.txt b.txt c.txt ]`
* `$ ./set-calc [ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]`
* `$ ./set-calc [ SUM [ DIF a.txt [ SUM b.txt c.txt ] ] [ INT a.txt b.txt ] ]`