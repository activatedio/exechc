package exechc

import (
	"bytes"
	"os/exec"
	"reflect"

	"github.com/google/cel-go/cel"
)

const (
	keyStdout = "stdout"
	keyStderr = "stderr"
)

type checker struct {
	cmd string
	prg cel.Program
}

func (c checker) Check() (bool, error) {

	so := &bytes.Buffer{}
	se := &bytes.Buffer{}

	cmd := exec.Command("sh", "-c", c.cmd) //nolint:gosec // okay to pass in this command from config

	cmd.Stdout = so
	cmd.Stderr = se

	err := cmd.Run()

	if err != nil {
		return false, err
	}

	v, _, err := c.prg.Eval(map[string]any{
		keyStdout: so.String(),
		keyStderr: se.String(),
	})

	if err != nil {
		return false, err
	}

	res, err := v.ConvertToNative(reflect.TypeFor[bool]())

	if err != nil {
		return false, err
	}

	return res.(bool), nil

}

// NewChecker constructs a checker
func NewChecker(cfg *Runtime) Checker {

	env, err := cel.NewEnv(
		cel.Variable(keyStdout, cel.StringType),
		cel.Variable(keyStderr, cel.StringType),
	)

	Must(err)

	ast, is := env.Compile(cfg.CheckExpression)

	if is != nil && is.Err() != nil {
		panic(is.Err())
	}

	prg, err := env.Program(ast)

	Must(err)

	return &checker{
		cmd: cfg.CheckCmd,
		prg: prg,
	}
}
