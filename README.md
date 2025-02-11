# skiplist


## https://cmps-people.ok.ubc.ca/ylucet/DS/SkipList.html 
## So what is a skiplist? It is made up of levels, each of which is a linked list.
## At the bottom level the linked list has all your data in.
## The layer above has half the data in and the layer above has half of that, and so on.
## If a piece of data is on one level it most be on all the ones below it as well,
## and we link these together to let us descend the structure.
