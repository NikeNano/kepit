# kepit
Golang cli to remashal files to keep all fields even if empty and set to be empty. Work in progress. 


## Examples

### One

```bash

rm examples/examplePersonOut.jso

kepit run -f examples/examplePerson.go -j examples/examplePerson.json -s ExamplePerson -o examples/examplePersonOut.json
```

### Two

```bash

kepit run -f examples/exampleAnimal.go -j examples/exampleAnimal.json -s Animal -o examples/exampleAnimalOut.json
```


## To do 

- Make it work with directories to run on multiple json objects as well
- Make ir work with files that hold multiple json objects in one file, not true json I guess or.....
- Clean up the code base. 


