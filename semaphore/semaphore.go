package semaphore

type inventory struct {
	content     string
	inventoryId int
}

type template struct {
	templateId int
	inventory
}

type task struct {
	projectId    int
	repositoryId int
	inventoryId  int
	template     template
	environment  string
	playbook     string
	debug        bool
	dryRun       bool
}

func (t *task) setTemplate() error{
	return nil
}

func ()