package runtime

type (
	LambdaNewRuntime struct {
		root *Object
	}
)

func New() (*LambdaNewRuntime, error) {

	r, err := createRootObject()
	if err != nil {
		return nil, err
	}

	// initialize the runtime with defaults
	rt := LambdaNewRuntime{
		root: r,
	}

	return &rt, nil
}

func createRootObject() (*Object, error) {
	o := NewObject("root", ObjectTypeNode, nil)
	return o, nil
}
