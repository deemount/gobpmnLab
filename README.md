# gobpmnLab

The project has moved. gobpmn is now called gobpmnLab. This repository is gobpmn's experimental laboratory, i.e. this is where the new modularity is being trialled, the shift away from the idea of dependent reflection towards a "read and write", with no frills.

## Examples

The following examples are particularly worth mentioning in this repository

1. Create File
2. Create Instance
3. Create Log
4. View Elements and Methods

### Create File

To create a BPMN file with a model from the given examples, you should perform the following steps, which are described below

First, clone the repository to your GOPATH (e.g. "~/go/src/github.com/[your_username]/gobpmnLab")

with HTTPS

```bash
git clone https://github.com/deemount/gobpmnLab.git
```

or with SSH

```bash
git clone git@github.com:deemount/gobpmnLab.git
```

Then, change directory to the repository

```bash
cd gobpmnLab
```

Finally, enter the following command in your shell

```go
go run main.go
```

The result is a file in BPMN and JSON format, which is located in the /files/bpmn or /files/json directory. The name has a prefix and a consecutive number. A file with the name "diagram_[NUM].bpmn" or "diagram_[NUM].json" is therefore created after each command.

The settings for this execution are located in the /repository directory in the BpmnFactory.go file.

Open the file with an editor and make the change in the set method (see example below)

```go
func (factory *bpmnFactory) set() {

	log.Println("BpmnFactory.go: set model")

	// e.g. use interface pointer (no argument)

	//def := new(models.Definitions)
	//repositoryModel := examples.NewModel(def)

	// e.g. use struct pointer (no argument)

	// 1
	//factory.Repo = small_process.New().Build().Call()
	factory.Repo = collaborative_process.New().Build().Call()

	// 2
	// repositoryModel := examples.NewBlackBoxModel()
	//d := repositoryModel.Create()
	//factory.Repo = d.Def()

	// 3
	//repositoryModel := examples.NewSimpleModel001()
	//d := repositoryModel.Create()
	//factory.Repo = d.Def()

}
```

**Note:**
For some examples, you must also make a change in the toBPMN method. Simply comment out the line of code marked as a comment.

### Create Instance

Once the models have been created, the BPMN file is read and a process instance is created. Not worth mentioning, it is an experiment, but can be viewed in the /repository directory in the BpmnEngine.go file.

### Create Log

I log as much as possible of what happens during execution in order to track what exactly happens where. This log can be viewed once in the terminal after executing **go run main.go** and as a more comprehensive version in the /logs directory, in the file **build.log**. There you will find further information that can help you to get an overview.

### View Elements and Methods

There is an attempt to create an interface using Go and HTML in the /public directory. The following steps must be carried out here:

```bash
cd public
```

```go
go run main.go
```

Now open the browser at [localhost:8080](https://localhost:8080)

## Wiki

There's actually no wiki for this repository installed ...

**Start here** [gobpmn](https://github.com/deemount/gobpmn)

## Modules

+ [gobpmnTypes](https://github.com/deemount/gobpmnTypes)
+ [gobpmnModels](https://github.com/deemount/gobpmnModels)
+ [gobpmnCamunda](https://github.com/deemount/gobpmnCamunda)
+ [gobpmnDiagram](https://github.com/deemount/gobpmnDiagram)
+ [gobpmnBuilder](https://github.com/deemount/gobpmnBuilder)
+ [gobpmnReflection](https://github.com/deemount/gobpmnReflection)

**Note:** gobpmnLab already imports the module gobpmnTypes and gobpmnDiagram (not complete).
