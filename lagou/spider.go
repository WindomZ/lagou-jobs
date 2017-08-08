package lagou

type Spider struct {
	Config
	Cookies string
}

func New() *Spider {
	return &Spider{}
}
