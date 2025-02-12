package shell

import (
	"bufio"
	"context"
	"io"
	"os"
	"strings"
)

type OneshotShell struct {
	endLineFn     EndLineFn
	textFn        TextFn
	commentPrefix string
	stderr        io.Writer
}

func NewOneshot(endLineFn EndLineFn, textFn TextFn) *OneshotShell {
	return &OneshotShell{
		endLineFn:     endLineFn,
		textFn:        textFn,
		commentPrefix: "",
		stderr:        os.Stderr,
	}
}

func (sh *OneshotShell) SetCommentPrefix(pfx string) {
	sh.commentPrefix = pfx
}

func (sh *OneshotShell) Run(ctx context.Context) error {
	if err := sh.readTextBasic(); err != nil {
		return err
	}
	return nil
}

func (sh *OneshotShell) readTextBasic() error {
	// NOTE: when this implementation is changed,
	// clc/shell/shell.go:readTextReadline should also change!
	var sb strings.Builder
	multiline := false
	scn := bufio.NewScanner(os.Stdin)
	for scn.Scan() {
		if scn.Err() != nil {
			return scn.Err()
		}
		p := scn.Text()
		if !multiline {
			p = strings.TrimSpace(p)
			if p == "" {
				continue
			}
			if sh.commentPrefix != "" && strings.HasPrefix(p, sh.commentPrefix) {
				continue
			}
		}
		text, end := sh.endLineFn(p, multiline)
		sb.WriteString(text)
		multiline = !end
		if end {
			if err := sh.textFn(context.Background(), sb.String()); err != nil {
				return err
			}
			sb.Reset()
		}
	}
	text := sb.String()
	if text != "" {
		return sh.textFn(context.Background(), sb.String())
	}
	return nil
}
